# Golem
Railsライクな感じでDB生成やモデル生成を手伝う的なモノを作成してます。

# Commands
`interact テーブル名 レコード名 型`
すべてスネークケースで

`interact Test Value string`

↓

```
type Model struct{
  gorm.Model
  Table: int
}
```