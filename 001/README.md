# Golem
Railsライクな感じでDB生成やモデル生成を手伝う的なモノを作成してます。

# Commands
`interact テーブル名 レコード名:型:オプション`

`foundation User Test:string:not_null`

↓

```
type User struct{
  gorm.Model
  Test: string `gorm:"not null"`
}
```

## TODO

[ ] 上記のテストの not null の前後にあるダブルクオーテーションが文字コードが初期化される。
    - おそらくテンプレートにExecuteする際にバッククオートが悪く動作してると思われる。

`go run .\main.go foundation User Test:string:not_null`

↓

``` User.go
package models

import (
	"time"
	"github.com/jinzhu/gorm"
	// MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
// TODO ドライバの指定の仕方は考える必要がある。

type User struct {
	gorm.Model
		Test string `gorm:&#34;not null&#34;`

}
```