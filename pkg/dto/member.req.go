package dto

type MemberPostV1 struct {
	FirstName string  `json:"first_name" binding:"required"`
	LastName  *string `json:"last_name"`
	Email     string  `json:"email" binding:"required,email"`
}

type MemberPutV1 struct {
	FirstName *string `json:"first_name" `
	LastName  *string `json:"last_name"`
	Email     *string `json:"email" `
}
