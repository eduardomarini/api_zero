package main

import (
	"github.com/labstack/echo/v4"
)

type Car struct {
	Name  string
	Model string
}

var cars []Car

func GenerateCars() {
	cars = append(cars, Car{Name: "Toyota", Model: "Corolla"})
	cars = append(cars, Car{Name: "Toyota", Model: "Camry"})
	cars = append(cars, Car{Name: "Toyota", Model: "Avalon"})
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func createCar(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	return c.JSON(200, cars)
}

func main() {

	GenerateCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCar)
	e.Logger.Fatal(e.Start(":3333"))

	// CreateCars()
	// for _, car := range cars {
	// 	fmt.Println(car.Name, car.Model)
	// }
}
