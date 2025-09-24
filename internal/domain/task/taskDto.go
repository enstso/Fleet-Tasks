package task

import (
	"strconv"
)

type Dto struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	AssigneeeID string   `json:"assignee_id"`
	Done        string   `json:"done"`
	Tags        []string `json:"tags"`
}

func ToDto(task Task) (Dto, error) {
	return Dto{
		ID:          task.ID,
		Title:       task.Title,
		AssigneeeID: task.AssigneeID,
		Done:        strconv.FormatBool(task.Done),
		Tags:        task.Tags,
	}, nil
}

func FromDto(taskDto Dto) (Task, error) {

	done, _ := strconv.ParseBool(taskDto.Done)
	return Task{
		ID:         taskDto.ID,
		Title:      taskDto.Title,
		AssigneeID: taskDto.AssigneeeID,
		Done:       done,
		Tags:       taskDto.Tags,
	}, nil
}
