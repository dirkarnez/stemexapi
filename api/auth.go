package api

import (
	"fmt"

	"github.com/dirkarnez/stemexapi/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/sessions"
	"gorm.io/gorm"
)

func Login(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		type LoginForm struct {
			UserName string `json:"user_name"`
			Password string `json:"password"`
		}
		var loginForm LoginForm
		/*err := */ ctx.ReadJSON(&loginForm)

		var authenticate = func(userName, password string) bool {
			var users []model.User

			err := dbInstance.Transaction(func(tx *gorm.DB) error {
				// do some database operations in the transaction (use 'tx' from this point, not 'db')
				if err := tx.Model(&model.User{}).
					Where(&model.User{UserName: userName, Password: password, IsActivated: true}).
					Limit(2).
					Group("user_name").
					Having("COUNT(user_name) = 1").
					Find(&users).Error; err != nil {
					// return any error will rollback
					return err
				}

				if len(users) > 1 {
					return fmt.Errorf("internal error")
				}

				userID := users[0].ID
				if len(users) == 1 {
					if err := tx.Create(&model.UserActivity{UserID: &userID}).Error; err != nil {
						return err
					}
				}

				// return nil will commit the whole transaction
				return nil
			})

			if err != nil {
				return false
			}

			return len(users) == 1
		}

		// Authenticate user
		if authenticate(loginForm.UserName, loginForm.Password) {
			sessions.Get(ctx).Set("user_name", loginForm.UserName)
			sessions.Get(ctx).Set("authenticated", true)

			ctx.JSON(iris.Map{"status": iris.StatusOK})
		} else {
			ctx.StatusCode(iris.StatusUnauthorized)
		}
	}
}

func Logout(ctx iris.Context) {
	session := sessions.Get(ctx)

	if session != nil {
		// Revoke users authentication
		// session.Set("authenticated", false)
		// Or to remove the variable:
		// session.Delete("authenticated")
		// Or destroy the whole session:
		session.Destroy()
	}

	ctx.JSON(iris.Map{"status": iris.StatusOK})
}
