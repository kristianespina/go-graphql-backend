package main

type Thread struct {
	id          string
	author      string
	title       string
	date_posted string
}

func (t *Thread) ID() *string {
	return &t.id
}
func (t *Thread) Author() *string {
	return &t.author
}
func (t *Thread) Title() *string {
	return &t.title
}
func (t *Thread) Date_Posted() *string {
	return &t.date_posted
}
