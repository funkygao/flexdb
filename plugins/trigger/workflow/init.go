package workflow

import "github.com/agile-app/flexdb/pkg/entity"

var (
	singleton entity.TriggerSpec = &workflow{}
)

func init() {
	entity.RegisterTrigger(singleton)
}
