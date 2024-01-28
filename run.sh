protoc -Iproto --go_opt=module=github.com/truong11t2/learn-go/greet --go_out=. --go-grpc_opt=module=github.com/truong11t2/learn-go/greet --go-grpc_out=. .\proto\*.proto

go run .\server\ *.go
go run .\client\ *.go

git init
git add greet/
git remote add origin https://github.com/truong11t2/learn-grpc
git pull origin main --allow-unrelated-histories

git commit -m "message"
git push origin main

protoc -Iproto --go_opt=module=github.com/truong11t2/learn-go/calculator --go_out=. --go-grpc_opt=module=github.com/truong11t2/learn-go/calculator --go-grpc_out=. .\proto\*.proto
go mod init github.com/truong11t2/learn-go/calculator
go mod tidy