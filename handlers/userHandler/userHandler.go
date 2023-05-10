package userHandler

import (
	"net/http"

	"github.com/dioncodes/go-surreal-starter/pkg/env"
	"github.com/dioncodes/go-surreal-starter/pkg/models/user"
	"github.com/gin-gonic/gin"
)

func Register(env *env.Env, r *gin.Engine) {
	g := r.Group("/users")
	{
		g.GET("", func(c *gin.Context) {
			getUsers(env, c)
		})
		g.POST("", func(c *gin.Context) {
			createUser(env, c)
		})

		g.GET(":id", func(c *gin.Context) {
			getUser(env, c)
		})
		g.PATCH(":id", func(c *gin.Context) {
			updateUser(env, c)
		})
		g.DELETE(":id", func(c *gin.Context) {
			deleteUser(env, c)
		})
	}
}

func getUsers(env *env.Env, c *gin.Context) {
	users, err := user.GetUsers(env)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "notFound"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func getUser(env *env.Env, c *gin.Context) {
	type Request struct {
		Id string `uri:"id" binding:"required"`
	}

	var r Request

	if err := c.ShouldBindUri(&r); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	user, err := user.GetUserById(env, r.Id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "notFound"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func createUser(env *env.Env, c *gin.Context) {
	type Request struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	var r Request

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	user := user.User{
		FirstName: r.FirstName,
		LastName: r.LastName,
	}

	if err := user.Save(env); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "dbError", "details": "error saving user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func updateUser(env *env.Env, c *gin.Context) {
	type Request struct {
		Id        string `uri:"id" binding:"required"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	var r Request

	if err := c.ShouldBindUri(&r); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	user, err := user.GetUserById(env, r.Id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "notFound"})
		return
	}

	user.FirstName = r.FirstName
	user.LastName = r.LastName

	if err := user.Save(env); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "dbError", "details": "error saving user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func deleteUser(env *env.Env, c *gin.Context) {
	type Request struct {
		Id        string `uri:"id" binding:"required"`
	}

	var r Request

	if err := c.ShouldBindUri(&r); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	user, err := user.GetUserById(env, r.Id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "notFound"})
		return
	}

	if err := user.Delete(env); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "dbError", "details": "error deleting user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
