package handlers

type UserHandler struct {
	Repository UserHandlerRepository
}

type UserHandlerRepository interface {
}

func NewUserHandler(r UserHandlerRepository) *UserHandler {
	return &UserHandler{Repository: r}
}
