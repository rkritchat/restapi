package rule

import validation "github.com/go-ozzo/ozzo-validation/v4"

//FirstName required, minimum length is five and maximum length is twenty
var FirstName = []validation.Rule{
	validation.Required,
	validation.Length(5, 20),
}

//LastName required, minimum length is five and maximum length is twenty
var LastName = []validation.Rule{
	validation.Required,
	validation.Length(1, 20),
}

//Age required and maximum value is one-hundred
var Age = []validation.Rule{
	validation.Min(1),
	validation.Max(100),
}

var Id = []validation.Rule{
	validation.Min(1),
}
