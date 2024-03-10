package api

import (
	"database/sql/driver"
	"fmt"
	"net/http"

	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/query"
	"github.com/dirkarnez/stemexapi/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/sessions"
	"github.com/samber/lo"
	"gorm.io/gen"
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

func Register(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		type RegisterForm struct {
			UserName      string `json:"user_name"`
			Password      string `json:"password"`
			Email         string `json:"email"`
			ContactNumber string `json:"contact_number"`
		}
		var registerForm RegisterForm
		/*err := */ ctx.ReadJSON(&registerForm)

		var q = query.Use(dbInstance)
		err := q.Transaction(func(tx *query.Query) error {
			newUser := model.User{UserName: registerForm.UserName, Password: registerForm.Password, Email: registerForm.Email, ContactNumber: registerForm.ContactNumber, IsActivated: false}
			err := tx.User.Not(gen.Exists(tx.User.Where(tx.User.UserName.Eq(registerForm.UserName)))).Create(&newUser)
			if err != nil {
				return err
			}

			activationKey := model.NewUUIDEx().ToString()
			err = tx.ParentUserActivating.Create(&model.ParentUserActivating{UserID: &newUser.ID, ActivationKey: activationKey})
			if err != nil {
				return err
			}

			host := ctx.Request().Host
			if host == "localhost" {
				host = fmt.Sprintf("%s:5678", host)
			}

			err = utils.SendActivationHTMLEmail(newUser.Email, fmt.Sprintf("https://%s/activation?key=%s", host, activationKey))
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			ctx.StopWithError(http.StatusInternalServerError, err)
		} else {
			ctx.JSON(iris.Map{"status": iris.StatusOK})
		}
	}
}

func Activation(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		type ActivationForm struct {
			Key string `json:"key"`
		}
		var activationForm ActivationForm
		/*err := */ ctx.ReadJSON(&activationForm)

		var q = query.Use(dbInstance)
		err := q.Transaction(func(tx *query.Query) error {
			users, err := tx.User.
				LeftJoin(tx.ParentUserActivating, tx.User.ID.EqCol(tx.ParentUserActivating.UserID)).
				Where(tx.ParentUserActivating.ActivationKey.Eq(activationForm.Key)).Find()
			if err != nil {
				// invalid key
				return err
			}

			_, err = tx.User.
				Where(tx.User.ID.In(lo.Map(users, func(user *model.User, index int) driver.Valuer {
					return user.ID
				})...)).
				Update(tx.User.IsActivated, true)
			return err
		})
		if err != nil {
			ctx.StopWithError(http.StatusInternalServerError, err)
		} else {
			ctx.JSON(iris.Map{"status": iris.StatusOK})
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
