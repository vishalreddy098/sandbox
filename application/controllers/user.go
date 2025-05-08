package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "sandbox/application/models"
)

var users = []models.User{}

func GetUsers(c *gin.Context) {
    c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
    var newUser models.User
    if err := c.BindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    users = append(users, newUser)
    c.JSON(http.StatusCreated, newUser)
}

func GetUser(c *gin.Context) {
    id := c.Param("id")
    for _, user := range users {
        if user.ID == id {
            c.JSON(http.StatusOK, user)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var updatedUser models.User
    if err := c.BindJSON(&updatedUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, user := range users {
        if user.ID == id {
            users[i] = updatedUser
            c.JSON(http.StatusOK, updatedUser)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
