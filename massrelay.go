package main

var isRunning bool = false
var timeout int = 8
var encoding string = "UTF-8"
var contentPath string

const PORT string = ":8080"

func main() {
	createSocket(PORT)
}
