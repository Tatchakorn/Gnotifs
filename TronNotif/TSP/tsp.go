package main

import (
	"fmt"
	"tsp/TronNotif/TSP/account"
	"tsp/TronNotif/TSP/handler"
)

func main() {
	url := handler.GetTronTokenURL("Hello", "World")
	fmt.Println(url)
	usdtToken := account.Token{TokenName: "USDT", TokenDecimal: 6}
	b := account.Balance{
		Token: usdtToken,
		Amount: 6347833,
	}
	b.GetAmountStr()
}
