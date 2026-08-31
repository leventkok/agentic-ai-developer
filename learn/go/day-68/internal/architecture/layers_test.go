package architecture_test

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func moduleRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("go.mod not found")
		}
		dir = parent
	}
}

func parseImports(t *testing.T, dir string) map[string][]string {
	t.Helper()
	imports := make(map[string][]string)
	fset := token.NewFileSet()
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return err
		}
		f, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			return err
		}
		for _, imp := range f.Imports {
			p := strings.Trim(imp.Path.Value, `"`)
			imports[path] = append(imports[path], p)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return imports
}

func TestDomainDoesNotImportOuterLayers(t *testing.T) {
	root := moduleRoot(t)
	imports := parseImports(t, filepath.Join(root, "internal", "domain"))
	forbidden := []string{"net/http", "database/sql", "internal/service", "internal/httpapi", "internal/repository"}
	for file, paths := range imports {
		for _, p := range paths {
			for _, f := range forbidden {
				if strings.Contains(p, f) {
					t.Fatalf("%s imports forbidden %s", file, p)
				}
			}
		}
	}
}

func TestServiceDoesNotImportHTTP(t *testing.T) {
	root := moduleRoot(t)
	imports := parseImports(t, filepath.Join(root, "internal", "service"))
	for file, paths := range imports {
		if strings.Contains(file, "testing"+string(filepath.Separator)+"fake") {
			continue
		}
		for _, p := range paths {
			if strings.Contains(p, "net/http") || strings.Contains(p, "internal/httpapi") {
				t.Fatalf("%s imports forbidden %s", file, p)
			}
		}
	}
}

func TestOnlyAppImportsSQLite(t *testing.T) {
	root := moduleRoot(t)
	dirs := []string{"internal/httpapi", "internal/service", "internal/domain", "internal/middleware"}
	for _, dir := range dirs {
		imports := parseImports(t, filepath.Join(root, dir))
		for file, paths := range imports {
			for _, p := range paths {
				if strings.Contains(p, "repository/sqlite") {
					t.Fatalf("%s imports sqlite directly: %s", file, p)
				}
			}
		}
	}
}

func TestNoImportCycles(t *testing.T) {
	root := moduleRoot(t)
	if _, err := parser.ParseDir(token.NewFileSet(), filepath.Join(root, "internal", "domain"), nil, parser.ImportsOnly); err != nil {
		t.Fatal(err)
	}
}
