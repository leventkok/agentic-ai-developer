

package jsonio




import (
	"encoding/json"
	"fmt"
	"os"
)

type MenuItem struct {
	Name string `json:"name"`
	Price int `json:"price"`
}

type Menu struct{
	Items []MenuItem `json:"items"`
}


type Order struct{
	Items []string `json:"items"`
	Total int `json:"total,omitempty"`
	Note *string `json:"note,omitempty"`
}


func MenuToJSON(menu Menu) ([]byte, error){
	return json.Marshal(menu)
}


func MenuFromJSON(data []byte) (Menu, error){
	var menu Menu
	err := json.Unmarshal(data, &menu)
	if err != nil{
		return Menu{}, fmt.Errorf("unmarshal menu: %w", err)
	}
	return menu, nil
}

func SaveMenu(path string, menu Menu) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create menu file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(menu); err != nil {
		return fmt.Errorf("encode menu: %w", err)
	}
	return nil
}

func LoadMenu(path string) (Menu, error) {
	file, err := os.Open(path)
	if err != nil {
		return Menu{}, fmt.Errorf("open menu file: %w", err)
	}
	defer file.Close()

	var menu Menu
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&menu); err != nil {
		return Menu{}, fmt.Errorf("decode menu: %w", err)
	}
	return menu, nil
}
