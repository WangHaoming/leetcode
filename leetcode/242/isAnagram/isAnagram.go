package isanagram

func isAnagram(s string, t string) bool {
	alphabets := make([]int, 26)
	sbytes := []byte(s)
	tbytes := []byte(t)
	if len(sbytes) != len(tbytes) {
		return false
	}
	for i := 0; i < len(sbytes); i++ {
		alphabets[sbytes[i]-'a']++
		alphabets[tbytes[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if alphabets[i] != 0 {
			return false
		}
	}

	return true
}

func Test() {
	s := "car"
	t := "rec"
	println(isAnagram(s, t))
}
