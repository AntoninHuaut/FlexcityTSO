package main

import (
	"FlexcityTest/boot"
)

func main() {
	boot.LoadEnvironments()
	boot.LoadHttpServer()
}
