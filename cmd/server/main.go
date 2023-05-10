package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dioncodes/go-surreal-starter/internal/router"
	"github.com/dioncodes/go-surreal-starter/pkg/env"
	"github.com/dioncodes/go-surreal-starter/pkg/models/user"
	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("BASE_DIR") == "" {
		os.Setenv("BASE_DIR", ".")
	}

	fmt.Println("Starting server...")

	// init env container with config and db connection
	env := env.Init()
	defer env.DB.Close()

	go func() {
		for {
			env.CheckDbConnection()
			time.Sleep(time.Second * 60 * 5)
		}
	}()

	if os.Getenv("ENV") != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create and save demo user
	john := user.User{
		FirstName: "John",
		LastName:  "Doe",
	}

	if err := john.Save(env); err != nil {
		fmt.Println("Error saving first user.", err.Error())
	}

	fmt.Println("----------")
	fmt.Println("Example user ID: " + john.Id)
	fmt.Println("You can get the user details using GET http://localhost:8100/users/" + strings.Replace(john.Id, "user:", "", 1))
	fmt.Println("----------")

	r := gin.Default()

	// setup routing and middleware
	router.Setup(env, r)

	r.Run(":" + os.Getenv("PORT"))
}
