# dsl
Service definition DSL. The DSL focuses on describing either what the service provides or what the service expects of it's environment.

Most functions are passthrough so you can use the NewX in the Add|SetX method and assign it to a variable and work with it. Ie:

```go
myEndpoint := myService.AddEndpoint(GET("/my/path"))
myEndpoint.Name = "myEndpoint"
// ...
```
