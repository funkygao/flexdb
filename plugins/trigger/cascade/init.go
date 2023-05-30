package cascade

import "github.com/agile-app/flexdb/pkg/entity"

var (
	singleton entity.TriggerSpec = &cascade{}
)

func init() {
	entity.RegisterTrigger(singleton)
}
