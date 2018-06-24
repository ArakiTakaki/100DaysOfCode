# Golem
Railsライクな感じでDB生成やモデル生成を手伝う的なモノを作成してます。

# Commands
`interact テーブル名 レコード名 型`

`interact Test Value string`

↓

```
type Test struct{
  gorm.Model
  Value: string
}
```