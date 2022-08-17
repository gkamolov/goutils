package goutils

// Function to get max of two integers
func Getmax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func NumberToWord(a int) string {
	var numbers = []string{"нол", "бир", "икки", "уч", "тўрт", "беш", "олти", "етти", "саккиз", "тўққиз"}

	if a >= 0 && a < 10 {
		return numbers[a]
	} else {
		return "Бу сон 1 ва 10 оралиғида эмас!"
	}
}
