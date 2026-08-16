package report

import (
	"encoding/json"
	"fmt"
	"os"
)

type Summary struct {
	ShopName    string         `json:"shop_name"`
	TotalWords  int            `json:"total_words"`
	UniqueWords int            `json:"unique_words"`
	Counts      map[string]int `json:"counts"`
}

func Build(shopName string, counts map[string]int) Summary {
	total := 0
	for _, n := range counts {
		total += n
	}
	return Summary{
		ShopName:    shopName,
		TotalWords:  total,
		UniqueWords: len(counts),
		Counts:      counts,
	}
}

func Write(path string, summary Summary) error {
	data, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal summary: %w", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write summary: %w", err)
	}
	return nil
}
