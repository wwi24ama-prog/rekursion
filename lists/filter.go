package lists

// Liefert eine Liste mit allen Elementen aus list, die kleiner oder gleich key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterLess(list []int, key int) []int {
	lesslist := []int{}

	FilterLessHelper(list, key, &lesslist)

	return lesslist
}

func FilterLessHelper(list []int, key int, lesslist *[]int) {
	if Empty(list) {
		return
	}
	if list[0] <= key {
		*lesslist = append(*lesslist, list[0])
	}

	FilterLessHelper(list[1:], key, lesslist)
}

// Liefert eine Liste mit allen Elementen aus list, die echt größer als key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterGreater(list []int, key int) []int {
	// Gehen Sie analog zu FilterLess vor.

	// TODO
	return []int{}
}
