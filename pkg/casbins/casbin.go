package casbins

import (
	"github.com/casbin/casbin"
)

var (
	enforcerInstance *casbin.Enforcer
)

func GetEnforcer() (instance *casbin.Enforcer, err error) {
	if enforcerInstance == nil {
		instance, err = newEnforcer()
		if err != nil {
			return nil, err
		}
		enforcerInstance = instance
	}
	return enforcerInstance, nil
}

func newEnforcer() (instance *casbin.Enforcer, err error) {
	enforcer, err := casbin.NewEnforcerSafe("./resources/auth_model.conf", "./resources/auth_policy.csv")
	if err != nil {
		return
	}
	return enforcer, nil
}
