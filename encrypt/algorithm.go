package encrypt 

func Nimbus(str string) string {
	encryptedstr := ""

	for _, c := range str{

		ascii := int(c)
		char := string(ascii + 3)

		encryptedstr += char
	}
	return encryptedstr
	
}

func Name() int {
	
	return 2
}