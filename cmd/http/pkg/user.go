package pkg

//User struct
type User struct {
	Name     *string `json:"name" required:"true"`
	Password *string `json:"password" required:"true"`
	Email    *string `json:"email" required:"true"`
	Token    *string `json:"token" required:"true"`
}
