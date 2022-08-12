package factory

import "fmt"

type CarFactory struct{}

type Vehicle interface {
	setName(name string)
	setPrice(price float64)
	getName() string
	getPrice() float64
}

type Car struct {
	name  string
	price float64
}

// setters

func (c *Car) setName(name string) {
	c.name = name
}

func (c *Car) setPrice(price float64) {
	c.price = price
}

// getters

func (c *Car) getName() string {
	return c.name
}

func (c *Car) getPrice() float64 {
	return c.price
}

type Ford struct {
	Car
}

type BMW struct {
	Car
}

func newFord() Vehicle {
	return &Ford{
		Car: Car{
			name:  "Ford",
			price: 105000.00,
		},
	}
}

func newBMW() Vehicle {
	return &BMW{
		Car: Car{
			name:  "BMW",
			price: 250000.00,
		},
	}
}

func (cf *CarFactory) createCar(carName string) (Vehicle, error) {

	switch carName {
	case "Ford":
		return newFord(), nil
	case "BMW":
		return newBMW(), nil
	}

	return nil, fmt.Errorf("invalid car name")
}

func printCarDetails(v Vehicle) {
	fmt.Printf("Car name: %s, car price: %f PLN\n", v.getName(), v.getPrice())
}
