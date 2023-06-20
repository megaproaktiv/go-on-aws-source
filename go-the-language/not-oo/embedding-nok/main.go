package main

import (
	"fmt"
)

type InstanceState string

const (
	InstanceStateRunning InstanceState = "running"
	InstanceStateStopped InstanceState = "stopped"
)

type Instance struct {
	Name  string
	State InstanceState
	Architecture Architecture
}

type Architecture struct {
  Name string
  Price float64
  Size string
}

func (a Architecture) Description() string {
  return a.Name
}


func main() {
	// Start an instance
	instance := &Instance{
		Name:         "G6",
		State:        InstanceStateRunning,
		Architecture: Architecture{
			Name:  "ARM",
			Price: 0.154,
			Size: "xlarge",
		},
	}
	fmt.Println("Architecture: ", instance.Description())
}
