package handlers

import(
	"github.com/jackc/pgx/v5/pgxpool"
	//"github.com/jackc/pgx/v5"
	"github.com/gin-gonic/gin"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"todo_api/internal/models"
	"todo_api/internal/repository"
	"todo_api/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type RegisterRequest struct{
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct{
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct{
	Token string `json:"token"`
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

func LoginHandler(pool *pgxpool.Pool, cfg *config.Config)gin.HandlerFunc{
	return func(c *gin.Context){
		var loginRequest LoginRequest

		if err:= c.BindJSON(&loginRequest); err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		user, err:= repository.GetUserByEmail(pool, loginRequest.Email)
		if err!=nil{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"User is not authorized"})
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(loginRequest.Password))
		if err!=nil{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"User is not authorized"})
			return
		}

		//map[string]any{}
		claims:=jwt.MapClaims{
			"user_id":user.ID,
			"email":user.Email,
			"exp":time.Now().Add(24*time.Hour).Unix(),
		}

		token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

		tokenString,err:=token.SignedString([]byte(cfg.JWTSecret))
		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to generate token"+err.Error()})
			return
		}

		c.JSON(http.StatusOK, LoginResponse{Token:tokenString})
	}
}