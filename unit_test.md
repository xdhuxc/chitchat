### 运行方式

#### 测试单个文件
当只想测试 utils_test.go 文件时，使用命令：
```markdown
go test -v -cover=true ./src/utils/utils_test.go ./src/utils/utils.go
```
执行结果如下：
```
wanghuans-MacBook-Pro:chitchat wanghuan$ go test -v -cover=true ./src/utils/utils_test.go ./src/utils/utils.go
=== RUN   TestSuccessStringInSlice
--- PASS: TestSuccessStringInSlice (0.00s)
    utils_test.go:7: pass
=== RUN   TestFailedStringInSlice
--- PASS: TestFailedStringInSlice (0.00s)
    utils_test.go:17: pass
PASS
coverage: 100.0% of statements
ok      command-line-arguments  0.005s  coverage: 100.0% of statements
```

#### 测试整个 utils 包
测试整个 utils 包时，使用命令为：
```markdown
go test -v -cover=true ./src/utils/...
```
执行结果如下：
```markdown
wanghuans-MacBook-Pro:chitchat wanghuan$ go test -v -cover=true ./src/utils/...
=== RUN   TestSuccessStringInSlice
--- PASS: TestSuccessStringInSlice (0.00s)
    utils_test.go:7: pass
=== RUN   TestFailedStringInSlice
--- PASS: TestFailedStringInSlice (0.00s)
    utils_test.go:17: pass
PASS
coverage: 100.0% of statements
ok      _/Users/wanghuan/GolandProjects/chitchat/src/utils      0.007s  coverage: 100.0% of statements
```

#### 测试单个测试用例
测试单个测试用例时，使用命令为：
```markdown
go test -v -cover=true ./src/utils -run TestSuccessStringInSlice
```
执行结果为：
```markdown
wanghuans-MacBook-Pro:chitchat wanghuan$ go test -v -cover=true ./src/utils -run TestSuccessStringInSlice
=== RUN   TestSuccessStringInSlice
--- PASS: TestSuccessStringInSlice (0.00s)
    utils_test.go:7: pass
PASS
coverage: 75.0% of statements
ok      _/Users/wanghuan/GolandProjects/chitchat/src/utils      0.005s  coverage: 75.0% of statements
```

### gotests 的使用
1、安装 gotests
```markdown
go get -u github.com/cweill/gotests/...
```
2、运行如下命令生成测试代码
```markdown
gotests -all -w ./src/utils/utils.go ./src/utils/utils_test.go
```
在生成的 utils_test.go 中，添加如下代码初始化测试数据：
```markdown
{
    // TODO: Add test cases.
    {
        name : "success",
        args : args{
            s: "a",
            slice: []string{"a", "b", "c"},
        },
        want: true,
    },
    {
        name : "failed",
        args : args{
            s: "d",
            slice: []string{"a", "b", "c"},
        },
        want: false,
    },
}
```
执行如下命令进行测试：
```markdown
wanghuans-MacBook-Pro:chitchat wanghuan$ go test -v -cover=true ./src/utils -run TestStringInSlice
=== RUN   TestStringInSlice
=== RUN   TestStringInSlice/success
=== RUN   TestStringInSlice/failed
--- PASS: TestStringInSlice (0.00s)
    --- PASS: TestStringInSlice/success (0.00s)
    --- PASS: TestStringInSlice/failed (0.00s)
PASS
coverage: 40.0% of statements
ok      _/Users/wanghuan/GolandProjects/chitchat/src/utils      0.005s  coverage: 40.0% of statements
```
