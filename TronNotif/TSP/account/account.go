package account

import (
	"strconv"
	"time"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// token Info
type Token struct {
	TokenName    string
	TokenDecimal int
}

// tokenName -> Token
type Tokens map[string]Token

// NoUpdateCount: used to determine which category
// the address should be in. (increment by 1 when balance not updated)
//
//	      NoUpdateCount <= 360 (6 hrs)  -> Normal
//	360 < NoUpdateCount <= 720 (10 hrs) -> Seldom
//	720 < NoUpdateCount                 -> Frozen
type Balance struct {
	Token         Token
	Amount        string
	NoUpdateCount int
}

// tokenName -> Balance
type Balances map[string]Balance

// address -> Balances
type Accounts map[string]Balances

// Returns balance in the human-readable format
// by adding decimal & thousands separator and
func (bal Balance) GetReadbleAmount() (string, error) {

	decimalPlace := bal.Token.TokenDecimal
	amount := bal.Amount
	p := message.NewPrinter(language.English)

	// amount >= 1
	if len(amount) > decimalPlace {
		decIndex := len(amount) - decimalPlace
		beforeDecimal := amount[:decIndex]
		n, err := strconv.Atoi(beforeDecimal)

		if err != nil { return "", err }
		// add thousands separator
		formattedBeforeDecimal := p.Sprintf("%d",  n)
		afterDecimal := amount[decIndex:]
		return formattedBeforeDecimal + "." + afterDecimal, nil
	}  

	// amount < 1

	beforeDecimal := "0"
	// left pad with 0 
	afterDecimal := p.Sprintf("%0*s", decimalPlace, amount)

	return  beforeDecimal + "." + afterDecimal, nil

}


// API Call frequecies:
// API Call frequecies:
// FREQ_1: (Normal) every 1 min
// FREQ_2: (Seldom)       2 min
// FREQ_3: (Frozen)       3 min

const (
	FREQ_1           = 1 * time.Second
	FREQ_2           = 2 * time.Second
	FREQ_3           = 3 * time.Second
	NUMBER_OF_QUEUES = 3
)

// Store each category in its own slice
type AccountDispatch struct {
	TokenTable Tokens
	Normal     Accounts
	Seldom     Accounts
	Frozen     Accounts
}

func (f FreqStat) String() string {
	switch f {
	case Normal:
		return "Normal"
	case Seldom:
		return "Seldom"
	case Frozen:
		return "Frozen"
	}
	return "<unknown>"
}

type FreqStat int8

const (
	Normal FreqStat = iota
	Seldom
	Frozen
)
