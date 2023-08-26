package decrypt

func Nimbus(str string) string {

	decryptedstr := ""

	for _,c := range str{
		ascii := int(c)
		char := string(ascii - 3)
		decryptedstr += char
		
	}

	return decryptedstr
}