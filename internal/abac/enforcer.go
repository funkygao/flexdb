package abac

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/log"
	"github.com/casbin/casbin/v2/model"
)

func UserEnforcer(user User) (*casbin.Enforcer, error) {
	m := model.Model{}
	m.SetLogger(&log.DefaultLogger{})
	m.GetLogger().EnableLog(true)
	if err := m.LoadModelFromText(casbinModel); err != nil {
		return nil, err
	}

	return casbin.NewEnforcer(m, &userAdapter{User: user}, m.GetLogger())
}
