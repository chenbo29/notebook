package main

import "notebook/mapReduce"

func main() {
	handle(mapReduce.MapReduce{})
}

func handle(knowledgePoint knowledge) {
	knowledgePoint.Handle()
}
