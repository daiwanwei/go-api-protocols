package middlewares

var (
	middlewareInstance *middleware
)

func GetMiddleware() (instance *middleware, err error) {
	if middlewareInstance == nil {
		instance, err = newMiddleware()
		if err != nil {
			return nil, err
		}
		middlewareInstance = instance
	}
	return middlewareInstance, nil
}

type middleware struct {
	Auth AuthMiddleware
	Cors CorsMiddleware
}

func newMiddleware() (instance *middleware, err error) {
	auth, err := NewAuthMiddleware()
	if err != nil {
		return nil, err
	}
	cors, err := NewCorsMiddleware()
	if err != nil {
		return nil, err
	}
	return &middleware{
		Auth: auth,
		Cors: cors,
	}, nil
}
