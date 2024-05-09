package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/penguinn/go-sdk/log"
	"github.com/penguinn/goframe/api/controller"
	"github.com/penguinn/goframe/api/middleware"
	"github.com/penguinn/goframe/config"
	"github.com/penguinn/goframe/dao"
	"github.com/penguinn/goframe/service/scanner"
)

var done = make(chan bool, 1)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatalf("invalid application configuration: %s", err)
	}

	err = dao.Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("successfully connected to database")

	// 启动scanner服务，可注释取消
	err = scanner.Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("successfully start scanner")

	// 暴露http服务，可注释取消
	serve()
}

func serve() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())
	v1 := r.Group("/api/v1")
	{
		v1.POST("/example", controller.ExampleController{}.Create)
	}
	err := r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
	if err != nil {
		log.Fatal(err)
	}
}

// 捕获sigterm(15)信号量放入channel，可以利用channel作销毁前的资源回收
func gracePeriod() <-chan bool {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Infof("capture signal %v", sig)
		done <- true
	}()

	return done
}
