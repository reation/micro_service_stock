package tool

import (
	"strconv"
)

func IntToStringArr(idList []int64) []string {
	var resp = make([]string, len(idList))
	for k, val := range idList {
		resp[k] = strconv.FormatInt(val, 10)
	}
	return resp
}
