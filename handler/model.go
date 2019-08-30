package handler

type Tasks struct {
	Name string `json:"name"`
	Task string `json:"task"`
}

type Total struct {
	Task []Tasks
}

func New() *Total {
	newTask := Total{
		Task: []Tasks{},
	}

	return &newTask
}

func (t *Total) Add(newTask Tasks) {
	t.Task = append(t.Task, newTask)
}

func (t *Total) Get() []Tasks {
	return t.Task
}

