package handler

import "fmt"

const (
	BASE_URL   = "https://apilist.tronscanapi.com/api/"
	ADDR_URL   = BASE_URL + "account/tokens?address=%s&token=%s"
	USDT_TOKEN = "USDT"
)

// get url for Get account's token list
// This endpoint returns a list of tokens held by the account with a balance greater than 0.
func GetTronTokenURL(addr string, token string) string {
	return fmt.Sprintf(ADDR_URL, addr, token)
}
