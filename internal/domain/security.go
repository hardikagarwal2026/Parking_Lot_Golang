package domain

type Security interface {
	OnLotFull(message string)
}