package model

import (
	"time"

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

var AllTables = []interface{}{
	&User{},
	&StudentToUser{},
	&ParentUserActivating{},
	&Role{},
	&File{},
	&UserActivity{},
	&CurriculumEntry{},
	&CurriculumCourse{},
	&CurriculumCoursePrerequisites{},
	&CurriculumCourseYoutubeVideoEntries{},
	&CurriculumCourseBlogEntries{},
	&CurriculumCourseLevel{},
	&CurriculumCourseLevelLesson{},
	&CurriculumCourseLessonResourceType{},
	&CurriculumCourseLevelLessonResources{},
}

type BaseModel struct {
	ID        UUIDEx         `gorm:"column:id;type:binary(16);primaryKey;default:UNHEX(REPLACE(UUID(), '-', ''))" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

type User struct {
	BaseModel
	FullName              string  `gorm:"column:full_name;type:varchar(255);not null" json:"full_name"`
	UserName              string  `gorm:"column:user_name;type:varchar(255);unique;not null" json:"user_name"`
	Password              string  `gorm:"column:password;type:varchar(15);not null" json:"-"`
	ContactNumberAreaCode string  `gorm:"column:contact_number_area_code;type:varchar(15);default:852" json:"contact_number_area_code"`
	ContactNumber         string  `gorm:"column:contact_number;type:varchar(500);not null" json:"contact_number"`
	Email                 string  `gorm:"column:email;type:varchar(255);unique;not null" json:"email"`
	IsDummy               bool    `gorm:"column:is_dummy;type:boolean;default:false" json:"is_dummy"`
	IsActivated           bool    `gorm:"column:is_activated;type:boolean;default:false" json:"is_activated"`
	RoleID                *UUIDEx `gorm:"column:role_id;type:binary(16)" json:"role_id"`
	Role                  *Role   `gorm:"foreignKey:RoleID" json:"role"`
}

// https://sheets.googleapis.com/v4/spreadsheets/1mRMBmxKuReGqp9MvcTiv-Z-QcxSDsHUHKwnxPORcj2Y/values/Form?key=AIzaSyBAuyTYKGijZn3jkwoMDlw0ZsR8JR5iOno
type StudentToUser struct {
	BaseModel
	GoogleSheetUserName string `gorm:"column:google_sheet_user_name;type:varchar(255);unique;not null" json:"-"`
	GoogleSheetPassword string `gorm:"column:google_sheet_password;type:varchar(255);unique;not null" json:"-"`
	GoogleSheetSID      string `gorm:"column:google_sheet_sid;type:varchar(255);unique;not null" json:"-"` // "SID"

	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"` //"Student Name"

	UserID UUIDEx `gorm:"column:user_id;type:binary(16);not null" json:"-"`
	User   User   `gorm:"foreignKey:UserID;not null" json:"-"`
}

type ParentUserActivating struct {
	BaseModel
	UserID        *UUIDEx `gorm:"column:user_id;type:binary(16)"`
	User          *User   `gorm:"foreignKey:UserID"`
	ActivationKey string  `gorm:"column:activation_key;type:varchar(255);unique;not null" json:"activation_key"`
	//UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type Role struct {
	BaseModel
	Name string `gorm:"column:name;unique;not null" json:"name"`
}

type UserActivity struct {
	BaseModel         //CreatedAt = login time, UpdatedAt = refresh session time
	UserID    *UUIDEx `gorm:"column:user_id;type:binary(16)"`
	User      *User   `gorm:"foreignKey:UserID"`
}

type File struct {
	BaseModel
	IsPublic         bool   `gorm:"column:is_public;type:boolean;default:false;not null"`
	SeqNo            uint64 `gorm:"column:seq_no;unique;not null;autoIncrement"`
	ObjectKey        string `gorm:"column:object_key;type:varchar(500);unique;not null"`
	FileNameUploaded string `gorm:"column:file_name_uploaded;type:varchar(500);not null"`
	//ContentHash      string `gorm:"column:content_hash;type:varchar(500);"`
}

type CurriculumEntry struct {
	BaseModel
	IconID         UUIDEx  `gorm:"column:icon_id;type:binary(16);not null"`
	Icon           File    `gorm:"foreignKey:IconID"` //constraint:OnDelete:SET NULL
	Description    string  `gorm:"column:description;type:varchar(500);not null;"`
	ParentID       *UUIDEx `gorm:"column:parent_id;type:binary(16);"`
	SeqNoSameLevel uint64  `gorm:"column:seq_no_same_level;not null;default:0;"`
}

type CurriculumCourse struct {
	BaseModel
	EntryID          UUIDEx          `gorm:"column:entry_id;type:binary(16);unique;not null"`
	Entry            CurriculumEntry `gorm:"foreignKey:EntryID"`
	CurriculumPlanID *UUIDEx         `gorm:"column:curriculum_plan_id;type:binary(16);"`
	CurriculumPlan   *File           `gorm:"foreignKey:CurriculumPlanID"` //constraint:OnDelete:SET NULL
}

type CurriculumCourseLevel struct {
	BaseModel
	Name        string           `gorm:"column:name;not null"`
	IconID      UUIDEx           `gorm:"column:icon_id;type:binary(16);not null"`
	Icon        File             `gorm:"foreignKey:IconID"` //constraint:OnDelete:SET NULL
	Title       string           `gorm:"column:content;type:varchar(255);not null"`
	Description string           `gorm:"column:description;type:varchar(1000);not null"`
	CourseID    UUIDEx           `gorm:"column:course_id;type:binary(16);not null"`
	Course      CurriculumCourse `gorm:"foreignKey:CourseID"`
}

type CurriculumCourseLevelLesson struct {
	BaseModel
	LessonNumber  uint64                `gorm:"column:lesson_number;not null"`
	CourseLevelID UUIDEx                `gorm:"column:course_level_id;type:binary(16);not null"`
	CourseLevel   CurriculumCourseLevel `gorm:"foreignKey:CourseLevelID"`
	// Content string           `gorm:"column:content;type:varchar(255);not null"`
}

type CurriculumCourseLessonResourceType struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(255);not null"`
	// SeqNo            uint64 `gorm:"column:seq_no;unique;not null;autoIncrement"`
	// ObjectKey        string `gorm:"column:object_key;type:varchar(500);unique;not null"`
	// FileNameUploaded string `gorm:"column:file_name_uploaded;type:varchar(500);not null"`
	//ContentHash      string `gorm:"column:content_hash;type:varchar(500);"`
}

