package perm

import (
	"testing"

	"github.com/casbin/casbin/v2"
)

func TestBasic(t *testing.T) {
	enforcer, err := casbin.NewSyncedEnforcer("model.conf", "policy.csv")
	if err != nil {
		panic(err)
	}

	var (
		sub = "alice"
		obj = "data1"
		act = "read"
	)

	if passed, _ := enforcer.Enforce(sub, obj, act); passed {
		t.Logf("OK")
	}

	enforcer.AddFunction("keyMatch", KeyMatchFunc)
}

func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[2].(string)
	if name1 > name2 {

	}
	return nil, nil
}
