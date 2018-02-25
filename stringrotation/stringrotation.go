package stringrotation

import (
	"strings"
)

func IsRotation(s1, s2 string) bool {
	lens1 := len(s1)
	if lens1 != len(s2) || lens1 == 0 {
		return false
	}
	s1 = s1 + s1
	return strings.Contains(s1, s2)
}

func IsRotationFF(s1, s2 string) bool {
	lens1 := len(s1)
	if lens1 != len(s2) || lens1 == 0 {
		return false
	}
	if s1 == s2 {
		return true
	}
	for i := 0; i < len(s1); i++ {
		s1 = s1[1:] + string(s1[0])
		if s1 == s2 {
			return true
		}
	}
	return false
}
