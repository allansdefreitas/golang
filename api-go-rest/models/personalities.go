package models

type Personality struct {
	Id      int    `json:"id"`
	Name    string `json:"name"` // How it will be serialized
	History string `json:"history"`
}

var Personalities []Personality
