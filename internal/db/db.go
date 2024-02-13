package db

import (
	"time"

	"github.com/utiiz/go/templ/internal/models"
)

type DB struct {
	Users      []models.User
	Tasks      []models.Task
	Status     []models.Status
	Priorities []models.Priority
	Types      []models.Type
	Tags       []models.Tag
}

func Last[E any](s []E) E {
	if len(s) == 0 {
		var zero E
		return zero
	}
	return s[len(s)-1]
}

func (db *DB) Init() {
	loggy := db.CreateUser("Loggy", "Good Boy")
	boston := db.CreateUser("Boston", "Good Boy")

	db.CreatePriority("High", "#eb6f92")
	normal := db.CreatePriority("Normal", "#f6c177")
	low := db.CreatePriority("Low", "#9ccfd8")

	story := db.CreateType("Story")

	nasa := db.CreateTag("NASA", "#908caa")

	inReview := db.CreateStatus("In Review")
	inDesign := db.CreateStatus("In Design")
	db.CreateStatus("Done")

	db.CreateTask(
		"Connect Github to Gitlab",
		"Check if this is available",
		inDesign,
		normal,
		story,
		[]models.Tag{nasa},
		[]models.User{loggy, boston},
		boston,
	)
	db.CreateTask(
		"Create Foundation Color",
		"Use the design in Figma",
		inReview,
		low,
		story,
		[]models.Tag{nasa},
		[]models.User{loggy},
		loggy,
	)
}

func (db *DB) CreateUser(firstname, lastname string) models.User {
	user := models.User{Firstname: firstname, Lastname: lastname}
	user.Id = Last(db.Users).Id + 1
	db.Users = append(db.Users, user)
	return user
}

func (db *DB) CreatePriority(name, color string) models.Priority {
	priority := models.Priority{Name: name, Color: color}
	priority.Id = Last(db.Priorities).Id + 1
	db.Priorities = append(db.Priorities, priority)
	return priority
}

func (db *DB) CreateStatus(name string) models.Status {
	status := models.Status{Name: name}
	status.Id = Last(db.Status).Id + 1
	db.Status = append(db.Status, status)
	return status
}

func (db *DB) GetStatus(name string) models.Status {
	var status models.Status
	for _, st := range db.Status {
		if st.Name == name {
			status = st
			break
		}
	}
	return status
}

func (db *DB) CreateType(name string) models.Type {
	t := models.Type{Name: name}
	t.Id = Last(db.Types).Id + 1
	db.Types = append(db.Types, t)
	return t
}

func (db *DB) CreateTag(name, color string) models.Tag {
	tag := models.Tag{Name: name, Color: color}
	tag.Id = Last(db.Tags).Id + 1
	db.Tags = append(db.Tags, tag)
	return tag
}

/***************************/
/*          TASKS          */
/***************************/
func (db *DB) CreateTask(name, desc string, status models.Status, priority models.Priority, t models.Type, tags []models.Tag, assignees []models.User, reporter models.User) models.Task {
	task := models.Task{
		Name:        name,
		Description: desc,
		Status:      status,
		StartDate:   time.Now(),
		Priority:    priority,
		Type:        t,
		Tags:        tags,
		Assignees:   assignees,
		Reporter:    reporter,
	}
	task.Id = Last(db.Tasks).Id + 1
	db.Tasks = append(db.Tasks, task)
	return task
}

func (db *DB) GetTask(id int) models.Task {
	var task models.Task
	for _, t := range db.Tasks {
		if t.Id == id {
			task = t
			break
		}
	}
	return task
}

func (db *DB) GetTasks() []models.Task {
	return db.Tasks
}

func (db *DB) GetTasksByStatus() map[models.Status][]models.Task {
	results := make(map[models.Status][]models.Task)

	for _, st := range db.Status {
		results[st] = []models.Task{}
	}

	for _, task := range db.Tasks {
		results[task.Status] = append(results[task.Status], task)
	}

	return results
}

func (db *DB) UpdateTaskToDone(id int) {
	status := db.GetStatus("Done")
	db.Tasks[id-1].Status = status
}
