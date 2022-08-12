package factory

import "testing"

func TestFactory(t *testing.T) {
	factory := CarFactory{}
	ford, _ := factory.createCar("Ford")
	bmw, _ := factory.createCar("BMW")
	printCarDetails(ford)
	printCarDetails(bmw)

}
