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
	Currency       string
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
	ParentItemTypeID int    `gorm:"uniqueIndex:unique_item_type"`
	CategoryDepth    int    `gorm:"uniqueIndex:unique_item_type"`
	Data             datatypes.JSON
	Url              string `gorm:"uniqueIndex:unique_item_type"`
	Count            int
	Enabled          bool `gorm:"default:true"`
	FinalItemTypeIds string
	Note             string
}

type Item struct {
	gorm.Model
	Heading             string `gorm:"uniqueIndex:unique_item"`
	Description         string
	UserID              int `gorm:"uniqueIndex:unique_item"`
	User                User
	TemporaryEmail      string
	TemporaryPhone      string
	ItemTypeID          int
	ItemType            ItemType
	ItemData            datatypes.JSON `gorm:"type:jsonb"`
	Price               float64
	DateTime            time.Time
	DeleteAfter         sql.NullTime
	DecreasePriceAfter  sql.NullTime
	DecreasePriceAmount float64
	Location            string `gorm:"index:item_location"`
	TemporaryPassword   string
	Currency            string
	Public              bool
	Premium             bool
	PremiumDateTime     time.Time
	PremiumLimitCount   int
	PremiumLimit        bool
	VIN                 string
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

type InvoiceRecord struct {
	gorm.Model
	PayerEmail   string
	PayerName    string
	PremiumCount int
	ItemID       int
	Item         Item
	DateTime     time.Time
	Amount       string
	Currency     string
	Note         string
	InvoiceData  datatypes.JSON
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
	UserID        int `gorm:"uniqueIndex:unique_user_alarm_record,priority:2"`
	User          User
	ItemTypeID    int `gorm:"uniqueIndex:unique_user_alarm_record,priority:1"`
	ItemType      ItemType
	ItemData      datatypes.JSON
	StartDateTime time.Time
	EndDateTime   time.Time
	Note          string
}

type WatchlistType struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:unique_watchlist_type"`
	Note string
}

type ItemWatchlistRecord struct {
	gorm.Model
	UserID            int `gorm:"uniqueIndex:unique_user_item_watchlist_record;uniqueIndex:unique_user_data_watchlist_record"`
	User              User
	ItemID            sql.NullInt32 `gorm:"uniqueIndex:unique_user_item_watchlist_record"`
	Item              Item
	WatchlistTypeId   sql.NullInt32 `gorm:"uniqueIndex:unique_user_item_watchlist_record;uniqueIndex:unique_user_data_watchlist_record"`
	WatchlistType     WatchlistType
	ItemTypeId        sql.NullInt32 `gorm:"uniqueIndex:unique_user_data_watchlist_record"`
	ItemType          ItemType
	EmailSent         bool
	EmailSentDateTime time.Time
	Data              datatypes.JSON
	Note              string
}

type Setting struct {
	gorm.Model
	Name    string `gorm:"uniqueIndex:unique_settings"`
	Value   string
	Enabled bool `gorm:"default:true"`
	Note    string
}

type Ruian struct {
	gorm.Model
	PostalCode string `gorm:"uniqueIndex:unique_ruian"`
	Location   string `gorm:"uniqueIndex:unique_ruian"`
	Longitude  float64
	Latitude   float64
	Note       string
}

type ItemHistoryRecord struct {
	gorm.Model
	ItemID   int
	Item     Item
	DateTime time.Time
	Count    int
	Note     string
}

type MaxbazarVisitRecords struct {
	gorm.Model
	DateTime string `gorm:"uniqueIndex:unique_maxbazar_visit_records"`
	Count    int
}

type RoadmapRecords struct {
	gorm.Model
	Heading string `gorm:"uniqueIndex:unique_roadmap_records"`
	Text    string
	Type    string
	Count   int `gorm:"default:0"`
	UserIds string
	Note    string
}
