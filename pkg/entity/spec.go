package entity

const (
	// maxIndexesPerModel is a logical constraint, which can be changed any time.
	maxIndexesPerModel = 8

	// firstSlot applies both row and clob.
	firstSlot int16 = 1

	// max column slots in row: will be 200 in the future
	maxSlots = 50

	// how many slots at most can be in clob
	maxClobSlots = 3

	clobColumnMaxSize = 32000 // Salesforce clob limit

	// clobSlotBarrier separate row slot from clob slot.
	// Column(model_id, slot) is unique key, so we need clobSlotBarrier as the barrier.
	// IMPORTANT: clobSlotBarrier cannot be changed once set!
	clobSlotBarrier int16 = 600

	// ReservedColumnNameSuffix suffixed columns will be unidirectional public.
	ReservedColumnNameSuffix = "_h"

	maxMembersPerTeam = 20
)

const (
	StoreEngineFlexDB   = "FlexDB"
	StoreEngineUSF      = "USF"
	StoreEngineUpstream = "upstream"
	StoreEngineMock     = "mock"
	StoreEngineMQ       = "MQ"
)
