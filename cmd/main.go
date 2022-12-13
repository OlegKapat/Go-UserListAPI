package main

import (
	//"bytes"
	"fmt"
	//"os"
	//"os/exec"

	"github.com/OlegKapat/ListOfWorks"
	"github.com/OlegKapat/ListOfWorks/pkg/handler"
	"github.com/OlegKapat/ListOfWorks/pkg/repository"
	"github.com/OlegKapat/ListOfWorks/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	var err error
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())

	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewSqlDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		DBtype:   viper.GetString("db.dbtype"),
	})
	if err != nil {
		logrus.Fatalf("fatal to initializ db", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	//createPgDb()
	//db.AutoMigrate(&todo.User{}, &todo.ListsItems{}, &todo.TodoItem{}, &todo.TodoList{}, &todo.UsersList{})

	//Run server
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatal("error to run http server: %s" + err.Error())
		panic(err)
	}
	fmt.Println("Server running")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

// Create DB
func createPgDb() {
	// cmd := exec.Command("User_list", "-p", "1433", "-h", "127.0.0.1", "-U", "superuser", "-e", "User_list")
	// var out bytes.Buffer
	// cmd.Stdout = &out
	// if err := cmd.Run(); err != nil {
	// 	logrus.Printf("Error: %v", err)
	// }
	// logrus.Printf("Output: %q\n", out.String())
}
