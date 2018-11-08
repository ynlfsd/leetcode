package codes

import "strings"

type T struct {
	ch byte
	t *T
}
func IsMatch(s string, p string) bool {

	if strings.Contains(p,".*"){return true}




	return true
}