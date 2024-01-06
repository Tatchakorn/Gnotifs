package account

import (
	"math"
	"slices"
	"strconv"
	"golang.org/x/text/message"
)

// token Info
type Token struct {
	TokenName    string
	TokenDecimal uint16
}

// Fields NoUpdateCount: If the fetched balance is the same as
// from the previous balane then increment by 1
type Balance struct {
	Token  Token
	Amount int64
}

type Account struct {
	Address string
	Balance Balance
}

type Accounts []Account

// Returns ac index of account from address
func (acc Accounts) GetIndexFromAddr(addr string) int {
	index := slices.IndexFunc[Accounts](acc, func(a Account) bool { return a.Address == addr })
	return index
}

// Returns balance in the human-readable format
// add decimal separator separator and
// and digit separator
func (bal Balance) GetAmountStr() string {
	p := message.NewPrinter(message.MatchLanguage("en"))
	decimalPlace := bal.Token.TokenDecimal
	beforeDecimalAmt := bal.Amount / int64(math.Pow(10, float64(decimalPlace)))
	beforeDecimal := p.Sprint(beforeDecimalAmt)
	Amount := strconv.Itoa(int(bal.Amount))
	afterDecimalAmt := Amount[len(Amount) - int(decimalPlace):]
	return  beforeDecimal + "." + afterDecimalAmt
}
