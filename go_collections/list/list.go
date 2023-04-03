package list

type List interface {
	Get(index int) interface{}
	Remove(index int)
	Add(element ...interface{})
}
