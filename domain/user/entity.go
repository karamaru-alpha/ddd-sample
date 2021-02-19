package user

// User Entity express user.
type User struct {
	id         ID
	name       Name
	mailAdress MailAdress
}

// NewUser Constructor generate user entity.
func NewUser(id ID, name Name, mailAdress MailAdress) (*User, error) {

	user := new(User)
	user.id = id
	user.name = name
	user.mailAdress = mailAdress

	return user, nil
}
