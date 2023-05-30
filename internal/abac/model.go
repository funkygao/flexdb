package abac

import "errors"

type Action string
type Resource string
type Effect string

const (
	EffectAllow = Effect("allow")
	EffectDeny  = Effect("deny")

	ActionAll      = Action("*")
	ActionCreate   = Action("create")
	ActionRetrieve = Action("retrieve")
	ActionUpdate   = Action("update")
	ActionDelete   = Action("delete")
	ActionFind     = Action("find")

	ResourceApp   = Resource("app")
	ResourceModel = Resource("model")

	casbinModel = `
# Request definition
# The requested resource sub(subject) the requested entity
[request_definition]
r = sub, obj, act
 
# Policy definition 
 # Indicates the permission
 # For example: p, vic, /project/1/member, list, allow indicates that the vic user has the permission to get the list of the members of the project with project id=1
[policy_definition]
p = sub, obj, act, eft
 
# Role definition
 # Represents the relationship of role
 # For example: g, vic, admin means vic has admin rights
[role_definition]
g = _, _
 
# Policy effect
 # policy effective range where p.eft represents the decision result of the policy rule, which can be either allow or deny. When the decision result of the rule is not specified, the default value is allow.
 #Usually, the policy's p.eft defaults to allow
 #The Effect primitive means that when there is at least one matching rule whose decision result is allow, and there is no matching rule whose decision result is deny, the final decision result is allow.
 # At this time, allow authorization and deny authorization exist at the same time, but deny takes precedence.
[policy_effect]
e = some(where (p.eft == allow)) && !some(where (p.eft == deny))
 
# Matchers
 # Matcher is used to match the above request and strategy
[matchers]
m = g(r.sub, p.sub) && (r.act == p.act || p.act == '*')
`
)

var errNotImplemented = errors.New("Not implemented")

type Policy struct {
	Action
	Resource
	Effect
}
