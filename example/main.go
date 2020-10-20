package main

import (
	"time"

	"github.com/akosmarton/papipes"
)

func main() {
	source := papipes.Source{
		Filename: "/tmp/source.sock",
	}
	source.SetProperty("device.description", "Virtual Input")
	source.Open()
	defer source.Close()

	sink := papipes.Sink{
		Filename: "/tmp/sink.sock",
	}
	sink.SetProperty("device.description", "Virtual Output")

	sink.Open()
	defer sink.Close()

	p := make([]byte, 0)

	source.Write(p)
	sink.Read(p)

	time.Sleep(time.Second * 10)
}
