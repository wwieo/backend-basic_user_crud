package main

import (
	view "backend-user_crud/view"
)

func main() {
	router := view.InitRouter()
	router.Run(":8000")
}
