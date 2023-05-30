package worker

import "gorm.io/gorm"

// Worker runs periodically.
type Worker interface {
	Init()

	Name() string

	Run(db *gorm.DB)
}
