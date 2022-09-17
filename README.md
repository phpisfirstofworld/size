# image

字节换算工具

## 安装
```shell
go get github.com/phpisfirstofworld/size
```

## 快速开始
```go
package main

import (
	"fmt"
	"github.com/phpisfirstofworld/size"
)

func main() {

	s := size.NewSize(56609)

	fmt.Println(s.Kb())
	fmt.Println(s.Mb())
	fmt.Println(s.Gb())

}
```