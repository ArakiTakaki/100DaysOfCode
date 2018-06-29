package foundation

import (
	"fmt"
	"os"
)

// AddHistory ADD
func AddHistroy(table string) {
	file, err := os.OpenFile("./db/history.dat", os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	isPanic(err, "histroyあたりのエラー")
	fmt.Fprintln(file, table)
}
