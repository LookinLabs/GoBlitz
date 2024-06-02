package api

import (
	"net/http"
	"os"

	helper "web/helpers"
	sql "web/repository/db"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type loginFormData struct {
	Email    string `form:"email"`
	Username string `form:"username"`
	Password string `form:"password"`
}

func SignUp(ctx *gin.Context) {
	var data loginFormData
	if err := ctx.Bind(&data); err != nil {
		ctx.Render(http.StatusBadRequest, render.Data{})
		return
	}

	checks := []helper.Check{
		helper.NewCheck(checkUsernameExists(data.Username), "User already exists", false),
		helper.NewCheck(checkEmailExists(data.Email), "Email already exists", false),
	}

	if helper.CheckAndRespond(ctx, checks) {
		return
	}

	user := sql.CreateUser(data.Email, data.Username, data.Password)
	if user == nil || user.ID == 0 {
		ctx.Render(http.StatusBadRequest, render.Data{})
		return
	}

	if err := helper.SetSession(ctx, user.ID); err != nil {
		return
	}

	ctx.Redirect(http.StatusFound, os.Getenv("API_PATH")+"ping")
}

func SignIn(ctx *gin.Context) {
	var data loginFormData
	if err := ctx.Bind(&data); err != nil {
		ctx.Render(http.StatusBadRequest, render.Data{})
		return
	}

	checks := []helper.Check{
		helper.NewCheck(checkUsernameExists(data.Username), "User does not exist", true),
	}

	if helper.CheckAndRespond(ctx, checks) {
		return
	}

	user, err := sql.CheckPasswordMatch(data.Username, data.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := helper.SetSession(ctx, user.ID); err != nil {
		return
	}

	ctx.Redirect(http.StatusFound, os.Getenv("API_PATH")+"ping")
}

func SignOut(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()

	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	ctx.Redirect(http.StatusFound, "/")
}

func checkUsernameExists(username string) helper.CheckFunc {
	return func() (bool, error) {
		return sql.CheckUserExistenceByUsername(username)
	}
}

func checkEmailExists(email string) helper.CheckFunc {
	return func() (bool, error) {
		return sql.CheckUserExistenceByEmail(email)
	}
}
