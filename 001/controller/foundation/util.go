package foundation

import (
	"fmt"
	"os"
)

func isPanic(work error, message string) {
	if work != nil {
		fmt.Println(message)
		panic(work)
	}
}

// IsExist 存在チェック(なかったらtrue)
func IsExist(table string) bool {
	table = "./model/" + table + ".go"
	_, err := os.Stat(table)
	if err != nil {
		return false
	}
	return true
}
