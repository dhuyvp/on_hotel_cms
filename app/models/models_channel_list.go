package models

type ChannelListManage struct {
	ChannelListID     *int    `db:"ChannelListID,omitempty" json:"ChannelListID,omitempty"`
	ChannelPackID     *int    `db:"ChannelPackID,omitempty" json:"ChannelPackID,omitempty"`
	ChannelListName   *string `db:"ChannelListName,omitempty" json:"ChannelListName,omitempty"`
	DetailDescription *string `db:"DetailDescription,omitempty" json:"DetailDescription,omitempty"`
	SortOrder         *int    `db:"SortOrder,omitempty" json:"SortOrder,omitempty"`
	CreatedAt         *string `db:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	UpdatedAt         *string `db:"UpdatedAt,omitempty" json:"UpdatedAt,omitempty"`
}
