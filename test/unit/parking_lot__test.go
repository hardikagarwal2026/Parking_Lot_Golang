package unit

import (
    "fmt"
	"parking-lot-system/internal/domain"
	"testing"
    "time"
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



//use case-6 Testing
// MockParkingAttendant for testing use case 6
type MockParkingAttendant struct {
	Name string
}

func (a *MockParkingAttendant) GetName() string {
	return a.Name
}

func TestParkingAttendant_ParkCar_ShouldReturnTrue_WhenSpaceAvailable(t *testing.T) {
	lot := domain.NewParkingLot(100)
	attendant := domain.NewParkingAttendant("John Doe")
	car := domain.Car{
		Plate: "MH12AB1234",
		Make:  "Toyota",
		Color: "Blue",
	}

	result := attendant.ParkCar(lot, car)

	if !result {
		t.Errorf("Expected attendant to park car successfully")
	}
	if lot.GetParkedCarsCount() != 1 {
		t.Errorf("Expected 1 car to be parked, got %d", lot.GetParkedCarsCount())
	}
}

func TestParkingAttendant_ParkCar_ShouldReturnFalse_WhenLotIsFull(t *testing.T) {
	lot := domain.NewParkingLot(1)
	attendant := domain.NewParkingAttendant("John Doe")
	car1 := domain.Car{Plate: "MH12AB1234", Make: "Toyota", Color: "Blue"}
	car2 := domain.Car{Plate: "MH12AB5678", Make: "Honda", Color: "White"}

	attendant.ParkCar(lot, car1)
	result := attendant.ParkCar(lot, car2)

	if result {
		t.Errorf("Expected attendant to fail parking when lot is full")
	}
}

func TestParkingAttendant_UnparkCar_ShouldReturnTrue_WhenCarIsParked(t *testing.T) {
	lot := domain.NewParkingLot(100)
	attendant := domain.NewParkingAttendant("John Doe")
	car := domain.Car{
		Plate: "MH12AB1234",
		Make:  "Toyota",
		Color: "Blue",
	}

	attendant.ParkCar(lot, car)
	result := attendant.UnparkCar(lot, car)

	if !result {
		t.Errorf("Expected attendant to unpark car successfully")
	}
	if lot.GetParkedCarsCount() != 0 {
		t.Errorf("Expected 0 cars after unparking, got %d", lot.GetParkedCarsCount())
	}
}


//use case -7
func TestParkingLot_FindCar_ShouldReturnSlotNumber_WhenCarIsParked(t *testing.T) {
	lot := domain.NewParkingLot(100)
	car := domain.Car{
		Plate: "MH12AB1234",
		Make:  "Toyota",
		Color: "Blue",
	}
	
	lot.Park(car)
	slotNumber := lot.FindCar(car.Plate)
	
	if slotNumber == -1 {
		t.Errorf("Expected to find car, but got -1")
	}
	if slotNumber != 0 { // First car should be at index 0
		t.Errorf("Expected car at slot 0, got %d", slotNumber)
	}
}

func TestParkingLot_FindCar_ShouldReturnMinusOne_WhenCarNotParked(t *testing.T) {
	lot := domain.NewParkingLot(100)
	
	slotNumber := lot.FindCar("MH12AB1234")
	
	if slotNumber != -1 {
		t.Errorf("Expected -1 when car not found, got %d", slotNumber)
	}
}

func TestParkingLot_FindCar_ShouldReturnCorrectSlot_WhenMultipleCarsParked(t *testing.T) {
	lot := domain.NewParkingLot(100)
	car1 := domain.Car{Plate: "MH12AB1234", Make: "Toyota", Color: "Blue"}
	car2 := domain.Car{Plate: "MH12AB5678", Make: "Honda", Color: "White"}
	car3 := domain.Car{Plate: "MH12AB9999", Make: "BMW", Color: "Black"}
	
	lot.Park(car1)
	lot.Park(car2)
	lot.Park(car3)
	
	slotNumber := lot.FindCar("MH12AB5678")
	
	if slotNumber != 1 { // Second car should be at index 1
		t.Errorf("Expected car at slot 1, got %d", slotNumber)
	}
}


//use case -8
func TestParkingLot_GetParkingTime_ShouldReturnTime_WhenCarIsParked(t *testing.T) {
    lot := domain.NewParkingLot(100)
    car := domain.Car{
        Plate: "MH12AB1234",
        Make:  "Toyota",
        Color: "Blue",
    }
    
    beforePark := time.Now()
    lot.Park(car)
    afterPark := time.Now()
    
    parkingTime := lot.GetParkingTime(car.Plate)
    
    if parkingTime.IsZero() {
        t.Errorf("Expected parking time to be recorded")
    }
    if parkingTime.Before(beforePark) || parkingTime.After(afterPark) {
        t.Errorf("Parking time should be between %v and %v, got %v", beforePark, afterPark, parkingTime)
    }
}

func TestParkingLot_GetParkingTime_ShouldReturnZeroTime_WhenCarNotParked(t *testing.T) {
    lot := domain.NewParkingLot(100)
    
    parkingTime := lot.GetParkingTime("MH12AB1234")
    
    if !parkingTime.IsZero() {
        t.Errorf("Expected zero time when car not parked, got %v", parkingTime)
    }
}

func TestParkingLot_GetParkingDuration_ShouldReturnDuration_WhenCarIsParked(t *testing.T) {
    lot := domain.NewParkingLot(100)
    car := domain.Car{
        Plate: "MH12AB1234",
        Make:  "Toyota",
        Color: "Blue",
    }
    
    lot.Park(car)
    time.Sleep(10 * time.Millisecond) // Small delay to ensure duration > 0
    
    duration := lot.GetParkingDuration(car.Plate)
    
    if duration <= 0 {
        t.Errorf("Expected positive duration, got %v", duration)
    }
}

func TestParkingLot_RemoveParkingTime_WhenCarIsUnparked(t *testing.T) {
    lot := domain.NewParkingLot(100)
    car := domain.Car{
        Plate: "MH12AB1234",
        Make:  "Toyota",
        Color: "Blue",
    }
    
    lot.Park(car)
    lot.Unpark(car)
    
    parkingTime := lot.GetParkingTime(car.Plate)
    
    if !parkingTime.IsZero() {
        t.Errorf("Expected parking time to be removed after unparking")
    }
}


//use-case -9
func TestParkingAttendant_EvenDistribution_ShouldParkInLotWithFewerCars(t *testing.T) {
    lot1 := domain.NewParkingLot(100)
    lot2 := domain.NewParkingLot(100)
    lots := []*domain.ParkingLot{lot1, lot2}
    attendant := domain.NewParkingAttendant("John Doe")
    
    car1 := domain.Car{Plate: "MH12AB1234", Make: "Toyota", Color: "Blue"}
    car2 := domain.Car{Plate: "MH12AB5678", Make: "Honda", Color: "White"}
    car3 := domain.Car{Plate: "MH12AB9999", Make: "BMW", Color: "Black"}
    
    // Park first car - should go to lot1 (both empty, picks first)
    result1 := attendant.ParkCarEvenly(lots, car1)
    if !result1 {
        t.Errorf("Expected first car to be parked")
    }
    if lot1.GetParkedCarsCount() != 1 {
        t.Errorf("Expected lot1 to have 1 car, got %d", lot1.GetParkedCarsCount())
    }
    
    // Park second car - should go to lot2 (to balance)
    result2 := attendant.ParkCarEvenly(lots, car2)
    if !result2 {
        t.Errorf("Expected second car to be parked")
    }
    if lot2.GetParkedCarsCount() != 1 {
        t.Errorf("Expected lot2 to have 1 car, got %d", lot2.GetParkedCarsCount())
    }
    
    // Park third car - should go to lot1 again (both have 1, picks first)
    result3 := attendant.ParkCarEvenly(lots, car3)
    if !result3 {
        t.Errorf("Expected third car to be parked")
    }
    if lot1.GetParkedCarsCount() != 2 {
        t.Errorf("Expected lot1 to have 2 cars, got %d", lot1.GetParkedCarsCount())
    }
}

func TestParkingAttendant_EvenDistribution_ShouldReturnFalse_WhenAllLotsFull(t *testing.T) {
    lot1 := domain.NewParkingLot(1)
    lot2 := domain.NewParkingLot(1)
    lots := []*domain.ParkingLot{lot1, lot2}
    attendant := domain.NewParkingAttendant("John Doe")
    
    car1 := domain.Car{Plate: "MH12AB1234", Make: "Toyota", Color: "Blue"}
    car2 := domain.Car{Plate: "MH12AB5678", Make: "Honda", Color: "White"}
    car3 := domain.Car{Plate: "MH12AB9999", Make: "BMW", Color: "Black"}
    
    // Fill both lots
    attendant.ParkCarEvenly(lots, car1)
    attendant.ParkCarEvenly(lots, car2)
    
    // Try to park third car - should fail
    result := attendant.ParkCarEvenly(lots, car3)
    if result {
        t.Errorf("Expected parking to fail when all lots are full")
    }
}


//use case-10
func TestParkingAttendant_HandicapParking_ShouldParkInNearestLot(t *testing.T) {
    lot1 := domain.NewParkingLot(100)
    lot2 := domain.NewParkingLot(100)
    lot3 := domain.NewParkingLot(100)
    lots := []*domain.ParkingLot{lot1, lot2, lot3}
    attendant := domain.NewParkingAttendant("John Doe")
    
    handicapCar := domain.Car{
        Plate: "MH12AB1234", 
        Make:  "Toyota", 
        Color: "Blue",
        Size:  domain.Small,
    }
    
    // Park handicap car - should go to first available lot (nearest)
    result := attendant.ParkHandicapCar(lots, handicapCar)
    
    if !result {
        t.Errorf("Expected handicap car to be parked")
    }
    if lot1.GetParkedCarsCount() != 1 {
        t.Errorf("Expected lot1 to have 1 car, got %d", lot1.GetParkedCarsCount())
    }
}

func TestParkingAttendant_HandicapParking_ShouldSkipFullLots(t *testing.T) {
    lot1 := domain.NewParkingLot(1)
    lot2 := domain.NewParkingLot(100)
    lots := []*domain.ParkingLot{lot1, lot2}
    attendant := domain.NewParkingAttendant("John Doe")
    
    // Fill first lot
    regularCar := domain.Car{Plate: "MH12AB5678", Make: "Honda", Color: "White"}
    lot1.Park(regularCar)
    
    handicapCar := domain.Car{
        Plate: "MH12AB1234", 
        Make:  "Toyota", 
        Color: "Blue",
        Size:  domain.Small,
    }
    
    // Park handicap car - should skip full lot1 and go to lot2
    result := attendant.ParkHandicapCar(lots, handicapCar)
    
    if !result {
        t.Errorf("Expected handicap car to be parked")
    }
    if lot2.GetParkedCarsCount() != 1 {
        t.Errorf("Expected lot2 to have 1 car, got %d", lot2.GetParkedCarsCount())
    }
}

func TestParkingAttendant_HandicapParking_ShouldReturnFalse_WhenAllLotsFull(t *testing.T) {
    lot1 := domain.NewParkingLot(1)
    lot2 := domain.NewParkingLot(1)
    lots := []*domain.ParkingLot{lot1, lot2}
    attendant := domain.NewParkingAttendant("John Doe")
    
    // Fill both lots
    car1 := domain.Car{Plate: "MH12AB5678", Make: "Honda", Color: "White"}
    car2 := domain.Car{Plate: "MH12AB9999", Make: "BMW", Color: "Black"}
    lot1.Park(car1)
    lot2.Park(car2)
    
    handicapCar := domain.Car{
        Plate: "MH12AB1234", 
        Make:  "Toyota", 
        Color: "Blue",
        Size:  domain.Small,
    }
    
    // Try to park handicap car - should fail
    result := attendant.ParkHandicapCar(lots, handicapCar)
    
    if result {
        t.Errorf("Expected handicap parking to fail when all lots are full")
    }
}

//use case-11
func TestParkingAttendant_LargeCarParking_ShouldParkInLotWithMostSpace(t *testing.T) {
    lot1 := domain.NewParkingLot(100)
    lot2 := domain.NewParkingLot(100)
    lot3 := domain.NewParkingLot(100)
    lots := []*domain.ParkingLot{lot1, lot2, lot3}
    attendant := domain.NewParkingAttendant("John Doe")
    
    // Fill lot1 with 50 cars to make it less spacious
    for i := 0; i < 50; i++ {
        car := domain.Car{
            Plate: fmt.Sprintf("MH12AB%04d", i),
            Make:  "Honda",
            Color: "White",
            Size:  domain.Small,
        }
        lot1.Park(car)
    }
    
    // Fill lot2 with 30 cars to make it more spacious than lot1
    for i := 0; i < 30; i++ {
        car := domain.Car{
            Plate: fmt.Sprintf("MH12CD%04d", i),
            Make:  "Toyota",
            Color: "Blue",
            Size:  domain.Small,
        }
        lot2.Park(car)
    }
    
    // lot3 remains empty (most spacious)
    
    largeCar := domain.Car{
        Plate: "MH12AB1234",
        Make:  "SUV",
        Color: "Black",
        Size:  domain.Large,
    }
    
    // Large car should go to lot3 (most free space)
    result := attendant.ParkLargeCar(lots, largeCar)
    
    if !result {
        t.Errorf("Expected large car to be parked")
    }
    if lot3.GetParkedCarsCount() != 1 {
        t.Errorf("Expected lot3 to have 1 car, got %d", lot3.GetParkedCarsCount())
    }
    if lot1.GetParkedCarsCount() != 50 {
        t.Errorf("Expected lot1 to remain unchanged with 50 cars, got %d", lot1.GetParkedCarsCount())
    }
    if lot2.GetParkedCarsCount() != 30 {
        t.Errorf("Expected lot2 to remain unchanged with 30 cars, got %d", lot2.GetParkedCarsCount())
    }
}

func TestParkingAttendant_LargeCarParking_ShouldReturnFalse_WhenAllLotsFull(t *testing.T) {
    lot1 := domain.NewParkingLot(1)
    lot2 := domain.NewParkingLot(1)
    lots := []*domain.ParkingLot{lot1, lot2}
    attendant := domain.NewParkingAttendant("John Doe")
    
    // Fill both lots
    car1 := domain.Car{Plate: "MH12AB5678", Make: "Honda", Color: "White", Size: domain.Small}
    car2 := domain.Car{Plate: "MH12AB9999", Make: "BMW", Color: "Black", Size: domain.Medium}
    lot1.Park(car1)
    lot2.Park(car2)
    
    largeCar := domain.Car{
        Plate: "MH12AB1234",
        Make:  "SUV",
        Color: "Black",
        Size:  domain.Large,
    }
    
    // Try to park large car - should fail
    result := attendant.ParkLargeCar(lots, largeCar)
    
    if result {
        t.Errorf("Expected large car parking to fail when all lots are full")
    }
}

//for use case - 11
func TestParkingAttendant_LargeCarParking_ShouldChooseCorrectLot_WhenMultipleLotsAvailable(t *testing.T) {
    lot1 := domain.NewParkingLot(100) // 100 free spaces
    lot2 := domain.NewParkingLot(50)  // 50 free spaces
    lot3 := domain.NewParkingLot(75)  // 75 free spaces
    lots := []*domain.ParkingLot{lot1, lot2, lot3}
    attendant := domain.NewParkingAttendant("John Doe")
    
    largeCar := domain.Car{
        Plate: "MH12AB1234",
        Make:  "SUV",
        Color: "Black",
        Size:  domain.Large,
    }
    
    // Large car should go to lot1 (highest capacity = most free space)
    result := attendant.ParkLargeCar(lots, largeCar)
    
    if !result {
        t.Errorf("Expected large car to be parked")
    }
    if lot1.GetParkedCarsCount() != 1 {
        t.Errorf("Expected lot1 to have 1 car, got %d", lot1.GetParkedCarsCount())
    }
}


//use case-12
func TestParkingLot_FindCarsByColor_ShouldReturnWhiteCars_WhenWhiteCarsParked(t *testing.T) {
    lot := domain.NewParkingLot(100)
    
    whiteCar1 := domain.Car{Plate: "MH12AB1234", Make: "Toyota", Color: "White", Size: domain.Small}
    whiteCar2 := domain.Car{Plate: "MH12AB5678", Make: "Honda", Color: "White", Size: domain.Medium}
    blueCar := domain.Car{Plate: "MH12AB9999", Make: "BMW", Color: "Blue", Size: domain.Large}
    
    lot.Park(whiteCar1)
    lot.Park(whiteCar2)
    lot.Park(blueCar)
    
    whiteCars := lot.FindCarsByColor("White")
    
    if len(whiteCars) != 2 {
        t.Errorf("Expected 2 white cars, got %d", len(whiteCars))
    }
    
    // Verify the correct cars are returned
    foundPlates := make(map[string]bool)
    for _, car := range whiteCars {
        foundPlates[car.Plate] = true
    }
    
    if !foundPlates["MH12AB1234"] {
        t.Errorf("Expected to find white car MH12AB1234")
    }
    if !foundPlates["MH12AB5678"] {
        t.Errorf("Expected to find white car MH12AB5678")
    }
    if foundPlates["MH12AB9999"] {
        t.Errorf("Blue car should not be in white cars list")
    }
}

func TestParkingLot_FindCarsByColor_ShouldReturnEmptySlice_WhenNoMatchingCars(t *testing.T) {
    lot := domain.NewParkingLot(100)
    
    blueCar := domain.Car{Plate: "MH12AB1234", Make: "Toyota", Color: "Blue", Size: domain.Small}
    redCar := domain.Car{Plate: "MH12AB5678", Make: "Honda", Color: "Red", Size: domain.Medium}
    
    lot.Park(blueCar)
    lot.Park(redCar)
    
    whiteCars := lot.FindCarsByColor("White")
    
    if len(whiteCars) != 0 {
        t.Errorf("Expected 0 white cars, got %d", len(whiteCars))
    }
}

func TestParkingLot_FindCarsByColor_ShouldReturnEmptySlice_WhenLotEmpty(t *testing.T) {
    lot := domain.NewParkingLot(100)
    
    whiteCars := lot.FindCarsByColor("White")
    
    if len(whiteCars) != 0 {
        t.Errorf("Expected 0 white cars in empty lot, got %d", len(whiteCars))
    }
}

//to verify the police investigation functionality
func TestPoliceDepartment_InvestigateWhiteCars_ShouldReturnAllWhiteCarsWithLocations(t *testing.T) {
    lot1 := domain.NewParkingLot(100)
    lot2 := domain.NewParkingLot(100)
    lots := []*domain.ParkingLot{lot1, lot2}
    police := domain.NewPoliceDepartment("City Police")
    
    whiteCar1 := domain.Car{Plate: "MH12AB1234", Make: "Toyota", Color: "White", Size: domain.Small}
    whiteCar2 := domain.Car{Plate: "MH12AB5678", Make: "Honda", Color: "White", Size: domain.Medium}
    blueCar := domain.Car{Plate: "MH12AB9999", Make: "BMW", Color: "Blue", Size: domain.Large}
    
    lot1.Park(whiteCar1)
    lot1.Park(blueCar)
    lot2.Park(whiteCar2)
    
    whiteCarLocations := police.InvestigateWhiteCars(lots)
    
    if len(whiteCarLocations) != 2 {
        t.Errorf("Expected 2 white cars, got %d", len(whiteCarLocations))
    }
    
    // Verify locations are correct
    for _, location := range whiteCarLocations {
        if location.Car.Color != "White" {
            t.Errorf("Expected white car, got %s", location.Car.Color)
        }
        if location.LotID < 0 || location.LotID >= len(lots) {
            t.Errorf("Invalid lot ID: %d", location.LotID)
        }
        if location.SlotID < 0 {
            t.Errorf("Invalid slot ID: %d", location.SlotID)
        }
    }
}

//use case-13
func TestParkingLot_FindCarsByMakeAndColor_ShouldReturnBlueToyotas_WhenBlueToyotasParked(t *testing.T) {
    lot := domain.NewParkingLot(100)
    
    blueToyota1 := domain.Car{Plate: "MH12AB1234", Make: "Toyota", Color: "Blue", Size: domain.Small}
    blueToyota2 := domain.Car{Plate: "MH12AB5678", Make: "Toyota", Color: "Blue", Size: domain.Medium}
    blueHonda := domain.Car{Plate: "MH12AB9999", Make: "Honda", Color: "Blue", Size: domain.Large}
    redToyota := domain.Car{Plate: "MH12AB7777", Make: "Toyota", Color: "Red", Size: domain.Small}
    
    lot.Park(blueToyota1)
    lot.Park(blueToyota2)
    lot.Park(blueHonda)
    lot.Park(redToyota)
    
    blueToyotas := lot.FindCarsByMakeAndColor("Toyota", "Blue")
    
    if len(blueToyotas) != 2 {
        t.Errorf("Expected 2 blue Toyota cars, got %d", len(blueToyotas))
    }
    
    // Verify the correct cars are returned
    foundPlates := make(map[string]bool)
    for _, car := range blueToyotas {
        foundPlates[car.Plate] = true
        if car.Make != "Toyota" || car.Color != "Blue" {
            t.Errorf("Expected blue Toyota, got %s %s", car.Color, car.Make)
        }
    }
    
    if !foundPlates["MH12AB1234"] {
        t.Errorf("Expected to find blue Toyota MH12AB1234")
    }
    if !foundPlates["MH12AB5678"] {
        t.Errorf("Expected to find blue Toyota MH12AB5678")
    }
    if foundPlates["MH12AB9999"] {
        t.Errorf("Blue Honda should not be in blue Toyota list")
    }
    if foundPlates["MH12AB7777"] {
        t.Errorf("Red Toyota should not be in blue Toyota list")
    }
}

func TestParkingLot_FindCarsByMakeAndColor_ShouldReturnEmptySlice_WhenNoMatchingCars(t *testing.T) {
    lot := domain.NewParkingLot(100)
    
    blueHonda := domain.Car{Plate: "MH12AB1234", Make: "Honda", Color: "Blue", Size: domain.Small}
    redToyota := domain.Car{Plate: "MH12AB5678", Make: "Toyota", Color: "Red", Size: domain.Medium}
    
    lot.Park(blueHonda)
    lot.Park(redToyota)
    
    blueToyotas := lot.FindCarsByMakeAndColor("Toyota", "Blue")
    
    if len(blueToyotas) != 0 {
        t.Errorf("Expected 0 blue Toyota cars, got %d", len(blueToyotas))
    }
}

func TestPoliceDepartment_InvestigateBlueToyotas_ShouldReturnCompleteInformation(t *testing.T) {
    lot1 := domain.NewParkingLot(100)
    lot2 := domain.NewParkingLot(100)
    lots := []*domain.ParkingLot{lot1, lot2}
    attendant := domain.NewParkingAttendant("Officer Smith")
    police := domain.NewPoliceDepartment("City Police")
    
    blueToyota1 := domain.Car{Plate: "MH12AB1234", Make: "Toyota", Color: "Blue", Size: domain.Small}
    blueToyota2 := domain.Car{Plate: "MH12AB5678", Make: "Toyota", Color: "Blue", Size: domain.Medium}
    blueHonda := domain.Car{Plate: "MH12AB9999", Make: "Honda", Color: "Blue", Size: domain.Large}
    
    // Park cars using attendant
    attendant.ParkCar(lot1, blueToyota1)
    attendant.ParkCar(lot1, blueHonda)
    attendant.ParkCar(lot2, blueToyota2)
    
    blueToyotaInvestigation := police.InvestigateBlueToyotas(lots, attendant)
    
    if len(blueToyotaInvestigation) != 2 {
        t.Errorf("Expected 2 blue Toyota investigations, got %d", len(blueToyotaInvestigation))
    }
    
    // Verify investigation details
    for _, investigation := range blueToyotaInvestigation {
        if investigation.Car.Make != "Toyota" || investigation.Car.Color != "Blue" {
            t.Errorf("Expected blue Toyota, got %s %s", investigation.Car.Color, investigation.Car.Make)
        }
        if investigation.AttendantName != "Officer Smith" {
            t.Errorf("Expected attendant name 'Officer Smith', got '%s'", investigation.AttendantName)
        }
        if investigation.LotID < 0 || investigation.LotID >= len(lots) {
            t.Errorf("Invalid lot ID: %d", investigation.LotID)
        }
        if investigation.SlotID < 0 {
            t.Errorf("Invalid slot ID: %d", investigation.SlotID)
        }
    }
}


//use case-14
func TestParkingLot_FindCarsByMake_ShouldReturnBMWCars_WhenBMWCarsParked(t *testing.T) {
    lot := domain.NewParkingLot(100)
    
    bmwCar1 := domain.Car{Plate: "MH12AB1234", Make: "BMW", Color: "Black", Size: domain.Large}
    bmwCar2 := domain.Car{Plate: "MH12AB5678", Make: "BMW", Color: "White", Size: domain.Medium}
    toyotaCar := domain.Car{Plate: "MH12AB9999", Make: "Toyota", Color: "Blue", Size: domain.Small}
    hondaCar := domain.Car{Plate: "MH12AB7777", Make: "Honda", Color: "Red", Size: domain.Medium}
    
    lot.Park(bmwCar1)
    lot.Park(bmwCar2)
    lot.Park(toyotaCar)
    lot.Park(hondaCar)
    
    bmwCars := lot.FindCarsByMake("BMW")
    
    if len(bmwCars) != 2 {
        t.Errorf("Expected 2 BMW cars, got %d", len(bmwCars))
    }
    
    // Verify the correct cars are returned
    foundPlates := make(map[string]bool)
    for _, car := range bmwCars {
        foundPlates[car.Plate] = true
        if car.Make != "BMW" {
            t.Errorf("Expected BMW car, got %s", car.Make)
        }
    }
    
    if !foundPlates["MH12AB1234"] {
        t.Errorf("Expected to find BMW MH12AB1234")
    }
    if !foundPlates["MH12AB5678"] {
        t.Errorf("Expected to find BMW MH12AB5678")
    }
    if foundPlates["MH12AB9999"] {
        t.Errorf("Toyota should not be in BMW list")
    }
    if foundPlates["MH12AB7777"] {
        t.Errorf("Honda should not be in BMW list")
    }
}

func TestParkingLot_FindCarsByMake_ShouldReturnEmptySlice_WhenNoMatchingCars(t *testing.T) {
    lot := domain.NewParkingLot(100)
    
    toyotaCar := domain.Car{Plate: "MH12AB1234", Make: "Toyota", Color: "Blue", Size: domain.Small}
    hondaCar := domain.Car{Plate: "MH12AB5678", Make: "Honda", Color: "Red", Size: domain.Medium}
    
    lot.Park(toyotaCar)
    lot.Park(hondaCar)
    
    bmwCars := lot.FindCarsByMake("BMW")
    
    if len(bmwCars) != 0 {
        t.Errorf("Expected 0 BMW cars, got %d", len(bmwCars))
    }
}

func TestPoliceDepartment_InvestigateBMWCars_ShouldReturnAllBMWsWithLocations(t *testing.T) {
    lot1 := domain.NewParkingLot(100)
    lot2 := domain.NewParkingLot(100)
    lots := []*domain.ParkingLot{lot1, lot2}
    police := domain.NewPoliceDepartment("City Police")
    
    bmwCar1 := domain.Car{Plate: "MH12AB1234", Make: "BMW", Color: "Black", Size: domain.Large}
    bmwCar2 := domain.Car{Plate: "MH12AB5678", Make: "BMW", Color: "White", Size: domain.Medium}
    toyotaCar := domain.Car{Plate: "MH12AB9999", Make: "Toyota", Color: "Blue", Size: domain.Small}
    
    lot1.Park(bmwCar1)
    lot1.Park(toyotaCar)
    lot2.Park(bmwCar2)
    
    bmwInvestigation := police.InvestigateBMWCars(lots)
    
    if len(bmwInvestigation) != 2 {
        t.Errorf("Expected 2 BMW investigations, got %d", len(bmwInvestigation))
    }
    
    // Verify investigation details
    for _, investigation := range bmwInvestigation {
        if investigation.Car.Make != "BMW" {
            t.Errorf("Expected BMW car, got %s", investigation.Car.Make)
        }
        if investigation.LotID < 0 || investigation.LotID >= len(lots) {
            t.Errorf("Invalid lot ID: %d", investigation.LotID)
        }
        if investigation.SlotID < 0 {
            t.Errorf("Invalid slot ID: %d", investigation.SlotID)
        }
    }
}
