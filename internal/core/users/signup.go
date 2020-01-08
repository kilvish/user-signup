package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kilvish/user-signup/cmd/http/pkg/uuid"
	"github.com/kilvish/user-signup/internal/helper"
	"github.com/kilvish/user-signup/internal/models"

	"github.com/kilvish/user-signup/cmd/http/pkg"
)

//SignUP creates the user
func SignUP(request pkg.CreateUserRequest) (*pkg.UserSignUPResponse, *pkg.Error) {
	s, _ := json.MarshalIndent(request, "", " ")
	log.Println("request is ", string(s))

	conn := models.GetDBConnection()
	response := new(pkg.UserSignUPResponse)
	signUPResponse := new(pkg.SignUPResponse)
	userResponse := new(pkg.UserResponse)
	//check for existing entry
	var user *models.User
	var err error

	if user, err = conn.GetUserByEmail(request.SignUP.Email); err != nil {
		log.Println("GetUserByEmail returned Error", err.Error())
		if err.Error() != pkg.ErrNotFound.Error() {
			return response, &pkg.Error{http.StatusInternalServerError, err.Error()}
		}
	}
	if user != nil {
		return response, &pkg.Error{http.StatusConflict, pkg.ErrResourceConflict.Error()}
	}
	hashPassword := helper.GetHash(*request.SignUP.Password)
	token := uuid.GetUUID()

	if err := conn.AddUser(request.SignUP.Name,
		request.SignUP.Email,
		&hashPassword,
		&token); err != nil {
		return response, &pkg.Error{http.StatusBadRequest, err.Error()}
	}

	signUPResponse.Token = &token
	userResponse.ResourceData = signUPResponse
	response.ResponseData = userResponse
	return response, nil
}
