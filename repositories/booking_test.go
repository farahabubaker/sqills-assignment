package repositories

import (
	"testing"
	"ticketing-service/models"
)

type MockDB struct{}

func (m MockDB) Create(data any, tableName string) error {
	print("fake mock but it worked!")
	return nil
}

func Test_Create(t *testing.T) {
	mockdb := MockDB{}
	b := NewBookingRespoitory(mockdb)

	idk := models.Booking{Id: "howdy"}
	t.Log("howdy there")

	_, err := b.Create(&idk)
	if err != nil {
		t.Log("welp.....")
	}
}
