package models

type ChannelPackManage struct {
	ChannelPackID     *int    `db:"ChannelPackID,omitempty" json:"ChannelPackID,omitempty"`
	HotelID           *int    `db:"HotelID,omitempty" json:"HotelID,omitempty"`
	Logo              *string `db:"Logo,omitempty" json:"Logo,omitempty"`
	ChannelPackName   *string `db:"ChannelPackName,omitempty" json:"ChannelPackName,omitempty"`
	DetailDescription *string `db:"DetailDescription,omitempty" json:"DetailDescription,omitempty"`
	Note              *string `db:"Note,omitempty" json:"Note,omitempty"`
	IsActive          *bool   `db:"IsActive,omitempty" json:"IsActive,omitempty"`
}
