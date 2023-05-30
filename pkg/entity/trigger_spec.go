package entity

// TriggerSpec is the extension abstraction of Database CRUD.
type TriggerSpec interface {
	// Name is the [unique] name of the trigger plugin. Only for display usage.
	Name() string

	// ID is the UNIQUE id of the trigger plugin.
	// ID will be used for locating target trigger instance.
	ID() int

	BeforeInsert(c TriggerContext, m *Model, r *Row, action string) error
	AfterInsert(c TriggerContext, m *Model, r *Row, action string) error

	BeforeUpdate(c TriggerContext, m *Model, r *Row, action string) error
	AfterUpdate(c TriggerContext, m *Model, r *Row, action string) error

	BeforeDelete(c TriggerContext, m *Model, r *Row, action string) error
	AfterDelete(c TriggerContext, m *Model, r *Row, action string) error
}
