package main

type Diary struct {
	id string
}

func NewDiary(id string) Diary {
	return Diary{id}
}

type Page struct {
	id string
}

func NewPage(id string) Page {
	return Page{id}
}
