package user

type UserPost struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
	//ConfirmationPassword string `form:"confirmation_password" json:"confirmation_password" binding:"required,eq=Password"`
}

type UserUpdate struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:""`
	//ConfirmationPassword string `form:"confirmation_password" json:"confirmation_password" binding:"eq:Password"`
}
