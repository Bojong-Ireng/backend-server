package domain

import (
	"time"

	"gorm.io/datatypes"
)

type User struct {
	ID           string
	Username     string
	TotalPoints  int64
	PhoneNumber  string
	RegisteredOn datatypes.Date
}

type LoginCredential struct {
	ID       string
	Email    string
	Password string
}

type GoogleCredential struct {
	ID    string
	Email string
}

type WalkTrackHistory struct {
	ID               uint
	Path             datatypes.JSON `gorm:"type:jsonb"`
	DistanceTraveled float32
	Duration         time.Duration `gorm:"type:interval"`
	Timestamp        time.Time
	UserID           string
}

type BusTrackHistory struct {
	ID               uint
	Path             datatypes.JSON `gorm:"type:jsonb"`
	DistanceTraveled float32
	Duration         time.Duration `gorm:"type:interval"`
	Timestamp        time.Time
	UserID           string
	BusID            string
}

// Probably remove json tag later, its only for testing purpose.
type Bus struct {
	BusID       string `gorm:"primaryKey" json:"bus_id"`
	BusName     string `json:"bus_name"`
	CompanyName string `json:"company_name"`
}
