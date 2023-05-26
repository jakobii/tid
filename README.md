# TRI

time + random id

This is useful when you need an id, that is not likely to collide with other ids, but also does not need to be completely random.

```Go
id := tri.New(4) 
fmt.Println(id)
// time + 4 bytes of randomness.
// output: hldwZPCfkqk
```

```Go
id := tri.Time()
fmt.Println(id)
// time with zero randomness.
// output: hldwZA
```



