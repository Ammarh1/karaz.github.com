package controllers

import (
	"encoding/json"
	"karaz/models"

	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	result, err := models.DB.Query("SELECT ID,First_name, Last_name from users")
	var users []models.User

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(result)

	for result.Next() {
		var user models.User
		// for each row, scan the result into our tag composite object
		err = result.Scan(&user.ID, &user.First_name, &user.Last_name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		users = append(users, user)
	}

	c.JSON(200, gin.H{
		"data": users,
	})

}

func PostUser(c *gin.Context) {
	var user models.User
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	err = json.Unmarshal(value, &user)
	if err != nil {
		panic(err.Error())
	}

	insert, err := models.DB.Query("INSERT INTO users (first_name,last_name) VALUES(?,?)", user.First_name, user.Last_name)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	//fmt.Println(string(value))

	c.JSON(200, gin.H{
		"message": "Success",
	})

}

func UpdateUser(c *gin.Context) {
	var user models.User
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	err = json.Unmarshal(value, &user)
	if err != nil {
		panic(err.Error())
	}

	insert, err := models.DB.Query("UPDATE users SET first_name =?, last_name=? WHERE ID=?", user.First_name, user.Last_name, user.ID)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	//fmt.Println(string(value))

	c.JSON(200, gin.H{
		"message": "Success",
	})

}

func DeletUser(c *gin.Context) {
	var user models.User
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	err = json.Unmarshal(value, &user)
	if err != nil {
		panic(err.Error())
	}

	insert, err := models.DB.Query("DELETE FROM users WHERE first_name =? AND last_name=?", user.First_name, user.Last_name)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	fmt.Println(string(value))

	c.JSON(200, gin.H{
		"message": "Success",
	})

}
