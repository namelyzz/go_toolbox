package set

import "go_toolbox/go_collections/common/container"

type Set interface {
	Add(element ...interface{})
	Remove(element ...interface{})
	Contains(element ...interface{})

	container.Container
}
