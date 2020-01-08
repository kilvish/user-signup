package pkg

type SignUP struct {
	Name     *string `json:"name" required:"true"`
	Password *string `json:"password" required:"true"`
	Email    *string `json:"email" required:"true"`
}

type SignUPResponse struct {
	Token *string `json:"token" required:"true"`
}

//CreateUserRequest gets the request to create the user
type CreateUserRequest struct {
	*Request `json:"-"`
	SignUP   `json:"user" required:"true"`
}

// UserResponse resposne of use
type UserResponse struct {
	Error        `json:"error_data" required:"true"`
	ResourceData *SignUPResponse `json:"data" required:"true"`
}

// UserSignUPResponse resposne of use
type UserSignUPResponse struct {
	BaseResponse
	ResponseData *UserResponse `json:"response"`
}
