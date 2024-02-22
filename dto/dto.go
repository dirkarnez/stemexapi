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
	IsCourse    bool          `json:"is_course"`
}

type CurriculumCourseYoutubeVideoEntries struct {
	ID  string `json:"id" form:"id"`
	URL string `json:"url" form:"url"`
}

type CurriculumCourseLevels struct {
	ID          string                         `json:"id" form:"id"`
	Name        string                         `json:"name" form:"name"`
	IconID      string                         `json:"icon_id" form:"icon_id"`
	Description string                         `json:"description" form:"description"`
	Lessons     []CurriculumCourseLevelLessons `json:"lessons" form:"lessons"`
}

type CurriculumCourseLevelLessons struct {
	ID                string                                 `json:"id" form:"id"`
	LessonNumber      uint64                                 `json:"lesson_number" form:"lesson_number"`
	PresentationNotes []CurriculumCourseLevelLessonResources `json:"presentation_notes" form:"presentation_notes"`
	StudentNotes      []CurriculumCourseLevelLessonResources `json:"student_notes" form:"student_notes"`
	TeacherNotes      []CurriculumCourseLevelLessonResources `json:"teacher_notes" form:"teacher_notes"`
	MiscMaterials     []CurriculumCourseLevelLessonResources `json:"misc_materials" form:"misc_materials"`
}

type CurriculumCourseLevelLessonResources struct {
	ID         string `json:"id" form:"id"`
	ResourseID string `json:"resource_id" form:"resource_id"`
	Name       string `json:"name" form:"name"`
}

type CurriculumCourseBlogEntries struct {
	ID          string `json:"id" form:"id"`
	ExternalURL string `json:"external_url" form:"external_url"`
	Title       string `json:"title" form:"title"`
}

// type CurriculumCourseInformationEntries struct {
// 	ID      model.UUIDEx  `json:"id"`
// 	IconID  *model.UUIDEx `json:"icon_id"`
// 	Title   string        `json:"title"`
// 	Content string        `json:"content"`
// }

// type CurriculumCourseDetails struct {
// 	ID               model.UUIDEx                          `json:"id"`
// 	ParentID         *model.UUIDEx                         `json:"parent_id"`
// 	Description      string                                `json:"description"`
// 	IconID           *model.UUIDEx                         `json:"icon_id"`
// 	Prerequisites    []string                              `json:"prerequisites"`
// 	YoutubeVideoURLs []CurriculumCourseYoutubeVideoEntries `json:"youtube_video_entries"`
// 	BlogEntries      []CurriculumCourseBlogEntries         `json:"blog_entries"`
// 	// InformationEntries []CurriculumCourseInformationEntries  `json:"information_entries"`
// }

type File struct {
	ID               model.UUIDEx `json:"id"`
	SeqNo            uint64       `json:"seq_no"`
	FileNameUploaded string       `json:"file_name_uploaded"`
	//ContentHash      string `gorm:"column:content_hash;type:varchar(500);unique;not null"`
}

type FileManagement struct {
	Files              []File `json:"files"`
	FromSeqNoInclusive int64  `json:"from_seq_no_inclusive"`
	ToSeqNoExclusive   int64  `json:"to_seq_no_exclusive"`
	TotalCount         int64  `json:"total_count"`
}

type CurriculumCourseForm struct {
	ID                     string                                `form:"id" json:"id"`
	IconID                 string                                `form:"icon_id" json:"icon_id"`
	Description            string                                `form:"description" json:"description"`
	ParentID               string                                `form:"parent_id" json:"parent_id"`
	CourseID               string                                `form:"course_id" json:"course_id"`
	CurriculumPlanID       string                                `form:"curriculum_plan_id" json:"curriculum_plan_id"`
	CurriculumPlanFileName string                                `form:"curriculum_plan_file_name" json:"curriculum_plan_file_name"` // uploaded
	BlogEntries            []CurriculumCourseBlogEntries         `form:"blog_entries" json:"blog_entries"`
	YoutubeVideoEntries    []CurriculumCourseYoutubeVideoEntries `form:"youtube_video_entries" json:"youtube_video_entries"`
	Levels                 []CurriculumCourseLevels              `form:"levels" json:"levels"`
}
