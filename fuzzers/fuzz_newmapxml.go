package myfuzz

import "github.com/clbanning/mxj"

func Fuzz(data []byte) int {
	_, err := mxj.NewMapXml(data)
	if err != nil { return 1 }
	return 0
}
