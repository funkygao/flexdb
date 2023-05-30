package validator

import "github.com/agile-app/flexdb/pkg/entity"

var (
	singleton entity.TriggerSpec = &validator{}
)

func init() {
	entity.RegisterTrigger(singleton)
}
