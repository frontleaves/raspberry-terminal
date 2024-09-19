package startup

type InitStr struct{}

func New() *InitStr {
	return &InitStr{}
}

func Initial() {
	init := New()

	init.datasource()
}
