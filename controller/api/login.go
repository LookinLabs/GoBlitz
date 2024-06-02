package api

import (
	"net/http"
	"os"
	helper "web/helpers"
	"web/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type loginFormData struct {
	Email    string `form:"email"`
	Username string `form:"username"`
	Password string `form:"password"`
}

func Signup(c *gin.Context) {
	var data loginFormData
	if err := c.Bind(&data); err != nil {
		c.Render(http.StatusBadRequest, render.Data{})
		return
	}

	if model.CheckUserExistance(data.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	user := model.CreateUser(data.Email, data.Username, data.Password)
	if user == nil || user.ID == 0 {
		c.Render(http.StatusBadRequest, render.Data{})
		return
	}

	err := helper.SetSession(c, user.ID)
	if err != nil {
		return
	}
	c.Redirect(http.StatusFound, os.Getenv("API_PATH")+"ping")
}

func Signin(c *gin.Context) {
	var data loginFormData
	if err := c.Bind(&data); err != nil {
		c.Render(http.StatusBadRequest, render.Data{})
		return
	}

	if !model.CheckUserExistance(data.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exists"})
		return
	}

	user := model.CheckPasswordMatch(data.Username, data.Password)
	if user.ID == 0 {
		c.Render(http.StatusUnauthorized, render.Data{})
		return
	}

	err := helper.SetSession(c, user.ID)
	if err != nil {
		return
	}
	c.Redirect(http.StatusFound, os.Getenv("API_PATH")+"ping")
}

func Signout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.Redirect(http.StatusFound, "/")
}
