package main

import(
	"github.com/gin-gonic/gin"
	"log"
	"todo_api/internal/config"
	"todo_api/internal/database"
	"github.com/jackc/pgx/v5/pgxpool"
	"todo_api/internal/handlers"

)

func main(){
	var cfg *config.Config
	var err error
	cfg,err = config.Load()
	
	if err!=nil{
		log.Fatal("Unable to load config:", err)
	}
	
	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)

	if err!=nil{
		log.Fatal("Unable to connect to Db:", err)
	}

	defer pool.Close()

	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			//map[string]interface -> map[string]any
			"message":"Todo API is running",
			"status":"success",
			"databse": "connected",
		})
	})
	router.POST("/todos", handlers.CreateTodoHandler(pool))
	router.GET("/todos", handlers.GetAllTodoshandler(pool))
	router.Run(":"+ cfg.Port)
}