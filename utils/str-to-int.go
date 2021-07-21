package utils

import "strconv"


func StrToInt (strList []string) []int {
	intList := make([]int, len(strList))
	for i, strElem := range strList {
		intElem, _ := strconv.Atoi(strElem)
		intList[i] = intElem
	}
	return intList
}
