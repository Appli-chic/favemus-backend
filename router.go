package main

import (
	"github.com/Favemus/config"
	"github.com/Favemus/controller"
	"github.com/Favemus/model"
	"github.com/Favemus/service"
	"github.com/Favemus/validator"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Creates all the routes
func InitRouter() *gin.Engine {
	identityKey := "id"
	authService := service.NewAuthService()
	userController := controller.NewUserController()
	r := gin.Default()

	// Add midllewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Realm",
		Key:         []byte(config.Conf.JwtSecret),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			u := &model.User{}
			u.ID = uint(claims[identityKey].(float64))
			return u
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals validator.AuthValidator
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userEmail := loginVals.Email
			password := loginVals.Password

			user, err := authService.GetUserByEmailAndPassword(userEmail, password)

			if err == nil {
				return user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*model.User); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// Add the routes
	r.POST("/auth/login", authMiddleware.LoginHandler)
	auth := r.Group("/")
	auth.GET("/auth/refresh_token", authMiddleware.RefreshHandler)

	// Protected routes
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		api := auth.Group("/api")
		{
			api.GET("/users/:id", userController.GetUser)
		}
	}

	return r
}
