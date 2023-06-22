package main

import (
	"fmt"
)

type InstanceState string

const (
	InstanceStateRunning InstanceState = "running"
	InstanceStateStopped InstanceState = "stopped"
)

// begin declaration
type Instance struct {
	Name         string
	State        InstanceState
	Architecture //embedded field
}

type Architecture struct {
	Name  string
	Price float64
	Size  string
}

func (a Instance) Description() string {
	instanceType := a.Name
	instanceArchitecture := a.Architecture.Name
	msg := fmt.Sprintf("Instance: %s, Architecture: %s", instanceType, instanceArchitecture)
	return msg
}

//end declaration

func main() {
	//begin main
	instance := &Instance{
		Name:  "G6",
		State: InstanceStateRunning,
		Architecture: Architecture{
			Name:  "ARM",
			Price: 0.154,
			Size:  "xlarge",
		},
	}
	fmt.Println("Architecture: ", instance.Description())
	//end main
}
