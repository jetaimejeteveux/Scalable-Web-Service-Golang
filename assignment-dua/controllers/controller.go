package controllers

import (
	"assignment-dua/structs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

var db *gorm.DB

func (Conn *DBconn) CreateOrder(c *gin.Context) {
	var order structs.Orders

	if err := c.BindJSON(&order); err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, order)

	Conn.DB.Create(&order)

	//result := gin.H{
	//	"result": order,
	//}
	//
	//c.JSON(http.StatusOK, result)
}

func (Conn *DBconn) GetOrderById(c *gin.Context) {
	var order structs.Orders
	var result gin.H

	id := c.Param("id")
	err := Conn.DB.Where("id = ?", id).First(&order).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusNotFound, result)
	} else {
		c.IndentedJSON(http.StatusOK, order)
	}

}

func (Conn *DBconn) GetOrders(c *gin.Context) {
	var orders []structs.Orders
	var result gin.H

	Conn.DB.Find(&orders)

	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
		}
		c.JSON(http.StatusNotFound, result)
	} else {
		c.IndentedJSON(http.StatusOK, orders)
	}

}

func (Conn *DBconn) UpdateOrder(c *gin.Context) {
	id := c.Query("id")
	var (
		oldOrder structs.Orders
		newOrder structs.Orders
		result   gin.H
	)

	err := Conn.DB.First(&oldOrder, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	if err := c.BindJSON(&newOrder); err != nil {
		return
	}

	err = Conn.DB.Model(&oldOrder).Updates(newOrder).Error
	if err != nil {
		result = gin.H{
			"result": "Update failed",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	c.IndentedJSON(http.StatusOK, oldOrder)
	return

}

func (Conn *DBconn) DeletePerson(c *gin.Context) {
	var (
		order  structs.Orders
		result gin.H
	)
	id := c.Param("id")
	err := Conn.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	err = Conn.DB.Delete(&order).Error
	if err != nil {
		result = gin.H{
			"result": "Delete failed",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"result": "Delete data success",
	}
	c.JSON(http.StatusOK, result)
	return
}
