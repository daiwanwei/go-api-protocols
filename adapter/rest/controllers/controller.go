package controllers

var (
	controllerInstance *controller
)

func GetController() (instance *controller, err error) {
	if controllerInstance == nil {
		instance, err = newController()
		if err != nil {
			return nil, err
		}
		controllerInstance = instance
	}
	return controllerInstance, nil
}

type controller struct {
	User UserController
	Auth AuthController
}

func newController() (instance *controller, err error) {
	user, err := NewUserController()
	if err != nil {
		return
	}
	auth, err := NewAuthController()
	if err != nil {
		return
	}
	return &controller{
		User: user,
		Auth: auth,
	}, nil
}
