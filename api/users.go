package api

import (
	"github.com/dirkarnez/stemexapi/bo"
	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateUser(userBO *bo.UsersBO) context.Handler {
	return func(ctx iris.Context) {
		userBO.CreateUser()
	}
}

func GetAllUsers(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		var users []model.User
		if err := dbInstance.Model(&model.User{}).Preload(clause.Associations).Find(&users).Error; err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		} else {
			ctx.JSON(users)
		}
	}
}

func GetAllRoles(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		var roles []model.Role
		if err := dbInstance.Model(&model.Role{}).Find(&roles).Error; err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		} else {
			ctx.JSON(roles)
		}
	}
}

func GetParentActivities(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		var userActivityResult []dto.UserActivityResult
		if err := dbInstance.Model(&model.User{}).
			Select("COUNT(`ua`.`id`) as `count`, `users`.`user_name`, CAST(`ua`.`created_at` AS DATE) as `login_at`").
			Joins("LEFT JOIN `roles` as `r` ON `users`.`role_id` = `r`.`id`").
			Joins("LEFT JOIN `user_activities` as `ua` ON `ua`.`user_id` = `users`.`id`").
			Group("`login_at`, `users`.`user_name`").
			Where("`r`.`name` IN ?", []string{"parent"}).
			Scan(&userActivityResult).Error; err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		} else {
			ctx.JSON(userActivityResult)
		}
	}
}

func GetProspectActivities(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		var userActivityResult []dto.UserActivityResult
		if err := dbInstance.Model(&model.User{}).
			Select("COUNT(`ua`.`id`) as `count`, `users`.`user_name`, CAST(`ua`.`created_at` AS DATE) as `login_at`").
			Joins("LEFT JOIN `roles` as `r` ON `users`.`role_id` = `r`.`id`").
			Joins("LEFT JOIN `user_activities` as `ua` ON `ua`.`user_id` = `users`.`id`").
			Group("`login_at`, `users`.`user_name`").
			Where("`r`.`name` IS NULL").
			Scan(&userActivityResult).Error; err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		} else {
			ctx.JSON(userActivityResult)
		}
	}
}

func GetInternalUserActivities(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		var userActivityResult []dto.UserActivityResult
		if err := dbInstance.Model(&model.User{}).
			Select("COUNT(`ua`.`id`) as `count`, `users`.`user_name`, CAST(`ua`.`created_at` AS DATE) as `login_at`").
			Joins("LEFT JOIN `roles` as `r` ON `users`.`role_id` = `r`.`id`").
			Joins("LEFT JOIN `user_activities` as `ua` ON `ua`.`user_id` = `users`.`id`").
			Group("`login_at`, `users`.`user_name`").
			Where("`r`.`name` IN ?", []string{"sales", "instructor", "admin"}).
			Scan(&userActivityResult).Error; err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		} else {
			ctx.JSON(userActivityResult)
		}
	}
}

func CreateOrUpdateUser(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		type CreateOrUpdateUserDTO struct {
			UserID   model.UUIDEx  `json:"user_id"`
			UserName string        `json:"user_name"`
			RoleID   *model.UUIDEx `json:"role_id"`
		}

		createOrUpdateUserDTO := CreateOrUpdateUserDTO{}
		/*err := */ ctx.ReadJSON(&createOrUpdateUserDTO)

		err := dbInstance.Transaction(func(tx *gorm.DB) error {
			// do some database operations in the transaction (use 'tx' from this point, not 'db')
			var userToUpdateOrCreate model.User

			var allZero = func(s []byte) bool {
				for _, v := range s {
					if v != 0 {
						return false
					}
				}
				return true
			}

			if !allZero(createOrUpdateUserDTO.UserID[:]) {
				if err := tx.First(&userToUpdateOrCreate, "id = ?", createOrUpdateUserDTO.UserID).Error; err != nil {
					return err
				}
			} else {
				userToUpdateOrCreate.FullName = createOrUpdateUserDTO.UserName
				userToUpdateOrCreate.Password = "stemex"
			}

			userToUpdateOrCreate.UserName = createOrUpdateUserDTO.UserName
			userToUpdateOrCreate.RoleID = createOrUpdateUserDTO.RoleID

			if err := tx.Save(&userToUpdateOrCreate).Error; err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
		} else {
			ctx.JSON(iris.Map{})
		}
	}
}
