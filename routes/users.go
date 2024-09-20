package routes

import (
	"errors"
	"net/http"

	"github.com/zedosoad1995/pokemon-wordle/models/user"
	route_types "github.com/zedosoad1995/pokemon-wordle/routes/types"
	"github.com/zedosoad1995/pokemon-wordle/utils"
	"gorm.io/gorm"
)

type CreateUserBody struct {
	Token string `json:"token"`
}

func createUserHandler(db *gorm.DB) route_types.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		body, err := utils.GetJSONBody[CreateUserBody](r)
		if err != nil {
			return err
		}

		userObj := user.User{
			Token: body.Token,
		}
		if err := db.Create(&userObj).Error; err != nil {
			if !errors.Is(err, gorm.ErrDuplicatedKey) {
				return utils.SendJSON(w, 400, route_types.ErrorRes{
					Message: "User already exists.",
				})
			}
			return err
		}

		return utils.SendJSON(w, 201, route_types.SuccessRes{Message: "User created"})
	}
}
