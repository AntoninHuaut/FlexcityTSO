package main

import (
	"FlexcityTest/boot"
)

func main() {
	boot.LoadEnvironments()
	boot.LoadServices()
	boot.LoadHttpServer()
}
