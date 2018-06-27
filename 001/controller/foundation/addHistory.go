package foundation

import (
	"fmt"
	"os"
)

func AddHistroy(table string) {
	file, err := os.OpenFile("./db/history.dat", os.O_CREATE|os.O_APPEND, 0666)
	isPanic(err, "histroyあたりのエラー")
	defer file.Close()

	fmt.Fprintln(file, table)
}
