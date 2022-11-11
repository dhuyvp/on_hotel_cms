package models

type HotelManage struct {
	HotelID           *int    `db:"HotelID,omitempty" json:"HotelID,omitempty"`
	Logo              *string `db:"Logo,omitempty" json:"Logo,omitempty"`
	HotelName         *string `db:"HotelName,omitempty" json:"HotelName,omitempty"`
	DevicesLimit      *int    `db:"DevicesLimit,omitempty" json:"DevicesLimit,omitempty"`
	DevicesNumber     *int    `db:"DevicesNumber,omitempty" json:"DevicesNumber,omitempty"`
	TotalRoom         *int    `db:"TotalRoom,omitempty" json:"TotalRoom,omitempty"`
	Address           *string `db:"Address,omitempty" json:"Address,omitempty"`
	DetailDescription *string `db:"DetailDescription,omitempty" json:"DetailDescription,omitempty"`
	IsActive          *bool   `db:"IsActive,omitempty" json:"IsActive,omitempty"`
}
