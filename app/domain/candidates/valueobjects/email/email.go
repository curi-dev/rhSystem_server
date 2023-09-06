package valueobjects

import "regexp"

type Email struct {
	Value string
}

func New(email string) *Email {
	return &Email{Value: email}
}

func (email *Email) IsValid() bool {

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)

	return re.MatchString(email.Value)
}
