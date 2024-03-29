// Import necessary packages
package main

import (
	"Say-Hi/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	initDB()
	defer DB.Close()
	initHandlers()
	r := gin.New()
	api := r.Group("/api")

	user := api.Group("/user")

	user.POST("/register", handler.User.RegisterHandler.Register)
	user.POST("/verify-email", handler.User.VerifyEmailHandler.VerifyEmail)
	user.POST("/login", handler.User.LoginHandler.Login)

	r.Use(auth.Middleware())
	user.POST("/logout", handler.User.LogoutHandler.Logout)
	user.POST("/forgot-password", handler.User.ForgotPasswordHandler.ForgotPassword)

	notification := api.Group("/notification")
	notification.POST("send-email", handler.Notification.SendEmailHandler.SendEmail)
	
	http.HandleFunc("/send-message", handler.MessageHandler.SendMessage)
	r.Run("localhost:8080")
}
