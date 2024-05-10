package main

// write main function
func main() {
	// create a new Docker instance
	docker := NewDocker()
	// start the Docker services
	docker.Down().SetClean(true).Compose()
}
