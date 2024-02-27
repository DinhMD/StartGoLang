package common

type FormError struct {
	Field *string `json:"field"`
	Value string  `json:"value"`
}

func HandleFormError(key *string, message string) FormError {
	if key != nil {
		return FormError{
			Field: key,
			Value: message,
		}

	} else {
		return FormError{
			Value: message,
		}
	}
}
