package models

type ChannelManage struct {
	ChannelID         *int    `db:"ChannelID,omitempty" json:"ChannelID,omitempty"`
	ChannelListID     *int    `db:"ChannelListID,omitempty" json:"ChannelListID,omitempty"`
	Logo              *string `db:"Logo,omitempty" json:"Logo,omitempty"`
	ChannelName       *string `db:"ChannelName,omitempty" json:"ChannelName,omitempty"`
	Link              *string `db:"Link,omitempty" json:"Link,omitempty"`
	DetailDescription *string `db:"DetailDescription,omitempty" json:"DetailDescription,omitempty"`
	IsActive          *bool   `db:"IsActive,omitempty" json:"IsActive,omitempty"`
	CreatedAt         *string `db:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
}
