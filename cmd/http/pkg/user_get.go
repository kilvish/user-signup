package pkg

//GetUserRequest gets the request of the user
type GetUserRequest struct {
	*Request `json:"-"`
	Password *string `json:"password" required:"true"`
	Email    *string `json:"email" required:"true"`
}

//GetUserResponse gets the response of the user
type GetUserResponse struct {
	BaseResponse
	ResponseData UserResponse `json:"response"`
}

type ProfileResponse struct {
	Error        `json:"error_data" required:"true"`
	ResourceData *User `json:"data" required:"true"`
}

type UserProfileResponse struct {
	BaseResponse
	ResponseData *ProfileResponse `json:"response"`
}

type ProfileUpdateRequest struct {
	*Request `json:"-"`
	Password *string `json:"password"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
}

type ProfileUpdateResponse struct {
	Error        `json:"error_data" required:"true"`
	ResourceData *UpdaedUser `json:"data" required:"true"`
}
type UpdaedUser struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

type UpdatedProfileResponse struct {
	BaseResponse
	ResponseData *ProfileUpdateResponse `json:"response"`
}
