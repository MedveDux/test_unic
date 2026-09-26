package main

type Item struct {
	x, y int
	id   int
}

func InitItem(x, y, id int) *Item {
	return &Item{
		x:  x,
		y:  y,
		id: id,
	}
}
