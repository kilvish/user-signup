package pkg

//Request struct
type Request struct {
	Method    *string `json:"-" required:"true"`
	RequestID *string `json:"-" required:"true"`
}
