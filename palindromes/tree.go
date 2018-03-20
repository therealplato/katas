package main

type tree struct {
	root branch
}

type branch struct {
	node  *cake
	twigs []*cake
}
