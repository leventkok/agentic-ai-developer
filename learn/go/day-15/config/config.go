


package config


import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ShopName   string `json:"shop_name"`
    InputFile  string `json:"input_file"`
    OutputFile string `json:"output_file"`
}


func (c Config) Validate() error {
	if c.ShopName == "" {
		return fmt.Errorf("shop name is required")
	}
	if c.InputFile == "" {
		return fmt.Errorf("input file is required")
	}
	if c.OutputFile == "" {
		return fmt.Errorf("output file is required")
	}
	return nil
}



func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil{
		return Config{}, fmt.Errorf("read config file: %w", err)
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil{
		return Config{}, fmt.Errorf("unmarshal config: %w", err)
	}
	return config, nil
}


 