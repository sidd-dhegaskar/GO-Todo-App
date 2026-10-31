package handlers

import(
	"github.com/jackc/pgx/v5/pgxpool"
	//"github.com/jackc/pgx/v5"
	"github.com/gin-gonic/gin"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"todo_api/internal/models"
	"todo_api/internal/repository"
)

type RegisterRequest struct{
	Email string `json:"title" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUserHandler(pool *pgxpool.Pool) gin.HandlerFunc{
	return func(c *gin.Context){
		var registerRequest RegisterRequest

		if err := c.BindJSON(&registerRequest); err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		if len(registerRequest.Password)<6{
			c.JSON(http.StatusBadRequest, gin.H{"error":"Password must be 6 characters long"})
			return
		}

		hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to hash password"+err.Error()})
			return
		}
		 
		user:= &models.User{
			Email: registerRequest.Email,
			Password: string(hashedPassword),
		}

		createdUser,err :=repository.CreateUser(pool, user)
		if err!=nil{
			if err.Error()!=""{
				c.JSON(http.StatusBadRequest, gin.H{"error":"Email already registered"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusCreated, createdUser)
	}
}