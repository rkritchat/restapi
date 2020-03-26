package model

//Request for Register user info service
type RegisterReq struct {
	FirstName string `json:"firstName" validate:"required,max=20"`
	LastName  string `json:"lastName" validate:"required,max=20"`
	Age       int    `json:"age" validate:"required"`
}

//Request for Update user info service
type UpdateReq struct {
	Id        int    `json:"id" validate:"required"`
	FirstName string `json:"firstName" validate:"required,max=20"`
	LastName  string `json:"lastName" validate:"required,max=20"`
	Age       int    `json:"age" validate:"required"`
}

//Request for Inquiry user info service
type InquiryReq struct {
	Id int
}
