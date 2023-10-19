package encrypt

func Nimbus(str string) string {
	encryptedStr := ""
	for _, ch := range str {
		asciiCode := int(ch)
		character := string(asciiCode + 3)
		encryptedStr += character
	}
	return encryptedStr
}
