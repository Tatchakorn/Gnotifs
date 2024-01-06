package account

import (
	"log"
	"testing"
)

// for single test
func setupAccountsSuite(tb testing.TB) (func(tb testing.TB), Accounts) {
	log.Println("setup suite")

	usdtToken := Token{TokenName: "USDT", TokenDecimal: 6}
	acc := Accounts{
		Account{
			Address: "addr1",
			Balance: Balance{
				Token: usdtToken, Amount: 29206347833}},
		Account{
			Address: "addr2",
			Balance: Balance{
				Token: usdtToken, Amount: 29206347833}},
		Account{
			Address: "addr3",
			Balance: Balance{
				Token: usdtToken, Amount: 29206347833}},
	}

	return func(tb testing.TB) {
		log.Println("teardown suite")
	}, acc
}

// for collection of tests
func setupAccountsTest(tb testing.TB) func(tb testing.TB) {
	log.Println("setup test")

	return func(tb testing.TB) {
		log.Println("teardown test")
	}
}

func TestAccounts_GetIndexFromAddr(t *testing.T) {
	teardownSuite, accounts := setupAccountsSuite(t)
	defer teardownSuite(t)

	table := []struct {
		name     string
		addr     string
		expected int
	}{
		{"address one", "addr1", 0},
		{"address rwo", "addr2", 1},
		{"not exist", "Noaddr", -1},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			teardownTest := setupAccountsTest(t)
			defer teardownTest(t)

			actual := accounts.GetIndexFromAddr(tc.addr)

			if actual != tc.expected {
				t.Errorf("expected: %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestAccount_GetAmountStr(t *testing.T) {

	usdtToken := Token{TokenName: "USDT", TokenDecimal: 6}

	table := []struct {
		name     string
		bal      Balance
		expected string
	}{
		{
			"With decimal and digit separator",
			Balance{
				Token:  usdtToken,
				Amount: 29206347833,
			},
			"29,206.347833",
		},
		{
			"With decimal",
			Balance{
				Token:  usdtToken,
				Amount: 6347833,
			},
			"6.347833",
		},
		{
			"With digit separator",
			Balance{
				Token:  usdtToken,
				Amount: 29206000000,
			},
			"29,206.000000",
		},
		{
			"With nothing",
			Balance{
				Token:  usdtToken,
				Amount: 6000000,
			},
			"6.000000",
		},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			teardownTest := setupAccountsTest(t)
			defer teardownTest(t)

			actual := tc.bal.GetAmountStr()

			if actual != tc.expected {
				t.Errorf("expected: \"%s\", got \"%s\"", tc.expected, actual)
			}
		})
	}

}
