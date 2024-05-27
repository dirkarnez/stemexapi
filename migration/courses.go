package migration

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dirkarnez/stemexapi/bo"
	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/query"
	"github.com/dirkarnez/stemexapi/utils"
	"github.com/samber/lo"
)

func AddCourse(qOrTx *query.Query, s3 *utils.StemexS3Client,
	prefix, rootDir, parentID, iconFilePath, curriculumPlanFilePath string,
	blogs []dto.CurriculumCourseBlogEntries,
	youtube []dto.CurriculumCourseYoutubeVideoEntries,
	levels []dto.CurriculumCourseLevels,
) error {
	var lessonCount uint64 = 0
	for {
		// only do increment when exists
		lessonFolder := fmt.Sprintf(`%s\%s\Lesson %d`, prefix, rootDir, lessonCount+1)
		_, err := os.Stat(lessonFolder)
		if os.IsNotExist(err) {
			break
		} else {
			lessonCount++
		}
	}

	iconFile, err := utils.CreateMultipartFileHeader(fmt.Sprintf(`%s\%s\%s`, prefix, rootDir, iconFilePath))
	if err != nil {
		log.Println("?????????????????????????????")
		log.Fatalln(err)
	}

	curriculumPlanFile, err := utils.CreateMultipartFileHeader(fmt.Sprintf(`%s\%s\%s`, prefix, rootDir, curriculumPlanFilePath))
	if err != nil {
		log.Println("?????????????????????????????")
		log.Fatalln(err)
	}

	// files := []string{}

	lo.ForEach(levels, func(level dto.CurriculumCourseLevels, index int) {
		iconFile, err := utils.CreateMultipartFileHeader(fmt.Sprintf(`%s\%s\%s`, prefix, rootDir, level.IconPath))
		if err != nil {
			log.Fatalln(err)
		}

		level.IconFile = iconFile
		level.Lessons = lo.Map(make([]uint64, lessonCount), func(lessonNumber uint64, i int) dto.CurriculumCourseLevelLessons {
			var getFiles = func(folderName string) []dto.CurriculumCourseLevelLessonResources {
				filePaths, err := filepath.Glob(fmt.Sprintf(`%s\%s\Lesson %d\%s\*`, prefix, rootDir, lessonNumber, folderName))
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

			return dto.CurriculumCourseLevelLessons{
				LessonNumber:      lessonNumber,
				PresentationNotes: getFiles("Presentation Notes"),
				TeacherNotes:      getFiles("Teacher Notes"),
				StudentNotes:      getFiles("Student Notes"),
				MiscMaterials:     getFiles("Misc Materials"),
			}
		})
	})

	dtoInput := dto.CurriculumCourseForm{
		ParentID:            parentID,
		IconFile:            iconFile,
		CurriculumPlanFile:  curriculumPlanFile,
		BlogEntries:         blogs,
		YoutubeVideoEntries: youtube,
		Levels:              levels,
	}

	dtoOutput, err := bo.CreateOrUpdateCurriculumCourse(&dtoInput, s3, qOrTx)

	if err != nil {
		return err
	} else {
		fmt.Printf("%+v", dtoOutput)
	}

	return nil
}
