package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Import pprof package

	"github.com/labstack/echo/v4" // Import Echo package
	users "test.go/user"
)

func main() {
	// Initialize a new in-memory user repository
	repo := users.NewInMemoryUserRepository()

	// Create a new user
	user := users.User{ID: 1, Name: "John Doe", Email: "john@example.com"}
	err := repo.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	// Read the user
	readUser, err := repo.Read(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Read user:", readUser)

	// Update the user
	updateUser := users.User{ID: 1, Name: "John Doe", Email: "updated@example.com"}
	err = repo.Update(1, updateUser)
	if err != nil {
		log.Fatal(err)
	}

	// Read the updated user
	readUpdatedUser, err := repo.Read(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated user:", readUpdatedUser)

	// Delete the user
	err = repo.Delete(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User deleted successfully")

	// Start a simple web server to expose CRUD operations and profiling endpoints
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil)) // pprof HTTP server
	}()

	// Start a simple web server to expose CRUD operations using Echo
	e := echo.New()

	e.POST("/user/create", func(c echo.Context) error {
		var newUser users.User
		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if err := repo.Create(newUser); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusCreated)
	})

	e.GET("/user/read", func(c echo.Context) error {
		id := c.QueryParam("id")
		var userID int
		_, err := fmt.Sscanf(id, "%d", &userID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		user, err := repo.Read(userID)
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	})

	e.PUT("/user/update", func(c echo.Context) error {
		var updatedUser users.User
		if err := c.Bind(&updatedUser); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if err := repo.Update(updatedUser.ID, updatedUser); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusOK)
	})

	e.DELETE("/user/delete", func(c echo.Context) error {
		id := c.QueryParam("id")
		var userID int
		_, err := fmt.Sscanf(id, "%d", &userID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if err := repo.Delete(userID); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusOK)
	})

	log.Fatal(e.Start(":8080"))
}
