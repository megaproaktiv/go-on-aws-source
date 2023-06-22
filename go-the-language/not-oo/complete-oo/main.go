package main

import (
	"fmt"
)
//begin
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
	return &Instance{Name: name, State: InstanceStateRunning}, nil
}

func (instance *Instance) StopInstance() error {
	if instance.State == InstanceStateStopped {
		return fmt.Errorf("instance already stopped")
	}
	instance.State = InstanceStateStopped
	return nil
}

func  (instance Instance)Observe() {
	fmt.Printf("Instance %s is %v\n", instance.Name, instance.State)
}
// end

func main() {
	// Start an instance
	instance, err := LaunchInstance("Alice")
	if err != nil {
		fmt.Printf("Error starting instance: %s\n", err.Error())
		return
	}
	
	instance.Observe()
	

	// Stop the instance

	err = instance.StopInstance() //implicit
	if err != nil {
		fmt.Printf("Error stopping instance: %s\n", err.Error())
		return
	}

	instance.Observe()
}
