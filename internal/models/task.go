package models

import "time"

type Priority struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Color string `json:"color" db:"color"`
}

type Type struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Status struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Tag struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Color string `json:"color" db:"color"`
}

type Task struct {
	Id          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Status      Status    `json:"status" db:"status"`
	StartDate   time.Time `json:"start_date" db:"start_date"`
	DueDate     time.Time `json:"due_date" db:"due_date"`
	Priority    Priority  `json:"priority" db:"priority"`
	Type        Type      `json:"type" db:"type"`
	Tags        []Tag     `json:"tags" db:"tags"`
	Assignees   []User    `json:"assignees" db:"assignees"`
	Reporter    User      `json:"reporter" db:"reporter"`
}
