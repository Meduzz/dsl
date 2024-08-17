# dsl
Service definition DSL

Most functions are passthrough so you can use the NewX in the Add|SetX method and assign it to a variable and work with it. Ie:

```go
myEndpoint := myService.AddEndpoint(GET("/my/path"))
myEndpoint.Name = "myEndpoint"
// ...
```
