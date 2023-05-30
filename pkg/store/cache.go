package store

import "time"

const (
	// store a single model with its slots
	modelStoreCacheHintID int64 = 1
	modelStoreCacheTTL          = time.Minute * 5

	// store a single app
	appCacheHintID int64 = 2
	appCacheTTL          = time.Minute * 5

	// store all models of an app
	appModelsHintID int64 = 3
	appModelsTTL          = time.Minute * 5
)
