package unit

import (
	"testing"
	"parking-lot-system/internal/domain"

)


//to test the parking lot , it should return true , when space available
func TestParkingLot_Park_true(t *testing.T) {
	lot := domain.NewParkingLot(100)
	car:= domain.Car{
		Plate: "RJ14LJ81110",
		Make: "Honda",
		Color: "Black",
	}

	result := lot.Park(car)

	if !result{
		t.Errorf("Expected car to be parked successfully")
	}
}

//to test the parking lot , it should return false , when space not available in lot
func TestParkingLot_Park_false(t *testing.T) {
	lot := domain.NewParkingLot(1) //parking lot with capacity one
	car1 := domain.Car{Plate: "MH12AB1234", Make: "Toyota", Color: "Blue"}
    car2 := domain.Car{Plate: "MH12AB5678", Make: "Honda", Color: "White"}
	lot.Park(car1)
	result := lot.Park(car2)

	if result {
		t.Errorf("Expected Car not to be parked when lot is full")
	}
}
