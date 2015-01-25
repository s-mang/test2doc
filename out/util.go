package out

import (
	"log"
	"strings"
)

func CommaJoinStrs(args ...interface{}) string {
	var strList []string

	for _, arg := range args {
		strs, ok := arg.([]string)
		if ok {
			for _, str := range strs {
				strList = append(strList, str)
			}
		} else {
			log.Println("Error: CommaJoinStrs called with non []string argument.")
		}

	}

	return strings.Join(strList, ", ")
}
