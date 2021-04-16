package validator_test

import (
	validator "github.com/MrWormHole/simple-validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	param    string
	expected bool
}

func TestIsEmpty(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"asdas", false},
		{"", true},
		{"   ", true},
	}

	for _, t := range testCases {
		actual := validator.IsEmpty(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsHTML(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"<html>", true},
		{"</script>", true},
		{"<stillworks>", true},
		{"</html", false},
		{"<script></script>", true},
		{"<//script>", false},
		{"<123nonsense>", false},
		{"test", false},
		{"", false},
	}

	for _, t := range testCases {
		actual := validator.IsHTML(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsHTMLEncoded(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"&#x3c;", true},
		{"&#xaf;", true},
		{"&#x00;", true},
		{"&#xf0;", true},
		{"&#x3c", true},
		{"&#xaf", true},
		{"&#x00", true},
		{"&#xf0", true},
		{"&#ab", true},
		{"&lt;", true},
		{"&gt;", true},
		{"&quot;", true},
		{"&amp;", true},
		{"#x0a", false},
		{"&x00", false},
		{"&#x1z", false},
		{"", false},
	}

	for _, t := range testCases {
		actual := validator.IsHTMLEncoded(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsURLEncoded(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"%20", true},
		{"%af", true},
		{"%ff", true},
		{"<%az", false},
		{"%test%", false},
		{"a%b", false},
		{"1%2", false},
		{"%%a%%", false},
		{"", false},
	}

	for _, t := range testCases {
		actual := validator.IsURLEncoded(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsETHAddress(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"0x52908400098527886E0F7030069857D2E4169EE7", true},
		{"0x8617E340B3D01FA5F11F306F4090FD50E238070D", true},
		{"0xde709f2102306220921060314715629080e2fb77", true},
		{"0x27b1fdb04752bbc536007a920d24acb045561c26", true},
		{"0x123f681646d4a755815f9cb19e1acc8565a0c2ac", true},
		{"0x02F9AE5f22EA3fA88F05780B30385bECFacbf130", true},
		{"0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed", true},
		{"0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359", true},
		{"0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB", true},
		{"0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb", true},
		{"0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDB", false}, // Invalid checksum.
		{"", false},
		{"D1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb", false},    // Missing "0x" prefix.
		{"0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDbc", false}, // More than 40 hex digits.
		{"0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aD", false},   // Less than 40 hex digits.
		{"0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDw", false},  // Invalid hex digit "w".
	}

	for _, t := range testCases {
		actual := validator.IsETHAddress(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsDomainName(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"test.example.com", true},
		{"example.com", true},
		{"example24.com", true},
		{"test.example24.com", true},
		{"test24.example24.com", true},
		{"example", false},
		{"EXAMPLE", false},
		{"1.foo.com", true},
		{"test.example.com.", false},
		{"example.com.   ", false},
		{"example24.com.", false},
		{"test.example24.com.", false},
		{"test24.example24.com.", false},
		{"example.", false},
		{"192.168.0.1", false},
		{"email@example.com", false},
		{"2001:cdba:0000:0000:0000:0000:3257:9652", false},
		{"2001:cdba:0:0:0:0:3257:9652", false},
		{"2001:cdba::3257:9652", false},
		{"example..........com", false},
		{"1234", false},
		{"abc1234", false},
		{"example. com", false},
		{"ex ample.com", false},
	}

	for _, t := range testCases {
		actual := validator.IsDomainName(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsIPAddress(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"1.2.3.4", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"::", true},
		{"ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff", true},
		{"2001::f:1234", true},
		{"1200:0000:AB00:1234:0000:2552:7777:1313", true},
		{"", false},
		{"   ", false},
		{"foo", false},
		{"01.02.03.04", false},
		{"256.256.256.256", false},
		{"1.2.3", false},
		{"1.2.3.4.5", false},
		{"-1.2.3.4.5", false},
		{":", false},
		{":::", false},
		{"2001::f::1234", false},
		{"2001:g::", false},
		{"1200:0000:AB00:1234:O000:2552:7777:1313", false},
	}

	for _, t := range testCases {
		actual := validator.IsIPAddress(t.param)
		assert.Equal(t.expected, actual)
	}
}
