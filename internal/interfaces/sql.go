package interfaces

type SqlModel interface {
	Keys() []string
	Values() []any
	PtrVec() []any
}
