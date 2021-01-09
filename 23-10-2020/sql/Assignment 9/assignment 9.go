package main

import (
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

type Cycle struct {
	gorm.Model
	Frame         string `json:"frame"`
	Frame_price   int    `json:"frame_price"`
	Brakes        string `json:"break"`
	Brakes_price  int    `json:"brakes_price"`
	Wheel         string `json:"wheel"`
	Wheel_price   int    `json:"wheel_price"`
	Chain         string `json:"chain"`
	Chain_price   int    `json:"chain_price"`
	Seating       string `json:"seating"`
	Seating_price int    `json:"seating_price"`
	Total_price   int    `json:"total_price"`
	Month_buying  int    `json:"month_buying"`
	Year_buying   int    `json:"year_buying"`
}

var db *gorm.DB
var err error
var router *gin.Engine

func main() {
	db, err = gorm.Open("postgres", "host="+"localhost"+" user="+"postgres"+" dbname="+"hari"+" sslmode=disable password="+"root")

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Cycle{})
	router = gin.Default()
	router.GET("/", Menu)
	router.POST("/order", orders)
	router.GET("/order", allorders)
	router.Run(":123")
}

func Menu(c *gin.Context) {
	menu1 := "Chose your order : frame:( steel, Stainless steel, iron) , brakes:( 1 gear, 2 gear, 3 gear, 4 gear) ,  wheels:(tubeless tyre, tube tyre)  , chain:(gear, without gear)  , seating:(double , single)  , buying_month:(1-12) ,  buying_year(after 2016):"
	c.JSON(200, menu1)
}

func orders(c *gin.Context) {
	var month int
	var F_price int
	var B_price int
	var W_price int
	var C_price int
	var S_price int
	var total_price int
	var input Cycle
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	year := input.Year_buying - 2016
	month = year*12 + input.Month_buying
	month = month * 10

	if input.Frame == "steel" {
		F_price = 1800
	} else if input.Frame == "stainless steel" {
		F_price = 2400
	} else {
		F_price = 1600
	}

	if input.Brakes == "1 gear" {
		B_price = 1200
	} else if input.Brakes == "2 gear" {
		B_price = 1500
	} else if input.Brakes == "3 gear" {
		B_price = 1800
	} else {
		B_price = 2100
	}

	if input.Wheel == "tubeless" {
		W_price = 2500
	} else {
		W_price = 1700
	}

	if input.Chain == "gear" {
		C_price = 2000
	} else {
		C_price = 1300
	}

	if input.Seating == "double" {
		S_price = 2500
	} else {
		S_price = 1700
	}

	F_price = F_price + month
	B_price = B_price + month
	W_price = W_price + month
	C_price = C_price + month
	S_price = S_price + month
	total_price = F_price + B_price + W_price + C_price + S_price

	cycles := Cycle{Frame: input.Frame, Frame_price: F_price, Brakes: input.Brakes, Brakes_price: B_price, Wheel: input.Wheel, Wheel_price: W_price, Chain: input.Chain, Chain_price: C_price, Seating: input.Seating, Seating_price: S_price, Total_price: total_price, Month_buying: input.Month_buying, Year_buying: input.Year_buying}
	db.Create(&cycles)
	c.JSON(200, cycles)
}

func allorders(c *gin.Context) {
	var order []Cycle
	db.Find(&order)

	c.JSON(http.StatusOK, order)
}
