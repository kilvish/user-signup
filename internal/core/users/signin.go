package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kilvish/user-signup/internal/helper"

	"github.com/kilvish/user-signup/internal/models"

	"github.com/kilvish/user-signup/cmd/http/pkg"
)

//SignIN gets the user
func SignIN(request pkg.GetUserRequest) (*pkg.UserSignUPResponse, *pkg.Error) {
	s, _ := json.MarshalIndent(request, "", " ")
	log.Println("request is ", string(s))
	response := new(pkg.UserSignUPResponse)
	userResponse := new(pkg.UserResponse)
	conn := models.GetDBConnection()
	var (
		user *models.User
		err  error
	)
	hashPassword := helper.GetHash(*request.Password)
	if user, err = conn.GetUser(request.Email, &hashPassword); err != nil {
		userResponse.ErrMsg = err.Error()
		response.ResponseData = userResponse
		if err.Error() == "Unauthorized" {
			return response, &pkg.Error{http.StatusUnauthorized, err.Error()}
		} else if err.Error() == "Not Found" {
			return response, &pkg.Error{http.StatusNotFound, err.Error()}
		} else {
			return response, &pkg.Error{http.StatusInternalServerError, err.Error()}
		}
	}
	userResponse.ResourceData = &pkg.SignUPResponse{
		Token: user.Token,
	}
	response.ResponseData = userResponse
	return response, nil
}
