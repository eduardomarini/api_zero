package main

import (
	"github.com/labstack/echo/v4"
)

type Car struct {
	Name  string
	Model string
}

var cars []Car

func CreateCars() {
	cars = append(cars, Car{Name: "Toyota", Model: "Corolla"})
	cars = append(cars, Car{Name: "Toyota", Model: "Camry"})
	cars = append(cars, Car{Name: "Toyota", Model: "Avalon"})
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func main() {

	CreateCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.Logger.Fatal(e.Start(":3333"))

	// CreateCars()
	// for _, car := range cars {
	// 	fmt.Println(car.Name, car.Model)
	// }
}
