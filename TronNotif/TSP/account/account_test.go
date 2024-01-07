package account

import (
	"log"
	"testing"
)

// for single test
func setupAccountsSuite(tb testing.TB) (func(tb testing.TB), AccountDispatch) {
	log.Println("setup suite")

	tokenTable := map[string]Token{
		"USDT": {TokenName: "USDT", TokenDecimal: 10},
		"ETH":  {TokenName: "ETH", TokenDecimal: 18},
		"DOGE": {TokenName: "DOGE", TokenDecimal: 8},
	}

	ad := AccountDispatch{
		Normal: Accounts{
			"TUAAqYySyBJDLJDxnQKqtcPBW1JXJuqrSS": Balances{
				"USDT": Balance{
					Token:  tokenTable["USDT"],
					Amount: "130205193380000"},
				"ETH": Balance{
					Token:  tokenTable["ETH"],
					Amount: "1162312289316921"},
				"DOGE": Balance{
					Token:  tokenTable["DOGE"],
					Amount: "293532778"},
			},
			"TT2T17KZhoDu47i2E4FWxfG79zdkEWkU9N": Balances{
				"USDT": Balance{
					Token:  tokenTable["USDT"],
					Amount: "130200000000000"},
				"DOGE": Balance{
						Token:  tokenTable["DOGE"],
						Amount: "12311001"},
			},
		},
	}

	return func(tb testing.TB) {
		log.Println("teardown suite")
	}, ad
}

// for collection of tests
func setupAccountsTest(tb testing.TB) func(tb testing.TB) {
	log.Println("setup test")

	return func(tb testing.TB) {
		log.Println("teardown test")
	}
}

func TestAccount_GetReadbleAmount(t *testing.T) {

	teardownSuite, ad := setupAccountsSuite(t)
	defer teardownSuite(t)

	table := []struct {
		name     string
		bal      Balance
		expected string
	}{
		{
			"(USDT) decimal and digit separator",
			ad.Normal["TUAAqYySyBJDLJDxnQKqtcPBW1JXJuqrSS"]["USDT"],
			"13,020.5193380000",
		},
		{
			"(ETH) decimal only",
			ad.Normal["TUAAqYySyBJDLJDxnQKqtcPBW1JXJuqrSS"]["ETH"],
			"0.001162312289316921",
		},
		{
			"(DOGE) decimal only",
			ad.Normal["TUAAqYySyBJDLJDxnQKqtcPBW1JXJuqrSS"]["DOGE"],
			"2.93532778",
		},
		{
			"(USDT) No decimal",
			ad.Normal["TT2T17KZhoDu47i2E4FWxfG79zdkEWkU9N"]["USDT"],
			"13,020.0000000000",
		},
		{
			"(DOGE) No decimal",
			ad.Normal["TT2T17KZhoDu47i2E4FWxfG79zdkEWkU9N"]["DOGE"],
			"0.12311001",
		},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			teardownTest := setupAccountsTest(t)
			defer teardownTest(t)

			actual, err := tc.bal.GetReadbleAmount()

			if err != nil {
				t.Fatal(err)
			}

			if actual != tc.expected {
				t.Errorf("expected: \"%s\", got \"%s\"", tc.expected, actual)
			}
		})
	}

}
