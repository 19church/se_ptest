package entity

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name string

	Admins []Admin `gorm:"foreignKey:AdminID"`
}

type Student struct {
	gorm.Model
	Name string

	Students []Student `gorm:"foreignKey:StudentID"`
}

type DisciplineType struct {
	gorm.Model
	Name string

	DisciplineTypes []DisciplineType `gorm:"foreignKey:DisciplineTypeID"`
}

type Discipline struct {
	gorm.Model

	//Admin FK
	AdminID *uint
	Admin   Admin `gorm:"references:id"`

	//Student FK
	StudentID *uint
	Student   Student `gorm:"references:id"`

	//DisciplineType FK
	DisciplineTypeID *uint
	DisciplineType   DisciplineType `gorm:"references:id"`

	Discipline_Reason     string
	Discipline_Punishment string
	Discipline_Point      uint
	Added_Time            time.Time
}
