package dto

import (
	"time"

	"github.com/dirkarnez/stemexapi/model"
)

type UserActivityResult struct {
	Count    int        `gorm:"column:count" json:"count"`
	UserName string     `gorm:"column:user_name" json:"user_name"`
	LoginAt  *time.Time `gorm:"column:login_at" json:"login_at,omitempty"`
}

// type UserActivityPerDayResult struct {
// 	Count    int       `gorm:"column:count" json:"count"`
// 	UserName string    `gorm:"column:user_name" json:"user_name"`
// 	LoginAt  time.Time `gorm:"column:login_at" json:"login_at"`
// }

type CurriculumEntry struct {
	ID          model.UUIDEx  `json:"id"`
	Description string        `json:"description"`
	ParentID    *model.UUIDEx `json:"parent_id"`
}

type CurriculumCourseBlogEntries struct {
	ExternalURL string `json:"external_url"`
	Title       string `json:"title"`
}

type CurriculumCourseInformationEntries struct {
	ImageSrc string `json:"image_src"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type CurriculumCourseDetails struct {
	ID                 model.UUIDEx                         `json:"id"`
	Description        string                               `json:"description"`
	Prerequisites      []string                             `json:"prerequisites"`
	YoutubeVideoURLs   []string                             `json:"youtube_video_urls"`
	BlogEntries        []CurriculumCourseBlogEntries        `json:"blog_entries"`
	InformationEntries []CurriculumCourseInformationEntries `json:"information_entries"`
}
