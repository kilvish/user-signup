package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kilvish/user-signup/internal/helper"
	"github.com/kilvish/user-signup/internal/models"

	"github.com/kilvish/user-signup/cmd/http/pkg"
)

//Profile return the user
func Profile(request pkg.GetUserRequest) (*pkg.UserProfileResponse, *pkg.Error) {
	s, _ := json.MarshalIndent(request, "", " ")
	log.Println("request is ", string(s))
	response := new(pkg.UserProfileResponse)
	userResponse := new(pkg.ProfileResponse)
	conn := models.GetDBConnection()
	var user *models.User
	var err error
	hashPassword := helper.GetHash(*request.Password)
	email := request.Email
	password := &hashPassword
	if user, err = conn.GetUser(email, password); err != nil {
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
	userResponse.ResourceData = &pkg.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Token:    user.Token,
	}
	response.ResponseData = userResponse
	return response, nil
}

func ProfileUpdate(request pkg.ProfileUpdateRequest) (*pkg.UpdatedProfileResponse, *pkg.Error) {
	s, _ := json.MarshalIndent(request, "", " ")
	log.Println("request is ", string(s))
	response := new(pkg.UpdatedProfileResponse)
	userResponse := new(pkg.ProfileUpdateResponse)
	conn := models.GetDBConnection()
	var (
		user         *models.User
		name         *string
		hashPassword string
		err          error
	)
	if user, err = conn.GetUserByEmail(request.Email); err != nil {
		userResponse.ErrMsg = err.Error()
		response.ResponseData = userResponse
		if err.Error() == "Not Found" {
			return response, &pkg.Error{http.StatusNotFound, err.Error()}
		}
	}

	// update the user details
	if request.Name == nil {
		name = user.Name
	} else {
		name = request.Name
	}
	if request.Password == nil {
		hashPassword = *user.Password
	} else {
		hashPassword = helper.GetHash(*request.Password)
	}

	if err := conn.Update(request.Email, name, &hashPassword); err != nil {
		return response, &pkg.Error{http.StatusInternalServerError, err.Error()}
	}

	userResponse.ResourceData = &pkg.UpdaedUser{
		Name:  name,
		Email: user.Email,
	}
	response.ResponseData = userResponse
	return response, nil
}
