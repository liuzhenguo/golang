## 单元测试
* 1. 测试文件名必须是 *_test.go  ,  星号（×）要和测试的文件名一致。
* 2. 测试的函数必须以TestXxx或者Test_xxx，必须以Test开头。
* 3. 执行 go test对当前目录下的所有 *_test.go进行编译并自动执行测试。
* 4. 测试某个文件使用 go test -file *_test.go ，-file可以省略。
* 5. go test -run="Test_xxx or TestXxx"执行某个测试方法。
* 6.go test -v 全部执行，没有错误的pass有错误的报错，go test 如果遇到报错，打印出报错停止。


## 性能测试
* 1.测试文件必须要以 *_b_test.go开头 星号(×)号要和函数名称相同。
* 2.测试函数名必须要以Benchmark_xxx或BenchmarkXxx开头。
* 3.所有go文件的benchmark进行测试 go test -bench=".*" 或 go test . -bench=".*"
* 4.对某个go文件的benchmark进行测试 go test *_b_test.go -benhc=".*"进行测试。

## Memery&CPU
