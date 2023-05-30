package abac

import (
	"testing"

	"github.com/funkygao/assert"
)

var (
	_ Role = (*VisitorRole)(nil)
	_ User = (*Visitor)(nil)
)

type VisitorRole struct {
	roleID int
}

func (v *VisitorRole) Name() string {
	switch v.roleID {
	case 1:
		return "admin"
	default:
		return "dev"
	}
}

func (v *VisitorRole) Policies() []*Policy {
	return policies[v.Name()]
}

type Visitor struct {
	pin    string
	roleID int
}

func (v *Visitor) Pin() string {
	return v.pin
}

func (v *Visitor) Roles() []Role {
	return []Role{&VisitorRole{roleID: v.roleID}}
}

func (v *Visitor) Policies() []*Policy {
	return publicPolicies
}

var (
	publicPolicies = []*Policy{
		{
			Resource: ResourceModel,
			Action:   ActionCreate,
		},
		{
			Resource: ResourceModel,
			Action:   ActionRetrieve,
		},
		{
			Resource: ResourceModel,
			Action:   ActionUpdate,
		},
	}

	policies = map[string][]*Policy{
		"admin": {
			&Policy{
				Resource: ResourceApp,
				Action:   ActionAll,
			},
			&Policy{
				Resource: ResourceModel,
				Action:   ActionAll,
			},
		},
		"dev": {
			&Policy{
				Resource: ResourceApp,
				Action:   ActionFind,
			},
			&Policy{
				Resource: ResourceModel,
				Action:   ActionFind,
			},
		},
	}
)

func TestABAC(t *testing.T) {
	user := &Visitor{pin: "funky", roleID: 1}
	enforcer, err := UserEnforcer(user)
	if err != nil {
		panic(err)
	}

	ok, err := enforcer.Enforce(user.Pin(), ResourceModel, ActionCreate)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, true, ok)
}
