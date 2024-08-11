package controllers

import (
	"net/http"
	"rental-mobil/database"
	"rental-mobil/repository"
	"rental-mobil/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllRentals(c *gin.Context) {
	rentals, err := repository.GetAllRentals(database.DbConnection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rentals})
}

func InsertRental(c *gin.Context) {
	var rental structs.Rental
	if err := c.ShouldBindJSON(&rental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.InsertRental(database.DbConnection, rental)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rental})
}

func UpdateRental(c *gin.Context) {
	var rental structs.Rental
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rental ID"})
		return
	}

	if err := c.ShouldBindJSON(&rental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rental.RentalID = id

	err = repository.UpdateRental(database.DbConnection, rental)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rental})
}

func DeleteRental(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = repository.DeleteRental(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rental deleted"})
}
