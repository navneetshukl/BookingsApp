package models

import "time"

//Reservation holds reservation data

type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

//Users is user model
type Users struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	Created     time.Time
	UpdatedAt   time.Time
}

//Rooms is room model
type Rooms struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Restrictions is restriction model
type Restrictions struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservations is the reservation model
type Reservations struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Rooms
}

// RoomRestrictions is the room restriction model

type RoomRestrictions struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	RoomIDs       int
	ReservationID int
	RestrictionID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Room          Rooms
	Reservation   Reservations
	Restriction   Restrictions
}
