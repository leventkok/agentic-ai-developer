package config

import (
	"fmt"
	"sync"
)

var (
	shopName string
	once     sync.Once
)

func loadShop() {
	fmt.Println("Loading shop config... (expensive, runs once)")
	shopName = "MasterFabric Cafe"
}

func ShopName() string {
	once.Do(loadShop)
	return shopName
}
