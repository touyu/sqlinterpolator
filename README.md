# sqlinterpolator

## Examples
```golang
func main() {
    query := "select * from users where id < ?"
    args := []interface{}{6}

    res, _ := sqlinterpolator.Interpolate(query, args)
    fmt.Println(res)  // select * from users where id < 6
}
```
