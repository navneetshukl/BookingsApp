package dbrepo

import (
	"bookings-udemy/internal/models"
	"errors"
	"time"
)

func (m *testDBRepo) AllUsers() bool {

	return true
}
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {

	return 1, nil
}

// InsertRoomRestriction Inserts the room restriction in database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {

	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if room availability exists  for roomID else it returns false
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {

	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available room if any,for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	var rooms []models.Room
	return rooms, nil

}

// GetRoomByID gets a room by id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {

	var room models.Room
	if id > 2 {
		return room, errors.New("Some error")

	}

	return room, nil
}

func (m *testDBRepo) GetUserByID(id int) (models.User,error){
	var u models.User
	return u,nil
}

func(m *testDBRepo) UpdateUser(u models.User) error{
	return nil;
}

func (m* testDBRepo) Authenticate(email,testPassword string) (int,string,error){
	return 1,"",nil
}
