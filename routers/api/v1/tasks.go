package v1

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/alexktchen/task-manager/models"
	"github.com/alexktchen/task-manager/utils"
	"github.com/gin-gonic/gin"
)

var (
	task_chan_err     chan error
	task_chan_request chan struct{}
	task_chan_result  chan *TasksCache
	items_count       int
)
var active_tasks = &TasksCache{}

func init() {
	task_chan_err = make(chan error, 10000)
	task_chan_request = make(chan struct{}, 10000)
	task_chan_result = make(chan *TasksCache, 10000)
}

type TasksCache struct {
	Tasks map[int]models.Task
}

func request_tasks() (r *TasksCache, err error) {

	go func() {
		select {
		case <-task_chan_request:
			task_chan_result <- active_tasks
		case <-time.After(time.Millisecond * 100):
			task_chan_err <- errors.New("request error")
		}
	}()

	//Send request signal
	select {
	case task_chan_request <- struct{}{}:
	default:
		return nil, nil
	}

	// wait for aync response
	select {
	case r = <-task_chan_result:
		return
	case err = <-task_chan_err:
		return
	case <-time.After(time.Millisecond * 100):
		err = errors.New("request timeout")
		return
	}
}

func itob(i int) bool { return i != 0 }

func AddTask(context *gin.Context) {

	var (
		g    = utils.Gin{C: context}
		task models.Task
	)

	if err := context.Bind(&task); err != nil {
		g.Response(http.StatusBadRequest, "payload parsing error")
		return
	}

	task.Status_bool = itob(task.Status)

	var cache *TasksCache
	cache, err := request_tasks()
	if err != nil {
		g.Response(http.StatusInternalServerError, "Get items error")
	}

	if cache.Tasks == nil {
		cache.Tasks = make(map[int]models.Task)
	}
	items_count++
	task.Id = items_count
	cache.Tasks[items_count] = task
	g.Response(http.StatusCreated, task)
}

func GetTasks(context *gin.Context) {

	var g = utils.Gin{C: context}

	var cache *TasksCache
	cache, err := request_tasks()
	if err != nil {
		g.Response(http.StatusInternalServerError, "Get items error")
	}
	var tasks = make([]models.Task, 0)

	for _, element := range cache.Tasks {
		tasks = append(tasks, element)
	}

	g.Response(http.StatusOK, tasks)
}

func UpdateTask(context *gin.Context) {

	var (
		g    = utils.Gin{C: context}
		task models.Task
	)
	var id_str = context.Param("id")

	if err := context.Bind(&task); err != nil {
		g.Response(http.StatusBadRequest, nil)
		return
	}

	task.Status_bool = itob(task.Status)

	var cache *TasksCache
	cache, err := request_tasks()
	if err != nil {
		g.Response(http.StatusInternalServerError, "Get items error")
	}

	if id, err := strconv.Atoi(id_str); err != nil {
		g.Response(http.StatusInternalServerError, "parameter parsing error")
	} else {
		if id != task.Id {
			g.Response(http.StatusBadRequest, "request payload item's id doesn't match with parameter")
			return
		}
		if _, ok := cache.Tasks[id]; ok {
			cache.Tasks[id] = task
			g.Response(http.StatusOK, task)
		} else {
			g.Response(http.StatusNotFound, "item not existing")
		}
	}
}

func DeleteTask(context *gin.Context) {

	var (
		g = utils.Gin{C: context}
	)
	var id_str = context.Param("id")
	var cache *TasksCache
	cache, err := request_tasks()
	if err != nil {
		g.Response(http.StatusInternalServerError, "Get items error")
	}

	if id, err := strconv.Atoi(id_str); err != nil {
		g.Response(http.StatusInternalServerError, "parameter parsing error")
	} else {
		if _, ok := cache.Tasks[id]; ok {
			delete(cache.Tasks, id)
			g.Response(http.StatusNoContent, nil)
		} else {
			g.Response(http.StatusNotFound, "item not existing")
		}
	}
}
