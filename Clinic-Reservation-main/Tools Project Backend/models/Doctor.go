package models

type Doctor struct {
	ID       uint `gorm:primaryKey`
	Name     string
	Mail     string
	Password string
	Type     string
	Slots    []Slot
}
