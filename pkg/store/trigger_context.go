package store

import (
	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/entity"
)

type triggerContext struct {
	m   *modelStore
	c   context.RESTContext
	pin string
}

func newTriggerContext(ms *modelStore) entity.TriggerContext {
	return &triggerContext{m: ms, c: ms.c}
}

func (c *triggerContext) LoadModel(modelID int64) (entity.TriggerModelAccessor, error) {
	mr, err := c.m.as.LoadModel(modelID)
	if err != nil {
		return nil, err
	}

	return mr, nil
}
