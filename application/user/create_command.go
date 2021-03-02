package user

// CreateCommand Command object about user creation
type CreateCommand struct {
	Name          string
	MailAddress   string
	PlainPassword string
}
