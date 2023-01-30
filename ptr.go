package ptr

func To[T any](v T) *T {
	return &v
}

func From[T any](p *T) T {
	return *p
}
