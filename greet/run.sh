protoc -Iproto --go_opt=module=github.com/truong11t2/learn-go/greet --go_out=. --go-grpc_opt=module=github.com/truong11t2/learn-go/greet --go-grpc_out=. .\proto\*.proto

go run .\server\ *.go
go run .\client\ *.go

git init
git add greet/
git remote add origin https://github.com/truong11t2/learn-grpc