go get github.com/lib/pq
go get github.com/gorilla/mux
go build server.go database_interface.go MyDatabase.go PostgresDatabase.go
server.exe -d=true
pause