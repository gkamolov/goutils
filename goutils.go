package goutils

func NumberToWord(a int) string {
	var numbers = []string{"нол", "бир", "икки", "уч бурчак", "тўрт", "беш", "олти", "етти", "саккиз", "тўққиз"}

	if a >= 0 && a < 10 {
		return numbers[a]
	} else {
		return "Бу сон 1 ва 10 оралиғида эмас!"
	}
}
