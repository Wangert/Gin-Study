package param

type RegisterParam struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	RePassword  string	`json:"re_password"`
}
