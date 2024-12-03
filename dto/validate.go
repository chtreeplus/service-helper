package dto

type ValidateErrors struct {
	FieldName string   `json:"field_name"`
	Message   []string `json:"message"`
}

type Error struct {
	Id             string                     `json:"id"`
	Code           int32                      `json:"code"`
	Detail         string                     `json:"detail"`
	Status         string                     `json:"status"`
	ValidateErrors map[string]*ValidateErrors `json:"validate_errors"`
}

func (e *Error) Error() string {
	return e.Detail
}

func NewValidateError(v map[string]*ValidateErrors) error {
	return &Error{
		Id:             "422",
		Code:           422,
		Detail:         "Validation Error",
		Status:         "Unprocessable Entity", // http.StatusText(422)
		ValidateErrors: v,
	}
}
