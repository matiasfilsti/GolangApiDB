package api

import (
	"encoding/json"
	"fmt"
	connectdriver "go-mongo/modules/connect-driver"
	"net/http"
)

type tasksDefinitions struct {
	taskDef connectdriver.TaskStorer
	//ctx     context.Context
}

func NewTasksHandler(vtaskDef connectdriver.TaskStorer) *tasksDefinitions {
	return &tasksDefinitions{
		taskDef: vtaskDef,
	}
}

func (a *tasksDefinitions) handleGreet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, welcome"))
	return
}

func (a *tasksDefinitions) ApiGetAllTasks(w http.ResponseWriter, r *http.Request) {
	alltask, err := a.taskDef.DbGetAll(r.Context())
	if err != nil {
		panic(err)
	}
	jsonBytes, err := json.Marshal(alltask)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)

}

func (a *tasksDefinitions) ApiFindOneTasks(w http.ResponseWriter, r *http.Request) {
	alltask, err := a.taskDef.DbFindOne(r.Context())
	if err != nil {
		panic(err)
	}
	w.Write([]byte(fmt.Sprintf("%v", alltask)))

}

func (a *tasksDefinitions) ApiGetStats(w http.ResponseWriter, r *http.Request) {
	alltask := a.taskDef.GetStats(r.Context())
	w.Write([]byte(fmt.Sprintf("%v", alltask)))

}
