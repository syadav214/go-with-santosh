mkdir -p src/github.com/syadav214/todo
go mod init github.com/syadav214/todo
go get .
go get -u github.com/go-pg/pg

go build .