package Models

type Auth struct {
	Email    string
	Password string
}

type IAuthUseCase interface {
	Login(request *Auth) (string, error)
}
