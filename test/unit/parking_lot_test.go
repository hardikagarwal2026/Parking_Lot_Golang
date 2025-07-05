package unit

import (
	"parking-lot-system/internal/domain"
	"testing"
)

// for owner notification testing use case -3
type MockOwner struct {
	WasNotified bool
	Message     string
    SpaceNotified bool
	SpaceMessage string
}

func (m *MockOwner) OnLotFull(message string) {
	m.WasNotified = true
	m.Message = message
}

func (m *MockOwner) OnSpaceAvailable(message string) {
	m.SpaceNotified = true
	m.SpaceMessage = message
}

// for security notification testing use case -4
type MockSecurity struct {
	WasNotified bool
	Message     string
}

func (m *MockSecurity) OnLotFull(message string) {
	m.WasNotified = true
	m.Message = message
}

// use case-1
// to test the parking lot , it should return true , when space available
func TestParkingLot_Park_true(t *testing.T) {
	lot := domain.NewParkingLot(100)
	car := domain.Car{
		Plate: "RJ14LJ81110",
		Make:  "Honda",
		Color: "Black",
	}

	result := lot.Park(car)

	if !result {
		t.Errorf("Expected car to be parked successfully")
	}
}

// to test the parking lot , it should return false , when space not available in lot
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

// use case-2
// after unparking, it should return true, when car is parked
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

// use case-3
func TestParkingLot_IsFull_ShouldReturnTrue_WhenLotIsFull(t *testing.T) {
	// Arrange
	lot := domain.NewParkingLot(2)
	car1 := domain.Car{Plate: "MH12AB1234", Make: "Toyota", Color: "Blue"}
	car2 := domain.Car{Plate: "MH12AB5678", Make: "Honda", Color: "White"}

	lot.Park(car1)
	lot.Park(car2)

	// Act
	result := lot.IsFull()

	// Assert
	if !result {
		t.Errorf("Expected lot to be full")
	}
}

func TestParkingLot_IsFull_ShouldReturnFalse_WhenLotIsNotFull(t *testing.T) {
	lot := domain.NewParkingLot(2)
	car := domain.Car{Plate: "MH12AB1234", Make: "Toyota", Color: "Blue"}

	lot.Park(car)
	result := lot.IsFull()
	if result {
		t.Errorf("Expected lot NOT to be full")
	}
}
func TestParkingLot_NotifyOwner_WhenLotBecomesFull(t *testing.T) {

	lot := domain.NewParkingLot(1)
	owner := &MockOwner{}
	lot.AddOwnerObserver(owner)

	car := domain.Car{Plate: "MH12AB1234", Make: "Toyota", Color: "Blue"}

	lot.Park(car)

	if !owner.WasNotified {
		t.Errorf("Expected owner to be notified when lot becomes full")
	}

	if owner.Message != "Lot is full" {
		t.Errorf("Expected 'Lot is full' message, got '%s'", owner.Message)
	}
}

// use case -4
func TestParkingLot_NotifySecurity_WhenLotBecomesFull(t *testing.T) {
	lot := domain.NewParkingLot(1)
	security := &MockSecurity{}
	lot.AddSecurityObserver(security)
	car := domain.Car{Plate: "MH12AB1234", Make: "Toyota", Color: "Blue"}

	lot.Park(car)

	if !security.WasNotified {
		t.Errorf("Expected security to be notified when lot becomes full")
	}
	if security.Message != "Lot is full" {
		t.Errorf("Expected 'Lot is full' message, got '%s'", security.Message)
	}
}
