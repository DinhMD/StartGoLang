package common

import "strconv"

func ToInt(str string) int {
	uintNum, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(uintNum)
}

func ToUInt(str string) uint {
	uintNum, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint(uintNum)
}
