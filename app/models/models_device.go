package models

type DeviceManage struct {
	DeviceID          *int    `db:"DeviceID,omitempty" json:"DeviceID,omitempty"`
	HotelID           *int    `db:"HotelID,omitempty" json:"HotelID,omitempty"`
	MacWired          *string `db:"MacWired,omitempty" json:"MacWired,omitempty"`
	MacWireless       *string `db:"MacWireless,omitempty" json:"MacWireless,omitempty"`
	DeviceName        *string `db:"DeviceName,omitempty" json:"DeviceName,omitempty"`
	DetailDescription *string `db:"DetailDescription,omitempty" json:"DetailDescription,omitempty"`
	IsActive          *bool   `db:"IsActive,omitempty" json:"IsActive,omitempty"`
	CreatedAt         *string `db:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	UpdatedAt         *string `db:"UpdatedAt,omitempty" json:"UpdatedAt,omitempty"`
}
