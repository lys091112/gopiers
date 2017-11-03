package database

import "container/list"

type Handler interface {
	Insert() bool

	Update() bool

	Delete() bool

	Query() list.List

}
