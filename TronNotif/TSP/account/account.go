package account

type Token struct {
	tokenName    string
	tokenDecimal uint16
}

type Balance struct {
	token   Token
	amount int64
}

type Account struct {}

type addressBal struct{}
type addresses []addressBal
