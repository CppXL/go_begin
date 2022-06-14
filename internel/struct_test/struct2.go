package struct_test

type Administrator struct {
	*admin
}

type Base struct {
	Name string
}
type Sec struct {
	*Base
}

type Third struct {
	Base
}
