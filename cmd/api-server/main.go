package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"database/sql"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	// get key from file for token use
	key, err := os.ReadFile("./key.txt")
	if err != nil {
		log.Fatal("key not found")
	}

	// connect to the database
	connStr := "postgres://postgres:user@localhost/card2go?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("could not connect to database")
	}

	// #region printusertable
	// prints everything inside the user table
	var users []User
	// only get the id and name columns
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal("could not execute SQL query ", err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var user User
		// assign id and name columns to a user object
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			log.Fatal("wtf ", err.Error()) // xd
		}
		// add to users array
		users = append(users, user)
	}

	log.Println("users table")
	for _, user := range users { // loop through the users array and print every user, don't need index so it is a _
		log.Printf("User Id %d; Name %s;\n", user.ID, user.Name)
	}
	// #endregion printusertable

	// create router for the endpoints
	router := gin.Default()

	// #region endpoints
	// test endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// auth endpoint, to be used for logging in
	router.POST("/auth", func(c *gin.Context) {
		/*
			the body for this post request goes like this
			{
				"username": "vaugnh",
				"password": "lovesyou"
			}
		*/
		// login object temporary, represents the json object above
		type LoginAttempt struct {
			Username string `json: "username" binding:"required"`
			Password string `json: "password" binding:"required"`
		}

		var attempt LoginAttempt //todo figure out how to do this without a struct
		if c.ShouldBindJSON(&attempt) == nil {
			var id int64
			// verify if person received has an actual entry in the database and get the id for that person
			row := db.QueryRow("SELECT id FROM users WHERE name = $1 AND password = $2;", attempt.Username, attempt.Password)
			if err := row.Scan(&id); err == nil {
				// create auth token
				t := jwt.NewWithClaims(jwt.SigningMethodHS256,
					jwt.MapClaims{
						"id": id, // include the user's id with the token so we can know who it is when they use the token
					})
				s, err := t.SignedString(key)
				if err == nil {
					// give the token to the sender
					c.JSON(200, gin.H{
						"token": s,
					})
				} else {
					// how does this happen
					c.JSON(400, gin.H{
						"error": fmt.Sprintf("could not create a token %s", err.Error()),
					})
				}
			} else {
				// i think QueryRow fails if there is no result idk
				c.JSON(400, gin.H{
					"error": "account not found",
				})
			}
		} else {
			c.JSON(400, gin.H{
				"error": fmt.Sprintf("invalid request %s", err.Error()),
			})
		}
	})
	// #endregion endpoints

	router.Run()
	log.Print("execution finished")
}

// user model in the database, don't need to store password anywhere so it is not included
type User struct {
	ID   int64
	Name string
}
