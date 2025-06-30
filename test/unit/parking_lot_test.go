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

//after unparking, it should return true, when car is parked
func TestParkingLot_Unpark_ShouldReturnTrue_WhenCarIsParked(t *testing.T) {
    // Arrange
    lot := domain.NewParkingLot(100)
    car := domain.Car{
        Plate: "MH12AB1234",
        Make:  "Toyota",
        Color: "Blue",
    }
    lot.Park(car)
    result := lot.Unpark(car)
    if !result {
        t.Errorf("Expected car to be unparked successfully")
    }
}

// after unparking,it should return false, when car is not parked
func TestParkingLot_Unpark_ShouldReturnFalse_WhenCarNotParked(t *testing.T) {
    // Arrange
    lot := domain.NewParkingLot(100)
    car := domain.Car{
        Plate: "MH12AB1234",
        Make:  "Toyota",
        Color: "Blue",
    }
    result := lot.Unpark(car)
    if result {
        t.Errorf("Expected unpark to fail when car is not parked")
    }
}

func TestParkingLot_Unpark_ShouldReduceParkedCarsCount(t *testing.T) {
    lot := domain.NewParkingLot(100)
    car := domain.Car{
        Plate: "MH12AB1234",
        Make:  "Toyota",
        Color: "Blue",
    }
    lot.Park(car)
    lot.Unpark(car)
    if lot.GetParkedCarsCount() != 0 {
        t.Errorf("Expected 0 cars after unparking, got %d", lot.GetParkedCarsCount())
    }
}
