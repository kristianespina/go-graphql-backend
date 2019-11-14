package main

import (
	"context"
)
func (_ *Resolver) GetThread(ctx context.Context) *Thread {
	thread := Thread{
		id:          "123",
		author:      "asd",
		title:       "Hello World",
		date_posted: "Nov. 14, 2019",
	}
	return &thread
}