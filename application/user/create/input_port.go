package user

// IInputPort Interface of user creation interactor
type IInputPort interface {
	Handle(InputData) OutputData
}
