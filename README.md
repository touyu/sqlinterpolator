# sqlinterpolator

## Examples
```golang
func main() {
    query := "select * from users where id = ?;"
    args := []interface{}{6}

    res, _ := sqlinterpolator.Interpolate(query, args)
    fmt.Println(res)  // select * from users where id = 6;
}
```

### SQLBoiler
```golang
queryObject := models.NewQuery(From("users"), models.UserWhere.ID.EQ(6))
query, args := queries.BuildQuery(queryObject)
res, _ := sqlinterpolator.Interpolate(raw, args)
fmt.Println(res)  // select * from users where id < 6;
```
