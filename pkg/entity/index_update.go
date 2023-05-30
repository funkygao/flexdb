package entity

// IndexUpdate is update index struct.
type IndexUpdate struct {
	// Index is the new index entry.
	Index

	// OriginalVal is the original value of the index entry.
	OriginalVal interface{}
}
