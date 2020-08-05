
### 实验
    go build main.go
     go run main.go


### 打包
在windows下编译成Linux下可执行的二进制文件并且执行(cmd下)
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go

### 部署到linux服务器
赋予权限
    chmod 777 main
执行
    ./main
或者后台执行
    nohup ./main &
杀死进程
    killall ./main

