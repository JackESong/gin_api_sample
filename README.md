
### 实验
go build main.go
go run main.go


文章
获取指定文章：GET http://localhost:8000/api/v1/articles/1
获取全部文章：GET http://localhost:8000/api/v1/articles
新增文章：POST http://localhost:8000/api/v1/articles，表单可用 form-data 或 x-www-form-urlencoded 形式
修改文章：PUT http://localhost:8000/api/v1/articles/1
删除文章：DELETE http://localhost:8000/api/v1/articles/1
标签
获取指定标签：GET http://localhost:8000/api/v1/tags/1
获取全部标签：GET http://localhost:8000/api/v1/tags
新增标签：POST http://localhost:8000/api/v1/tags
修改标签：PUT http://localhost:8000/api/v1/tags/1
删除标签：DELETE http://localhost:8000/api/v1/tags/1

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

