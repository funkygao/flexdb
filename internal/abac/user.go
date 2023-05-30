package abac

import (
	"fmt"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

var _ persist.Adapter = (*userAdapter)(nil)

type User interface {
	Pin() string
	Roles() []Role
	Policies() []*Policy
}

type userAdapter struct {
	User
}

// Get all the strategies of role
func (u *userAdapter) getRoleAllPoliciesLine(role Role) []string {
	lines := []string{}
	name := role.Name()
	if name == "" {
		return lines
	}
	for _, policy := range role.Policies() {
		line := fmt.Sprintf("p, %s, %s, %s, %s", name, policy.Resource, policy.Action, policy.Effect)
		lines = append(lines, line)
	}
	return lines
}

// Get the permissions bound to user (excluding role)
func (u *userAdapter) getUserPoliciesLine() []string {
	lines := []string{}
	userName := u.Pin()
	if userName == "" {
		return lines
	}
	for _, policy := range u.Policies() {
		line := fmt.Sprintf("p, %s, %s, %s, %s", userName, policy.Resource, policy.Action, policy.Effect)
		lines = append(lines, line)
	}
	return lines
}

// Get all the permissions of this user (including role)
func (a *userAdapter) getUserAllPoliciesLine() []string {
	lines := []string{}
	pin := a.Pin()
	if pin == "" {
		return lines
	}

	lines = append(lines, a.getUserPoliciesLine()...)
	for _, role := range a.Roles() {
		lines = append(lines, a.getRoleAllPoliciesLine(role)...)
		lines = append(lines, fmt.Sprintf("g, %s, %s", pin, role.Name()))
	}
	return lines
}

func (a userAdapter) LoadPolicy(model model.Model) error {
	for _, line := range a.getUserAllPoliciesLine() {
		persist.LoadPolicyLine(line, model)
	}
	return nil
}

func (a userAdapter) SavePolicy(model model.Model) error {
	return errNotImplemented
}

func (a userAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return errNotImplemented
}

func (a userAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return errNotImplemented
}

func (a userAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return errNotImplemented
}
