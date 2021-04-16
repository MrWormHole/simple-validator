package validator

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	"strings"
)

func IsAlpha(str string) bool {
	return alphaRegex.MatchString(str)
}

func IsAlphaNumeric(str string) bool {
	return alphaNumericRegex.MatchString(str)
}

func IsAlphaUnicode(str string) bool {
	return alphaUnicodeRegex.MatchString(str)
}

func IsAlphaUnicodeNumeric(str string) bool {
	return alphaUnicodeNumericRegex.MatchString(str)
}

func IsNumeric(str string) bool {
	return numericRegex.MatchString(str)
}

func IsNumber(str string) bool {
	return numberRegex.MatchString(str)
}

func IsHexadecimal(str string) bool {
	return hexadecimalRegex.MatchString(str)
}

func IsHexcolor(str string) bool {
	return hexcolorRegex.MatchString(str)
}

func IsRGB(str string) bool {
	return rgbRegex.MatchString(str)
}

func IsRGBA(str string) bool {
	return rgbaRegex.MatchString(str)
}

func IsHSL(str string) bool {
	return hslRegex.MatchString(str)
}

func IsHSLA(str string) bool {
	return hslaRegex.MatchString(str)
}

func IsEmail(str string) bool {
	return emailRegex.MatchString(str)
}

func IsBase64(str string) bool {
	return base64Regex.MatchString(str)
}

func IsBase64URL(str string) bool {
	return base64URLRegex.MatchString(str)
}

func IsISBN10(str string) bool {
	return isbn10Regex.MatchString(str)
}

func IsISBN13(str string) bool {
	return isbn13Regex.MatchString(str)
}

func IsUUID3(str string) bool {
	return uuid3Regex.MatchString(str)
}

func IsUUID4(str string) bool {
	return uuid4Regex.MatchString(str)
}

func IsUUID5(str string) bool {
	return uuid5Regex.MatchString(str)
}

func IsUUID(str string) bool {
	return uuidRegex.MatchString(str)
}

func IsUUID3Mixed(str string) bool {
	return uuid3MixedRegex.MatchString(str)
}

func IsUUID4Mixed(str string) bool {
	return uuid4MixedRegex.MatchString(str)
}

func IsUUID5Mixed(str string) bool {
	return uuid5MixedRegex.MatchString(str)
}

func IsUUIDMixed(str string) bool {
	return uuidMixedRegex.MatchString(str)
}

func IsASCII(str string) bool {
	return asciiRegex.MatchString(str)
}

func IsPrintableASCII(str string) bool {
	return printableASCIIRegex.MatchString(str)
}

func IsMultibyte(str string) bool {
	return multibyteRegex.MatchString(str)
}

func IsDataURI(str string) bool {
	uri := strings.SplitN(str, ",", 2)

	if len(uri) != 2 {
		return false
	}

	if !dataURIRegex.MatchString(uri[0]) {
		return false
	}

	return base64Regex.MatchString(uri[1])
}

func IsLatitude(str string) bool {
	return latitudeRegex.MatchString(str)
}

func IsLongitude(str string) bool {
	return longitudeRegex.MatchString(str)
}

func IsIPAddress(str string) bool {
	return ipAddressRegex.MatchString(str)
}

func IsDomainName(str string) bool {
	domain := strings.Split(str, ".")
	if domain[len(domain) - 1] == "" {
		return false
	}

	return domainNameRegex.MatchString(str)
}

func IsETHAddress(str string) bool {
	if !ethAddressRegex.MatchString(str) {
		return false
	}

	if isETHAddressLower(str) || isETHAddressUpper(str) {
		return true
	}

	address := str[2:]
	h := sha3.NewLegacyKeccak256()
	_, _ = h.Write([]byte(strings.ToLower(address)))
	hash := hex.EncodeToString(h.Sum(nil))

	for i := 0; i < len(address); i++ {
		if address[i] <= '9' {
			continue
		}
		if hash[i] > '7' && address[i] >= 'a' || hash[i] <= '7' && address[i] <= 'F' {
			return false
		}
	}

	return true
}

func isETHAddressLower(str string) bool {
	return ethAddressRegexLower.MatchString(str)
}

func isETHAddressUpper(str string) bool {
	return ethAddressRegexUpper.MatchString(str)
}

func IsURLEncoded(str string) bool {
	return urlEncodedRegex.MatchString(str)
}

func IsHTMLEncoded(str string) bool {
	return htmlEncodedRegex.MatchString(str)
}

func IsHTML(str string) bool {
	return htmlRegex.MatchString(str)
}

func IsEmpty(str string) bool {
	return strings.Trim(str, " ") == ""
}