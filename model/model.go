package model

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CREATE TABLE login(
//    ,user_name VARCHAR (15) NOT NULL PRIMARY KEY
//    , VARCHAR() NOT NULL
//    ,sid VARCHAR(8)
//    ,student_name VARCHAR(255)
//    ,parent_telephone VARCHAR(15)
//    ,email VARCHAR(255)
//    ,access BOOLEAN
//    ,continue_id VARCHAR(60)
//    ,
//  );

type UUIDEx uuid.UUID

// GormDataType -> sets type to binary(16)
func (my UUIDEx) GormDataType() string {
	return "binary(16)"
}

// Scan --> From DB
func (my *UUIDEx) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*my = UUIDEx(parseByte)
	return err
}

// Value -> TO DB
func (my UUIDEx) Value() (driver.Value, error) {
	return uuid.UUID(my).MarshalBinary()
}

// func NewUUIDEx() UUIDEx {
// 	return UUIDEx(uuid.UUID{})
// }

type BaseModel struct {
	ID        UUIDEx         `gorm:"column:id;type:binary(16);primaryKey;default:UNHEX(REPLACE(UUID(), '-', ''))"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

type User struct {
	BaseModel
	FullName      string  `gorm:"column:full_name;type:varchar(255);not null"`
	UserName      string  `gorm:"column:user_name;type:varchar(15);unique;not null"`
	Password      string  `gorm:"column:password;type:varchar(15);not null"`
	ContactNumber string  `gorm:"column:contact_number;type:varchar(15);not null"`
	Email         string  `gorm:"column:email;type:varchar(255);not null"`
	IsDummy       bool    `gorm:"column:is_dummy;type:boolean;default:false"`
	RoleID        *UUIDEx `gorm:"column:role_id;type:binary(16)"`
	Role          *Role   `gorm:"foreignKey:RoleID"`
}

type Role struct {
	BaseModel
	Name string `gorm:"column:name;unique;not null"`
}

type UserActivity struct {
	BaseModel         //CreatedAt = login time, UpdatedAt = refresh session time
	UserID    *UUIDEx `gorm:"column:user_id;type:binary(16)"`
	User      *User   `gorm:"foreignKey:UserID"`
}

type File struct {
	BaseModel
	SeqNo            uint64 `gorm:"column:seq_no;unique;not null;autoIncrement"`
	PhysicalFileName string `gorm:"column:physical_file_name;type:varchar(500);unique;not null"`
	//ContentHash      string `gorm:"column:content_hash;type:varchar(500);unique;not null"`
}

type CurriculumEntry struct {
	BaseModel
	IconID      *UUIDEx `gorm:"column:icon_id;type:binary(16)"`
	Icon        *File   `gorm:"foreignKey:IconID"` //constraint:OnDelete:SET NULL
	Description string  `gorm:"column:description;type:varchar(255);unique;not null"`
	ParentID    *UUIDEx `gorm:"column:parent_id;type:binary(16)"`
}

type CurriculumCoursePrerequisites struct {
	BaseModel
	Content string           `gorm:"column:content;type:varchar(255);not null"`
	EntryID *UUIDEx          `gorm:"column:entry_id;type:binary(16)"`
	Entry   *CurriculumEntry `gorm:"foreignKey:EntryID"`
}

type CurriculumCourseYoutubeVideoEntries struct {
	BaseModel
	URL     string           `gorm:"column:url;type:varchar(500);not null"`
	Title   string           `gorm:"column:title;type:varchar(255);not null"`
	EntryID *UUIDEx          `gorm:"column:entry_id;type:binary(16)"`
	Entry   *CurriculumEntry `gorm:"foreignKey:EntryID"`
}

type CurriculumCourseBlogEntries struct {
	BaseModel
	ExternalURL string           `gorm:"column:external_url;type:varchar(500);not null"`
	Title       string           `gorm:"column:title;type:varchar(255);not null"`
	EntryID     *UUIDEx          `gorm:"column:entry_id;type:binary(16)"`
	Entry       *CurriculumEntry `gorm:"foreignKey:EntryID"`
}

type CurriculumCourseInformationEntries struct {
	BaseModel
	IconID  *UUIDEx          `gorm:"column:icon_id;type:binary(16)"`
	Icon    *File            `gorm:"foreignKey:IconID"` //constraint:OnDelete:SET NULL
	Title   string           `gorm:"column:title;type:varchar(255);not null"`
	Content string           `gorm:"column:content;type:varchar(1000);not null"`
	EntryID *UUIDEx          `gorm:"column:entry_id;type:binary(16)"`
	Entry   *CurriculumEntry `gorm:"foreignKey:EntryID"`
}
