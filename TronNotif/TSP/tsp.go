package main

import (
	"fmt"
	_ "tsp/TronNotif/TSP/account"
	"tsp/TronNotif/TSP/handler"
)

func main() {
	url := handler.GetTronTokenURL("Hello", "World")
	fmt.Println(url)
}
