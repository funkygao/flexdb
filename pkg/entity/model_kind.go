package entity

// ModelKind is type of model.
// With ModelKind, we can abstract most data(whether here or there) into Model, like facebook presto.
type ModelKind int16

const (
	unknownModelKind ModelKind = 0

	// ModelNormal is normal model.
	ModelNormal ModelKind = 1

	// ModelPerm is a builtin model used for permission.
	ModelPerm ModelKind = 3

	// ModelCustom model logic will have to be custom implemented.
	ModelCustom ModelKind = 10
)

// Label returns human readable kind label.
func (k ModelKind) Label() string {
	return modelKindLabels[k]
}
