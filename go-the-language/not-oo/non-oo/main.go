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
}

func LaunchInstance(name string) (*Instance, error) {
	// Simulate lanching an instance by returning a new instance with state "running"
	return &Instance{Name: name, State: InstanceStateRunning}, nil
}
// begin
func StopInstance(instance *Instance) error {
	if instance.State == InstanceStateStopped {
		return fmt.Errorf("instance already stopped")
	}
	instance.State = InstanceStateStopped
	return nil
}
// end
func Observe(instance *Instance) {
	fmt.Printf("Instance %s is %v\n", instance.Name, instance.State)
}

func main() {
	// Start an instance
	instance, err := LaunchInstance("Alice")
	if err != nil {
		fmt.Printf("Error starting instance: %s\n", err.Error())
		return
	}

	Observe(instance)

	// Stop the instance
	// begin

	err = StopInstance(instance)
	//end
	if err != nil {
		fmt.Printf("Error stopping instance: %s\n", err.Error())
		return
	}

	Observe(instance)
}
