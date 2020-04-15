package param

//用户登录所需的参数
type LoginParam struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}