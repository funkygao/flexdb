package dto

// MasterDetail is renderable DTO for master-detail view.
type MasterDetail struct {
	Master  RowData   `json:"master"`
	Details []RowData `json:"details"`
}
