# Simpleprint

一个简化Go打印操作的小型库，无需每次都使用`fmt.`前缀。

## 安装

```bash
go get github.com/yourusername/simpleprint
```

## 使用方式

使用点导入可以直接调用打印函数，无需包前缀：

```go
package main

import . "github.com/yourusername/simpleprint"

func main() {
    // 直接使用，无需包前缀
    Println("Hello world!")
    Printf("Hello %s\n", "Go")
}
```

## 包含的函数

- `Print`
- `Println`
- `Printf`
- `Sprintf`
- `Fprintf` 