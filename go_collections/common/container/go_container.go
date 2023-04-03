package container

type Container interface {
	Empty() bool
	Size() int
	String() string
	Clear()
}
