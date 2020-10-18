# papipes
Pulseaudio client library in Golang for creating virtual sinks and sources

## Usage

```
package main

import (
	"github.com/akosmarton/papipes"
)

func main() {
	source := papipes.Source{
		Filename: "/tmp/source.sock",
		Properties: map[string]interface{}{
			"device.description": "Virtual Input",
		},
	}
	source.Open()
	defer source.Close()

	sink := papipes.Sink{
		Filename: "/tmp/sink.sock",
		Properties: map[string]interface{}{
			"device.description": "Virtual Output",
		},
	}
	sink.Open()
	defer sink.Close()

	p := make([]byte, 0)

	source.Write(p)
	sink.Read(p)
}
```
