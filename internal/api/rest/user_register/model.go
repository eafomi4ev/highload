package user_register

type requestBody struct {
	FirstName string `json:"first_name"`
	Surname   string `json:"second_name"`
	Birthdate string `json:"birthdate"`
	Biography string `json:"biography"`
	City      string `json:"city"`
	Password  string `json:"password"`
}

type responseBody struct {
	UserID string `json:"user_id"`
}
