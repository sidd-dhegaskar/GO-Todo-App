package handlers

import(
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_api/internal/repository"
	"strconv"
)

type CreateTodoInput struct{
	Title string `json:"title" binding:"required"`
	Completed bool `json:"completed"`
}

type UpdateTodoInput struct{
	Title string `json:"title"`
	// &true ----> set completed as true
	// &false ----> set completed as false
	// &nil ----> did not set "completed" ---> default to false as per design
	Completed *bool `json:"completed"`
}

func CreateTodoHandler(pool *pgxpool.Pool) gin.HandlerFunc{
	return func(c *gin.Context){
		var input CreateTodoInput

		if err:= c.ShouldBindJSON(&input); err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		}

		todo, err:= repository.CreateTodo(pool, input.Title, input.Completed)

		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, todo)
	}
}

func GetAllTodosHandler(pool *pgxpool.Pool) gin.HandlerFunc{
	return func(c *gin.Context){
		todos,err :=repository.GetAllTodos(pool)

		if err !=nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, todos)
	}
}

func GetTodoByIDHandler(pool *pgxpool.Pool) gin.HandlerFunc{
	return func(c *gin.Context){
		idstr:= c.Param("id")
		id, err:= strconv.Atoi(idstr)
		if err!= nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid todo ID"})
			return
		}
		todo, err := repository.GetTodoByID(pool, id)
		if err != nil{
			if err==pgx.ErrNoRows{
				c.JSON(http.StatusNotFound, gin.H{"error":"Todo not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodoHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Todo Id"})
			return
		}

		var input UpdateTodoInput

		if err := c.ShouldBindJSON(&input); err != nil { //adding & so that we know where to write the values
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if input.Title == "" && input.Completed == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Provide minimum of 1 field"})
			return
		}

		var completed bool
		if input.Completed != nil {
			completed = *input.Completed
		}

		todo, err := repository.UpdateTodo(pool, id, input.Title, completed)

		if err != nil {
			if err == pgx.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todo)
	}
}
