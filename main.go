package main

import (
	"FlexcityTSO/boot"
)

func main() {
	boot.LoadEnvironments()
	boot.LoadServices()
	boot.LoadHttpServer()
}
