# while
带出口的死循环

# 调用示例

```go
//初始化循环对象，100为最大循环数量
while := NewWhile(100)
 
//todo 业务代码
//模拟业务代码，a初始为100，循环减一，当减到90时退出循环
a := 100
while.For(func() {
   a--
   if a == 90 {
      while.Break()
   }
})
fmt.Println(a)
```
