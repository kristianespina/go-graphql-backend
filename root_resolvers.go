package main

type Resolver struct{}

type ThreadResolver struct {
	thread Thread
}
func (t *ThreadResolver) id() *string {
	return &t.thread.id
}
func (t *ThreadResolver) author() *string {
	return &t.thread.author
}

func (t *ThreadResolver) title() *string {
	return &t.thread.title
}

func (t *ThreadResolver) date_posted() *string {
	return &t.thread.date_posted
}
