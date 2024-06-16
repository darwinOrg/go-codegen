go build -ldflags "-s -w" -o build/go-codegen-default cmd/default/main.go
go-codegen-default -i 输入sql文件路径 -p 包路径前缀 -o 输出文件目录路径