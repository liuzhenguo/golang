下载protoc https://github.com/google/protobuf/releases
下载  go get -u github.com/golang/protobuf/protoc-gen-go
创建test.proto
创建 test.pb.go 使用命令 protoc --go_out=. *.proto   
main文件中引用 test.pb.go的包
