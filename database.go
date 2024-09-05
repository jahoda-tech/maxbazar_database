package database

import (
	"database/sql"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName      string `gorm:"uniqueIndex:unique_user"`
	SecondName     string `gorm:"uniqueIndex:unique_user"`
	Password       string
	ExternalID     string
	LastAccessTime time.Time
	AverageRating  float64
	Location       string
	Email          string `gorm:"uniqueIndex:unique_user"`
	Phone          string
	Note           string
}

type UserReviewRecord struct {
	gorm.Model
	Heading   string
	Text      string
	UserInID  int `gorm:"uniqueIndex:unique_user_review_record"`
	UserIn    User
	UserOutID int `gorm:"uniqueIndex:unique_user_review_record"`
	UserOut   User
	DateTime  time.Time `gorm:"uniqueIndex:unique_user_review_record"`
	Rating    int
	Note      string
}

type UserFilterRecord struct {
	gorm.Model
	Name   string `gorm:"uniqueIndex:unique_user_filter_record"`
	UserID int    `gorm:"uniqueIndex:unique_user_filter_record"`
	User   User
	Data   datatypes.JSON
	Note   string
}

type ItemType struct {
	gorm.Model
	Name             string `gorm:"uniqueIndex:unique_item_type"`
	ParentItemTypeID int
	Data             datatypes.JSON
	Note             string
}

type PriceType struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_price_type"`
	Data datatypes.JSON
	Note string
}

type Item struct {
	gorm.Model
	Heading             string `gorm:"uniqueIndex:unique_item"`
	Text                string
	UserID              int `gorm:"uniqueIndex:unique_item"`
	User                User
	ItemTypeID          int
	ItemType            ItemType
	PriceTypeID         int
	PriceType           PriceType
	ItemData            datatypes.JSON
	Price               float64
	DateTime            time.Time
	DeleteAfter         sql.NullTime
	DecreasePriceAfter  sql.NullTime
	DecreasePriceAmount float64
	Description         string
	Location            string
	TemporaryPassword   string
	Public              bool
	Note                string
}

type UserMessageRecord struct {
	gorm.Model
	UserInID  int `gorm:"uniqueIndex:unique_user_message_record"`
	UserIn    User
	UserOutID int `gorm:"uniqueIndex:unique_user_message_record"`
	UserOut   User
	ItemID    int `gorm:"uniqueIndex:unique_user_message_record"`
	Item      Item
	DateTime  time.Time `gorm:"uniqueIndex:unique_user_message_record"`
	Message   string
	Note      string
}

type UserInvoiceRecord struct {
	gorm.Model
	UserID   int `gorm:"uniqueIndex:unique_user_invoice_record"`
	User     User
	ItemID   int `gorm:"uniqueIndex:unique_user_invoice_record"`
	Item     Item
	DateTime time.Time `gorm:"uniqueIndex:unique_user_invoice_record"`
	Text     string
	Note     string
}

type ItemCounterRecord struct {
	gorm.Model
	ItemID   int `gorm:"uniqueIndex:unique_item_counter_record"`
	Item     Item
	DateTime time.Time
	Count    int
	Note     string
}

type ItemImageRecord struct {
	gorm.Model
	ItemID   int `gorm:"uniqueIndex:unique_item_image_record"`
	Item     Item
	DateTime time.Time
	Url      string `gorm:"uniqueIndex:unique_item_image_record"`
	Note     string
}

type ItemAlarmRecord struct {
	gorm.Model
	UserID        int `gorm:"uniqueIndex:unique_user_alarm_record"`
	User          User
	ItemTypeID    int `gorm:"uniqueIndex:unique_user_alarm_record"`
	ItemType      ItemType
	ItemData      datatypes.JSON
	StartDateTime time.Time
	EndDateTime   time.Time
	Note          string
}

type ItemWatchlistRecord struct {
	gorm.Model
	UserID int `gorm:"uniqueIndex:unique_user_watchlist_record"`
	User   User
	ItemID int `gorm:"uniqueIndex:unique_user_watchlist_record"`
	Item   Item
	Note   string
}
