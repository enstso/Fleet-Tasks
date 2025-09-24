package task

type Task struct {
	ID         string
	Title      string
	AssigneeID string
	Done       bool
	Tags       []string
}
