package storage

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func ReadMenu(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("menu file not found at %q: %w", path, err)

		}
		return nil, fmt.Errorf("read menu %w", err)
	}
	return data, nil
}

func SaveOrder(path string, content string) error {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("write order: %w", err)
	}
	return nil
}

func ParseMenuLines(data []byte) map[string]int {
	menu := make(map[string]int)
	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		price, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		menu[parts[0]] = price
	}
	return menu
}
