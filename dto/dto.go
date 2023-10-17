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
	IconID      *model.UUIDEx `json:"icon_id"`
}

type CurriculumCourseYoutubeVideoEntries struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

type CurriculumCourseBlogEntries struct {
	ExternalURL string `json:"external_url"`
	Title       string `json:"title"`
}

type CurriculumCourseInformationEntries struct {
	IconID  *model.UUIDEx `json:"icon_id"`
	Title   string        `json:"title"`
	Content string        `json:"content"`
}

type CurriculumCourseDetails struct {
	ID                 model.UUIDEx                          `json:"id"`
	Description        string                                `json:"description"`
	IconID             *model.UUIDEx                         `json:"icon_id"`
	Prerequisites      []string                              `json:"prerequisites"`
	YoutubeVideoURLs   []CurriculumCourseYoutubeVideoEntries `json:"youtube_video_entries"`
	BlogEntries        []CurriculumCourseBlogEntries         `json:"blog_entries"`
	InformationEntries []CurriculumCourseInformationEntries  `json:"information_entries"`
}

type File struct {
	ID    model.UUIDEx `json:"id"`
	SeqNo uint64       `json:"seq_no"`
	//ContentHash      string `gorm:"column:content_hash;type:varchar(500);unique;not null"`
}

type FileManagement struct {
	Files              []File `json:"files"`
	FromSeqNoInclusive int64  `json:"from_seq_no_inclusive"`
	ToSeqNoExclusive   int64  `json:"to_seq_no_exclusive"`
	TotalCount         int64  `json:"total_count"`
}
