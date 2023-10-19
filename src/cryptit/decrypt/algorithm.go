// decrypt package consists of all decryption algorithms
package decrypt

// decrypts by reducing the ascii code by 3 for each character of input string
func Nimbus(str string) string {
	decryptedStr := ""
	for _, ch := range str {
		asciiCode := int(ch)
		character := string(asciiCode - 3)
		decryptedStr += character
	}
	return decryptedStr
}
