package abac

type Role interface {
	Name() string
	Policies() []*Policy
}
