package testdoc

import (
	"bytes"
	"encoding/json"
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

func IndentJSONBody(args ...interface{}) string {
	if len(args) != 1 {
		panic("IndentJSONBody: called with too many arguments.")
	}

	bodyStr, ok := args[0].(string)
	if !ok {
		panic("IndentJSONBody: argument should be a string.")
	}

	var outJSON bytes.Buffer
	json.Indent(&outJSON, []byte(bodyStr), "\t\t\t", "\t")

	return string(outJSON.Bytes())
}
