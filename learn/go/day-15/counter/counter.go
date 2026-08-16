package counter


import (
	"bufio"
	"fmt"
	"os"
)

func CounterWords(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil{
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	words := make(map[string]int)
	for scanner.Scan(){
		words[scanner.Text()]++
	}
	return words, nil
}