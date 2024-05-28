package migration

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/dirkarnez/stemexapi/bo"
	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/query"
	"github.com/dirkarnez/stemexapi/utils"
	"github.com/samber/lo"
)

func AddCourse(qOrTx *query.Query, s3 *utils.StemexS3Client,
	prefix, rootDir, parentID, description, iconFilePath string, curriculumPlanFilePath *string,
	blogs []dto.CurriculumCourseBlogEntries,
	youtube []dto.CurriculumCourseYoutubeVideoEntries,
	levels []dto.CurriculumCourseLevels,
) (*dto.CurriculumCourseForm, error) {
	// coursesRoot := fmt.Sprintf(`%s\%s\*`, prefix, rootDir)
	// coursesRootFolders, err := filepath.Glob(coursesRoot)
	// if err != nil {
	// 	log.Println("?????????????????????????????")
	// 	log.Fatalln(err)
	// }

	// a := filepath.Base(coursesRootFolders[0])
	// fmt.Println(a)

	iconFile, err := utils.CreateMultipartFileHeader(fmt.Sprintf(`%s\%s\%s`, prefix, rootDir, iconFilePath))
	if err != nil {
		log.Println("?????????????????????????????")
		log.Fatalln(err)
	}

	var curriculumPlanFile *multipart.FileHeader = nil
	if curriculumPlanFilePath != nil {
		curriculumPlanFile, err = utils.CreateMultipartFileHeader(fmt.Sprintf(`%s\%s\%s`, prefix, rootDir, *curriculumPlanFilePath))
		if err != nil {
			log.Println("?????????????????????????????")
			log.Fatalln(err)
		}
	}

	// files := []string{}

	for i := range levels {
		var lessonCount uint64 = 0
		for {
			// only do increment when exists
			lessonFolder := fmt.Sprintf(`%s\%s\%s\Lesson %d`, prefix, rootDir, levels[i].Name, lessonCount+1)
			_, err := os.Stat(lessonFolder)
			if os.IsNotExist(err) {
				break
			} else {
				lessonCount++
			}
		}

		iconFilePath := fmt.Sprintf(`%s\%s\%s\%s`, prefix, rootDir, levels[i].Name, levels[i].IconPath)
		levels[i].IconFile, err = utils.CreateMultipartFileHeader(iconFilePath)
		if err != nil {
			log.Fatalln(err)
		}

		levels[i].Lessons = func(j int) []dto.CurriculumCourseLevelLessons {
			return lo.Map(make([]uint64, lessonCount), func(_ uint64, i int) dto.CurriculumCourseLevelLessons {
				var lessonNumber uint64 = uint64(i) + 1

				var getFiles = func(folderName string) []dto.CurriculumCourseLevelLessonResources {
					dir := fmt.Sprintf(`%s\%s\%s\Lesson %d\%s\*`, prefix, rootDir, levels[j].Name, lessonNumber, folderName)
					filePaths, err := filepath.Glob(dir)
					if err != nil {
						log.Fatal(err)
					}

					return lo.Map(filePaths, func(filePath string, i int) dto.CurriculumCourseLevelLessonResources {
						file, err := utils.CreateMultipartFileHeader(filePath)
						if err != nil {
							log.Fatalln(err)
						}

						return dto.CurriculumCourseLevelLessonResources{
							File: file,
						}
					})
				}

				dto := dto.CurriculumCourseLevelLessons{
					LessonNumber:      lessonNumber,
					PresentationNotes: getFiles("Presentation Notes"),
					TeacherNotes:      getFiles("Teacher Notes"),
					StudentNotes:      getFiles("Student Notes"),
					MiscMaterials:     getFiles("Misc Materials"),
				}

				return dto
			})
		}(i)
	}

	dtoInput := dto.CurriculumCourseForm{
		ParentID:            parentID,
		IconFile:            iconFile,
		Description:         description,
		CurriculumPlanFile:  curriculumPlanFile,
		BlogEntries:         blogs,
		YoutubeVideoEntries: youtube,
		Levels:              levels,
	}

	return bo.CreateOrUpdateCurriculumCourse(&dtoInput, s3, qOrTx)
}
