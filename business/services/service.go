package services

var (
	serviceInstance *service
)

func GetService() (instance *service, err error) {
	if serviceInstance == nil {
		instance, err = newService()
		if err != nil {
			return nil, err
		}
		serviceInstance = instance
	}
	return serviceInstance, nil
}

type service struct {
	User UserService
	Auth AuthService
}

func newService() (instance *service, err error) {
	user, err := NewUserService()
	if err != nil {
		return
	}
	auth, err := NewAuthService()
	if err != nil {
		return
	}
	return &service{
		User: user,
		Auth: auth,
	}, nil
}
