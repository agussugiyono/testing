package handler

type UserRequest struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Team     string `json:"Team"`
	Status   string `json:"Status"`
}
