# go 调用外部动态库


1. 执行脚本生成动态共享库

2. 将共享库拷贝到可以被外部程序访问的地方

示例程序：

```go

package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L. -ladd
#include <libadd.h>
*/
import "C"
import "fmt"

func main() {
    z := C.add(1, 3)
    fmt.Println(z)
}

// CFLAGS中的-I（大写的i） 参数表示.h头文件所在的路径
// LDFLAGS中的-L(大写) 表示.so或者 dylib 文件所在的路径 -l(小写的L) 表示指定该路径下的库名称,省略了前缀lib以及后缀.so .dylib

```


## 2. Golang的构建模式

    -buildmode=archive
		Build the listed non-main packages into .a files. Packages named
		main are ignored.

	-buildmode=c-archive
		Build the listed main package, plus all packages it imports,
		into a C archive file. The only callable symbols will be those
		functions exported using a cgo //export comment. Requires
		exactly one main package to be listed.

	-buildmode=c-shared
		Build the listed main package, plus all packages it imports,
		into a C shared library. The only callable symbols will
		be those functions exported using a cgo //export comment.
		Requires exactly one main package to be listed.

	-buildmode=default
		Listed main packages are built into executables and listed
		non-main packages are built into .a files (the default
		behavior).

	-buildmode=shared
		Combine all the listed non-main packages into a single shared
		library that will be used when building with the -linkshared
		option. Packages named main are ignored.

	-buildmode=exe
		Build the listed main packages and everything they import into
		executables. Packages not named main are ignored.

	-buildmode=pie
		Build the listed main packages and everything they import into
		position independent executables (PIE). Packages not named
		main are ignored.

	-buildmode=plugin
		Build the listed main packages, plus all packages that they
		import, into a Go plugin. Packages not named main are ignored'
