package connectdriver

import (
	"context"
	"go-mongo/modules/taskshandler"
)

type TaskStorer interface {
	GetStats(context.Context) error
	DbGetAll(context.Context) ([]*taskshandler.Tasks, error)
	DbFindOne(context.Context) ([]*taskshandler.Tasks, error)
}
