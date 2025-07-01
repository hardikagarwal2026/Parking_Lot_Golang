package domain

//owner represents the parking lot owner who gets notified when the lot is full
type Owner interface {
	OnLotFull(message string) // this method called when parking lot reaches capacity
    OnSpaceAvailable(message string) //this method called when parking lot has space
}
