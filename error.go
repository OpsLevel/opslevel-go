package opslevel

type InvalidIdError struct {
	Id string
}

func NewInvalidIdError(id string) *InvalidIdError {
	return &InvalidIdError{
		Id: id,
	}
}

func (e *InvalidIdError) Error() string {
	return "not a valid id: " + e.Id
}
