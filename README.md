# TRI

time + random id

This is useful when you need an id, that is not likely to collide with other ids while being generated at the same time, but also does not need to be completely random.

```Go
id := tri.New(4)
fmt.Println(id)
// output: hldwZPCfkqk
```

```Go
id := tri.Time()
fmt.Println(id)
// output: nmVwZA
```



