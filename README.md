# go-rc6

`go-rc6` golang rc6 encrypt

```sh 
'   ██████╗  ██████╗       ██████╗  ██████╗ ██████╗ 
'  ██╔════╝ ██╔═══██╗      ██╔══██╗██╔════╝██╔════╝ 
'  ██║  ███╗██║   ██║█████╗██████╔╝██║     ███████╗ 
'  ██║   ██║██║   ██║╚════╝██╔══██╗██║     ██╔═══██╗
'  ╚██████╔╝╚██████╔╝      ██║  ██║╚██████╗╚██████╔╝
'   ╚═════╝  ╚═════╝       ╚═╝  ╚═╝ ╚═════╝ ╚═════╝ 
```


## Installation

Use [`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to install and update:

```sh
$ go get -u github.com/scloudrun/go-rc6
```

## Quick start
 
```sh
# assume the following codes in example.go file
$ cat example.go
```

```go
package main

import ( 
    "fmt"
    rc6lib "github.com/scloudrun/go-rc6"
)

var (
	defaultKey     = []byte("0000000000000000")
	defaultIv      = []byte("1111111111111111")
	defaultEncData = []byte("scloudrun")
	defaultEncByte = []byte{148, 250, 198, 116, 20, 46, 103, 128, 40, 160, 83, 239, 210, 246, 151, 15}
)

func main() {
    fmt.Println(rc6lib.Rc6Enc(defaultKey,defaultIv,defaultEncData,true))
}
```

```
# run example.go
$ go run example.go
```

## Todo
- extend
