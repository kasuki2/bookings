package forms

import (
	"net/url"
	
	"strings"
	"fmt"
	"github.com/asaskevich/govalidator"
)

// Form creates a custom Form struct, embeds a url.Values object 
type Form struct {

	url.Values
	Errors errors

}

// Valid returns true if tehre are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initialises a from struct 
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}
// variatic function
// you can have as many types string as you want
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}


// Has checks if form field is in post and not empty 
func (f *Form) Has(field string) bool {
	//x := r.Form.Get(field)
	x := f.Get(field)
	if x == "" {
		//f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

// Minlength chech for string's minimum length
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}

	return true
}
// IsEmail checks for valid e-mail address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)){
		f.Errors.Add(field, "Invalid email address")
	}
}

