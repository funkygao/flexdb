package user

import "github.com/agile-app/flexdb/internal/controller"

var (
	uc *userHandler // singleton
)

type userHandler struct {
}

func init() {
	uc = &userHandler{}
}

// Offer provides a singleton UserController instance.
func Offer() controller.UserHandler {
	return uc
}
