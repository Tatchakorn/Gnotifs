package handler

import "testing"

func TestGetTronTokenURL(t *testing.T) {

	expect := "https://apilist.tronscanapi.com/api/account/tokens?address=testToken123sadf&token=USDT"
	result := GetTronTokenURL("testToken123sadf", "USDT")

	if expect != result {
		t.Errorf("wrong url!\n(expect):%s\n!=\n(result):%s", expect, result)
	}
}
