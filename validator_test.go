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

func TestIsLongitude(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"-180.000", true},
		{"180.1", false},
		{"+73.234", true},
		{"+382.3811", false},
		{"23.11111111", true},
		{"180", true},
		{"-180.0", true},
		{"-180", true},
		{"180.1", false},
	}

	for _, t := range testCases {
		actual := validator.IsLongitude(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsLatitude(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"-90.000", true},
		{"+90", true},
		{"47.1231231", true},
		{"+99.9", false},
		{"108", false},
		{"90", true},
		{"-90.0", true},
		{"-90", true},
		{"90.1", false},
	}

	for _, t := range testCases {
		actual := validator.IsLatitude(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsDataURI(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"data:image/png;base64,TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=", true},
		{"data:text/plain;base64,Vml2YW11cyBmZXJtZW50dW0gc2VtcGVyIHBvcnRhLg==", true},
		{"image/gif;base64,U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", false},
		{"data:image/gif;base64,MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuMPNS1Ufof9EW/M98FNw" +
			"UAKrwflsqVxaxQjBQnHQmiI7Vac40t8x7pIb8gLGV6wL7sBTJiPovJ0V7y7oc0Ye" +
			"rhKh0Rm4skP2z/jHwwZICgGzBvA0rH8xlhUiTvcwDCJ0kc+fh35hNt8srZQM4619" +
			"FTgB66Xmp4EtVyhpQV+t02g6NzK72oZI0vnAvqhpkxLeLiMCyrI416wHm5Tkukhx" +
			"QmcL2a6hNOyu0ixX/x2kSFXApEnVrJ+/IxGyfyw8kf4N2IZpW5nEP847lpfj0SZZ" +
			"Fwrd1mnfnDbYohX2zRptLy2ZUn06Qo9pkG5ntvFEPo9bfZeULtjYzIl6K8gJ2uGZHQIDAQAB", true},
		{"data:image/png;base64,12345", false},
		{"", false},
		{"data:text,:;base85,U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", false},
		{"data:image/jpeg;key=value;base64,UEsDBBQAAAAI", true},
		{"data:image/jpeg;key=value,UEsDBBQAAAAI", true},
		{"data:;base64;sdfgsdfgsdfasdfa=s,UEsDBBQAAAAI", true},
		{"data:,UEsDBBQAAAAI", true},
	}

	for _, t := range testCases {
		actual := validator.IsDataURI(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestHasMultibyteChar(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"abc", false},
		{"123", false},
		{"<>@;.-=", false},
		{"ひらがな・カタカナ、．漢字", true},
		{"あいうえお foobar", true},
		{"test＠example.com", true},
		{"test＠example.com", true},
		{"1234abcDEｘｙｚ", true},
		{"ｶﾀｶﾅ", true},
	}

	for _, t := range testCases {
		actual := validator.HasMultibyteChar(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsASCII(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", true},
		{"ｆｏｏbar", false},
		{"ｘｙｚ０９８", false},
		{"１２３456", false},
		{"ｶﾀｶﾅ", false},
		{"foobar", true},
		{"0987654321", true},
		{"test@example.com", true},
		{"1234abcDEF", true},
		{"", true},
	}

	for _, t := range testCases {
		actual := validator.IsASCII(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsPrintableASCII(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", true},
		{"ｆｏｏbar", false},
		{"ｘｙｚ０９８", false},
		{"１２３456", false},
		{"ｶﾀｶﾅ", false},
		{"foobar", true},
		{"0987654321", true},
		{"test@example.com", true},
		{"1234abcDEF", true},
		{"newline\n", false},
		{"\x19test\x7F", false},
	}

	for _, t := range testCases {
		actual := validator.IsPrintableASCII(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsUUIDMixed(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"xxxa987Fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"a987Fbc9-4bed-3078-cf07-9141ba07c9f3xxx", false},
		{"a987Fbc94bed3078cf079141ba07c9f3", false},
		{"934859", false},
		{"987fbc9-4bed-3078-cf07a-9141ba07c9F3", false},
		{"aaaaaaaa-1111-1111-aaaG-111111111111", false},
		{"a987Fbc9-4bed-3078-cf07-9141ba07c9f3", true},
	}

	for _, t := range testCases {
		actual := validator.IsUUIDMixed(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsUUID5Mixed(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"xxxa987Fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"9c858901-8a57-4791-81Fe-4c455b099bc9", false},
		{"a987Fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"987Fbc97-4bed-5078-af07-9141ba07c9f3", true},
		{"987Fbc97-4bed-5078-9f07-9141ba07c9f3", true},
	}

	for _, t := range testCases {
		actual := validator.IsUUID5Mixed(t.param)
		assert.Equal(t.expected, actual)
	}
}


func TestIsUUID4Mixed(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9F3", false},
		{"a987fbc9-4bed-5078-af07-9141ba07c9F3", false},
		{"934859", false},
		{"57b73598-8764-4ad0-a76A-679bb6640eb1", true},
		{"625e63f3-58f5-40b7-83a1-a72ad31acFfb", true},
	}

	for _, t := range testCases {
		actual := validator.IsUUID4Mixed(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsUUID3Mixed(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"412452646", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9F3", false},
		{"a987fbc9-4bed-4078-8f07-9141ba07c9F3", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9F3", true},
	}

	for _, t := range testCases {
		actual := validator.IsUUID3Mixed(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsUUID(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3xxx", false},
		{"a987fbc94bed3078cf079141ba07c9f3", false},
		{"934859", false},
		{"987fbc9-4bed-3078-cf07a-9141ba07c9f3", false},
		{"aaaaaaaa-1111-1111-aaag-111111111111", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3", true},
	}

	for _, t := range testCases {
		actual := validator.IsUUID(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsUUID5(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"9c858901-8a57-4791-81fe-4c455b099bc9", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"987fbc97-4bed-5078-af07-9141ba07c9f3", true},
		{"987fbc97-4bed-5078-9f07-9141ba07c9f3", true},
	}

	for _, t := range testCases {
		actual := validator.IsUUID5(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsUUID4(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"a987fbc9-4bed-5078-af07-9141ba07c9f3", false},
		{"934859", false},
		{"57b73598-8764-4ad0-a76a-679bb6640eb1", true},
		{"625e63f3-58f5-40b7-83a1-a72ad31acffb", true},
	}

	for _, t := range testCases {
		actual := validator.IsUUID4(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsUUID3(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"412452646", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"a987fbc9-4bed-4078-8f07-9141ba07c9f3", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3", true},
	}

	for _, t := range testCases {
		actual := validator.IsUUID3(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsISBN13(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"foo", false},
		{"3-8362-2119-5", false},
		{"01234567890ab", false},
		{"978 3 8362 2119 0", false},
		{"9784873113685", true},
		{"978-4-87311-368-5", true},
		{"978 3401013190", true},
		{"978-3-8362-2119-1", true},
	}

	for _, t := range testCases {
		actual := validator.IsISBN13(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsISBN10(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"foo", false},
		{"3423214121", false},
		{"978-3836221191", false},
		{"3-423-21412-1", false},
		{"3 423 21412 1", false},
		{"3836221195", true},
		{"1-61729-085-8", true},
		{"3 423 21412 0", true},
		{"3 401 01319 X", true},
	}

	for _, t := range testCases {
		actual := validator.IsISBN10(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsBase64URL(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		// empty string, although a valid base64 string, should fail
		{"", false},
		// invalid length
		{"a", false},
		// base64 with padding
		{"Zg==", true},
		{"Zm8=", true},
		// base64 without padding
		{"Zm9v", true},
		{"Zg", false},
		{"Zm8", false},
		// base64 URL safe encoding with invalid, special characters '+' and '/'
		{"FPucA9l+", false},
		{"FPucA/lz", false},
		// base64 URL safe encoding with valid, special characters '-' and '_'
		{"FPucA9l-", true},
		{"FPucA_lz", true},
		// non base64 characters
		{"@mc=", false},
		{"Zm 9", false},
	}

	for _, t := range testCases {
		actual := validator.IsBase64URL(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsBase64(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"dW5pY29ybg==", true},
		{"dGhpIGlzIGEgdGVzdCBiYXNlNjQ=", true},
		{"dW5pY29ybg== foo bar", false},
	}

	for _, t := range testCases {
		actual := validator.IsBase64(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsEmail(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"test", false},
		{"test@", false},
		{"test.com", false},
		{"test@test.com", true},
		{"test@test.org", true},
		{"jack!_24@hotmail.com", true},
	}

	for _, t := range testCases {
		actual := validator.IsEmail(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsHSL(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"hsl(360,100%,50%)", true},
		{"hsl(0,0%,0%)", true},
		{"hsl(361,100%,50%)", false},
		{"hsl(361,101%,50%)", false},
		{"hsl(361,100%,101%)", false},
		{"hsl(-10,100%,100%)", false},
		{"1", false},
	}

	for _, t := range testCases {
		actual := validator.IsHSL(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsHSLA(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"hsla(360,100%,100%,1)", true},
		{"hsla(360,100%,100%,0.5)", true},
		{"hsla(0,0%,0%, 0)", true},
		{"hsl(361,100%,50%,1)", false},
		{"hsl(361,100%,50%)", false},
		{"hsla(361,100%,50%)", false},
		{"hsla(360,101%,50%)", false},
		{"hsla(360,100%,101%)", false},
		{"1", false},
	}

	for _, t := range testCases {
		actual := validator.IsHSLA(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsRGB(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"rgb(0,31,255)", true},
		{"rgb(0,  31, 255)", true},
		{"rgb(10%,  50%, 100%)", true},
		{"rgb(10%,  50%, 55)", false},
		{"rgb(1,349,275)", false},
		{"rgb(01,31,255)", false},
		{"rgba(0,31,255)", false},
		{"1", false},
	}

	for _, t := range testCases {
		actual := validator.IsRGB(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsRGBA(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"rgba(0,31,255,0.5)", true},
		{"rgba(0,31,255,0.12)", true},
		{"rgba(12%,55%,100%,0.12)", true},
		{"rgba( 0,  31, 255, 0.5)", true},
		{"rgba(12%,55,100%,0.12)", false},
		{"rgb(0,  31, 255)", false},
		{"rgb(1,349,275,0.5)", false},
		{"rgb(01,31,255,0.5)", false},
		{"1", false},
	}

	for _, t := range testCases {
		actual := validator.IsRGBA(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsHexcolor(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"#fff", true},
		{"#c2c2c2", true},
		{"#aCbCcC", true},
		{"#123123", true},
		{"fff", false},
		{"fffFF", false},
		{"#GGG", false},
		{"1", false},
	}

	for _, t := range testCases {
		actual := validator.IsHexcolor(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsHexadecimal(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"ff0044", true},
		{"0xff0044", true},
		{"0Xff0044", true},
		{"abcdefg", false},
		{"-1", false},
	}

	for _, t := range testCases {
		actual := validator.IsHexadecimal(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsNumber(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"1", true},
		{"+1", false},
		{"-1", false},
		{"1.12", false},
		{"1.o", false},
	}

	for _, t := range testCases {
		actual := validator.IsNumber(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsNumeric(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"1", true},
		{"+1", true},
		{"-1", true},
		{"1.12", true},
		{"1.o", false},
	}

	for _, t := range testCases {
		actual := validator.IsNumeric(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsAlphaUnicodeNumeric(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"abc", true},
		{"this is a test string", false},
		{"这是一个测试字符串", true},
		{"\u0031\u0032\u0033", true}, // unicode 5
		{"123", true},
		{"<>@;.-=", false},
		{"ひらがな・カタカナ、．漢字", false},
		{"あいうえおfoobar", true},
		{"test＠example.com", false},
		{"1234abcDE", true},
		{"ｶﾀｶﾅ", true},
	}

	for _, t := range testCases {
		actual := validator.IsAlphaUnicodeNumeric(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsAlphaUnicode(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"abc", true},
		{"this is a test string", false},
		{"这是一个测试字符串", true},
		{"123", false},
		{"<>@;.-=", false},
		{"ひらがな・カタカナ、．漢字", false},
		{"あいうえおfoobar", true},
		{"test＠example.com", false},
		{"1234abcDE", false},
		{"ｶﾀｶﾅ", true},
	}

	for _, t := range testCases {
		actual := validator.IsAlphaUnicode(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"abc!23", false},
		{"5", true},
		{"abcd123", true},
	}

	for _, t := range testCases {
		actual := validator.IsAlphaNumeric(t.param)
		assert.Equal(t.expected, actual)
	}
}

func TestIsAlpha(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		{"", false},
		{"abcd", true},
		{"abc®", false},
		{ "abc÷", false},
		{"abc1", false},
		{"this is a test string", false},
		{"1", false},
	}

	for _, t := range testCases {
		actual := validator.IsAlpha(t.param)
		assert.Equal(t.expected, actual)
	}
}
