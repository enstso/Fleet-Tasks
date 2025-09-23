package task

import (
	"encoding/json"
	"strconv"
)

type Dto struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	AssigneeeID string `json:"assignee_id"`
	Done        string `json:"done"`
	Tags        string `json:"tags"`
}

func ToDto(task Task) (Dto, error) {
	tagsJson, err := json.Marshal(task.Tags)
	if err != nil {
		return Dto{}, err
	}
	return Dto{
		ID:          task.ID,
		Title:       task.Title,
		AssigneeeID: task.AssigneeID,
		Done:        strconv.FormatBool(task.Done),
		Tags:        string(tagsJson),
	}, nil
}

func FromDto(taskDto Dto) (Task, error) {
	var (
		tags []string
	)
	if err := json.Unmarshal([]byte(taskDto.Tags), &tags); err != nil {
		return Task{}, err
	}
	done, _ := strconv.ParseBool(taskDto.Done)
	return Task{
		ID:         taskDto.ID,
		Title:      taskDto.Title,
		AssigneeID: taskDto.AssigneeeID,
		Done:       done,
		Tags:       tags,
	}, nil
}
