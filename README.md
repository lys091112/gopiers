# gopiers



###测试

- go tool cover 用来显示help， 覆盖率测试的命令
- go test -coverprofile=c.out  用来生成覆盖率测试文件到c.out
- go tool cover -html=c.out -o coverage.html 将c.out转化为html， 可以通过浏览器进行查看
- go test -bench=. 用来统计以Benchmark开头的测试方法的性能

白盒测试的一种技巧： 针对于方法测试中，包含有邮件发送，webhook发送等第三方依赖的代码，为了保证测试不依赖这些，可以将这些包装成一个接口，在测试的时候，传递的一个空实现，不修改现有代码逻辑，还可以屏蔽第三方的依赖影响
