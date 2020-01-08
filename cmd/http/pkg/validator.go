package pkg

//Validator interface
type Validator interface {
	Validate() *Error
}
