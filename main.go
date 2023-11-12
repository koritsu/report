package main

import (
	"fmt"
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"restapi-go/api"
	"restapi-go/base"
	"restapi-go/config"
	"restapi-go/database"
	_ "restapi-go/docs"
	"restapi-go/logging"
	"time"
)

const (
	GeneralErrorExitCode = 1
)

// @title Rest API server written by GO
// @version 1.0
// @host localhost:8080
func main() {

	var err error

	err = config.LoadConfig()
	if err != nil {
		fmt.Printf("Config error: %s", err)
		os.Exit(GeneralErrorExitCode)
	}

	appConfig := config.Config()
	err = logging.InitLog(appConfig)
	if err != nil {
		fmt.Printf("Log error: %s", err)
		os.Exit(GeneralErrorExitCode)
	}

	// init DB
	err = database.InitDB()
	if err != nil {
		fmt.Printf("DB error: %s", err)
		os.Exit(GeneralErrorExitCode)
	}

	// savePID
	base.SavePID(appConfig)

	// init gin
	router := initGin()

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := ":8080"
	if err = router.Run(port); err != nil {
		panic(err)
	}

	fmt.Printf("Server is starting on %s\n", port)
}

func initGin() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// middleware
	// logger?
	gin.DefaultWriter = logging.Log().Logger
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: logging.Log().Logger,
	}))

	memoryStore := persist.NewMemoryStore(1 * time.Minute)

	// route
	router.POST("/create", api.Create)
	router.GET("/read/:id", cache.CacheByRequestURI(memoryStore, 2*time.Second), api.Read)
	router.GET("/read/:id/file", api.ReadFile)
	router.PUT("/update/:id", api.Update)
	router.DELETE("/delete/:id", api.Delete)

	return router
}
