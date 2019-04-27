package controllers

import (
	"fmt"
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) Login(c *gin.Context) {
	var (
		User   models.User
		result gin.H
	)
	username := c.PostForm("Username")
	password := c.PostForm("Password")
	err := idb.DB.Where("username = ?", username).First(&User).Error
	if err != nil {
		result = gin.H{
			"result": "Username tidak terdaftar",
			"status": "fail",
		}
		c.JSON(400, result)
	} else {
		if User.Password == password {
			c.JSON(http.StatusOK, gin.H{
				"result": User.Username,
				"status": "success",
			})
		} else {
			c.JSON(400, gin.H{
				"result": "Password Salah",
				"status": "fail",
			})
		}
	}

}

// to get one data with {id}
func (idb *InDB) GetUser(c *gin.Context) {
	var (
		User   models.User
		result gin.H
	)
	id := c.Param("Username")
	err := idb.DB.Where("username = ?", id).First(&User).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": User,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// get all data in User
func (idb *InDB) GetUsers(c *gin.Context) {
	var (
		Users  []models.User
		result gin.H
	)

	idb.DB.Find(&Users)
	if len(Users) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": Users,
			"count":  len(Users),
		}
	}

	c.JSON(http.StatusOK, result)
}

// create new data to the database
func (idb *InDB) CreateUser(c *gin.Context) {
	var (
		User   models.User
		result gin.H
	)
	fmt.Println(c.PostForm("Username"))
	Username := c.PostForm("Username")
	Password := c.PostForm("Password")

	err := idb.DB.Where("username = ?", Username).First(&User).Error
	if err != nil {
		User.Username = Username
		User.Password = Password
		idb.DB.Create(&User)
		result = gin.H{
			"result": User,
		}
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(400, gin.H{
			"result": "Username sudah terdaftar",
			"status": "fail",
		})
	}
}

// update data with {id} as query
func (idb *InDB) UpdateUser(c *gin.Context) {
	id := c.Query("id")
	Username := c.PostForm("Username")
	Password := c.PostForm("Password")
	var (
		User    models.User
		newUser models.User
		result  gin.H
	)

	err := idb.DB.First(&User, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	newUser.Username = Username
	newUser.Password = Password
	err = idb.DB.Model(&User).Updates(newUser).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

// delete data with {id}
func (idb *InDB) DeleteUser(c *gin.Context) {
	var (
		User   models.User
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&User, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&User).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
