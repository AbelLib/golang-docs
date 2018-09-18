# GOPATH环境变量
按照约定 GOPATH 下面需要建立 3 个目录：
* bin 存放编译后生成的可执行文件
* pkg 存放编译后生成的包文件
* src 存放项目源码文件

对于Unix和Linux默认环境变量在`~/go`下。使用`go env`:
```shell
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/abel/Library/Caches/go-build"
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/abel/go"
GORACE=""
GOROOT="/usr/local/Cellar/go/1.10.3/libexec"
GOTMPDIR=""
GOTOOLDIR="/usr/local/Cellar/go/1.10.3/libexec/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/97/4v707d557zq_8sxtsdhps5xh0000gp/T/go-build422873217=/tmp/go-build -gno-record-gcc-switches -fno-common
```