type CurriculumCourseLevelLessonResources struct {
	BaseModel
	LessonID       UUIDEx                             `gorm:"column:lesson_id;type:binary(16);not null"`
	Lesson         CurriculumCourseLevelLesson        `gorm:"foreignKey:LessonID"`
	ResourseTypeID UUIDEx                             `gorm:"column:resourse_type_id;type:binary(16);not null"`
	ResourseType   CurriculumCourseLessonResourceType `gorm:"foreignKey:ResourseTypeID"` //constraint:OnDelete:SET NULL
	ResourseID     UUIDEx                             `gorm:"column:resourse_id;type:binary(16);not null"`
	Resourse       File                               `gorm:"foreignKey:ResourseID"` //constraint:OnDelete:SET NULL
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

// type CurriculumCourseInformationEntries struct {
// 	BaseModel
// 	IconID   *UUIDEx                      `gorm:"column:icon_id;type:binary(16)"`
// 	Icon     *File                        `gorm:"foreignKey:IconID"` //constraint:OnDelete:SET NULL
// 	Title    string                       `gorm:"column:title;type:varchar(255);not null"`
// 	Content  string                       `gorm:"column:content;type:varchar(1000);not null"`
// 	LessonID *UUIDEx                      `gorm:"column:lesson_id;type:binary(16)"`
// 	Lesson   *CurriculumCourseLevelLesson `gorm:"foreignKey:LessonID"`
// }
