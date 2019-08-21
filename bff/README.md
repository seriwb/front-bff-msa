```
go mod init
go run github.com/99designs/gqlgen init
go run github.com/99designs/gqlgen -v
rm resolver.go
go run github.com/99designs/gqlgen
go run server/server.go
```