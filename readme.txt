go run cmd/main.go

go mod init github.com/OlegKapat/ListOfWorks
git init
go get -u github.com/gin-gonic/gin
go get github.com/spf13/viper  
go get github.com/golang-migrate/migrate/v4
go get github.com/sirupsen/logrus
go get github.com/gorilla/mux
go get github.com/denisenkom/go-mssqldb
or
 go get -u github.com/go-sql-driver/mysql  for mysql
 or
go get github.com/jinzhu/gorm  -- for mssql
scoop install migrate
migrate create  -ext sql -dir ./schema -seq init
migrate -path ./schema -database "mysql://root@tcp(localhost:3306)/postgress" -verbose up
go get github.com/joho/godotenv
go get github.com/dgrijalva/jwt-go
