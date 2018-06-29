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

const bufSize = 1024

// FileRead ファイル読み込み用
func FileRead(file *os.File) string {
	var cash string
	buf := make([]byte, bufSize)
	for {
		n, err := file.Read(buf)
		// nにbyteが入る
		// 初回は13バイト 2回目は0バイトだった
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		cash += string(buf[:n])
	}
	return cash

}
