# Simpleprint

一个简化Go打印操作的小型库，无需每次都使用`fmt.`前缀。

## 安装

```bash
go get github.com/cody1686/simpleprint
```

## 引入与使用步骤

### 1. 在您的项目中添加依赖

在项目的`go.mod`文件中添加依赖：

```bash
# 如果您的项目还没有初始化go模块
go mod init 您的项目名称

# 添加simpleprint依赖
go get github.com/cody1686/simpleprint
```

### 2. 在代码中引入

使用点导入语法在代码中引入，这样可以直接使用打印函数而无需包前缀：

```go
package main

import . "github.com/cody1686/simpleprint"

func main() {
    // 直接使用，无需包前缀
    Println("Hello world!")
    Printf("Hello %s\n", "Go")
}
```

### 3. 使用示例

```go
package main

import . "github.com/cody1686/simpleprint"

func main() {
    // 基本打印
    Println("直接打印，无需fmt前缀")
    
    // 格式化打印
    name := "Go"
    Printf("Hello %s!\n", name)
    
    // 打印多个值
    Println("多个值:", 42, true, 3.14)
    
    // 返回格式化字符串
    message := Sprintf("格式化字符串: %d, %t", 100, false)
    Println(message)
}
```

## 包含的函数

- `Print` - 打印内容，不换行
- `Println` - 打印内容并换行
- `Printf` - 格式化打印
- `Sprintf` - 返回格式化的字符串
- `Fprintf` - 写入格式化内容到io.Writer 
