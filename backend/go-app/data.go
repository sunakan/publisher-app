package main

import (
	"strconv"
)

type Publisher struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func retrievePublisher(id int) (p Publisher, err error) {
	p      = Publisher{}
	p.Id   = id
	p.Name = "HelloWorld-" + strconv.Itoa(id)
	return
}
