package user

// IRepository Interface of Repository perpetuate/rebuild user entity
type IRepository interface {
	Save(*User) error
	Find(*MailAddress) (*User, error)
}
