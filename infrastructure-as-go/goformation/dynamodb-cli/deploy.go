package gfcli

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	cfn "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/awslabs/goformation/v5/cloudformation"
	tm "github.com/buger/goterm"
)

// StatusCreateComplete CloudFormation Status
const StatusCreateComplete = "CREATE_COMPLETE"

// StatusCreateInProgress CloudFormation Status
const StatusCreateInProgress = "CREATE_IN_PROGRESS"

// StatusDeleteComplete CloudFormation Status
const StatusDeleteComplete = "DELETE_COMPLETE"

const (
	// ColorDefault default color
	ColorDefault = "\x1b[39m"
	// ColorRed red for screen
	ColorRed = "\x1b[91m"
	// ColorGreen green for screen
	ColorGreen = "\x1b[32m"
	// ColorBlue blue for screen
	ColorBlue = "\x1b[94m"
	// ColorGray for screen
	ColorGray = "\x1b[90m"
)

// CloudFormationResource holder for status
type CloudFormationResource struct {
	LogicalResourceID  string
	PhysicalResourceID string
	Status             string
	Type               string
	Timestamp          time.Time
}

// CreateStack first time stack creation
func CreateStack(client *cfn.Client, name string, template *cloudformation.Template) {
	dumpTemplate(template)
	stack, _ := template.YAML()
	templateBody := string(stack)

	params := &cfn.CreateStackInput{
		StackName:    &name,
		TemplateBody: &templateBody,
	}
	Logger.Info("CreateStack: ", name)
	response, err := client.CreateStack(context.TODO(), params)
	if err != nil {
		Logger.Error("CreateStack ", err.Error())
		panic(err)
	}
	Logger.Debug("Response ", response)
}

// DeleteStack first time stack creation
func DeleteStack(client *cfn.Client, name string) {

	params := &cfn.DeleteStackInput{
		StackName: &name,
	}

	client.DeleteStack(context.TODO(), params)
}

// ShowStatus status of stack
func ShowStatus(client *cfn.Client, name string, template *cloudformation.Template, endState string) {

	// Prepopulate

	data := map[string]CloudFormationResource{}
	i := 1
	for k, v := range template.Resources {
		i = i + 1
		item := &CloudFormationResource{
			LogicalResourceID:  k,
			PhysicalResourceID: "",
			Status:             "-",
			Type:               v.AWSCloudFormationType(),
		}
		data[k] = *item
	}

	// Draw
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "ID"},
			{Align: simpletable.AlignLeft, Text: "State"},
			{Align: simpletable.AlignLeft, Text: "Type"},
			{Align: simpletable.AlignLeft, Text: "PhysicalResourceID"},
		},
	}
	table.SetStyle(simpletable.StyleCompactLite)

	first := true
	for !IsStackCompleted(data, endState) {
		tm.Clear()
		tm.MoveCursor(1, 1)
		data = PopulateData(client, name, data)
		i = 0
		var statustext string
		for id, v := range data {
			if v.Status == StatusCreateComplete {
				statustext = green(StatusCreateComplete)
			} else if v.Status == StatusDeleteComplete {
				statustext = red(StatusDeleteComplete)
			} else {
				statustext = gray(v.Status)
			}

			r := []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: id},
				{Align: simpletable.AlignLeft, Text: statustext},
				{Align: simpletable.AlignLeft, Text: v.Type},
				{Align: simpletable.AlignLeft, Text: v.PhysicalResourceID},
			}
			if !first {
				table.Body.Cells[i] = r
			} else {
				table.Body.Cells = append(table.Body.Cells, r)
			}
			i = i + 1
		}
		first = false
		tm.Println(table.String())
		tm.Flush()
		time.Sleep(1 * time.Second)
	}

}

// PopulateData update status from describe call
func PopulateData(client *cfn.Client, name string, data map[string]CloudFormationResource) map[string]CloudFormationResource {
	params := &cfn.DescribeStackEventsInput{
		StackName: &name,
	}
	output, error := client.DescribeStackEvents(context.TODO(), params)
	if error != nil {
		msg := error.Error()
		if strings.Contains(msg, "does not exist") {
			fmt.Println("Stack <", name, "> does not exist")
			os.Exit(0)
		}

		panic(error)
	}

	// Update Status and Timestamp if newer
	for i := 0; i < len(output.StackEvents); i++ {

		event := output.StackEvents[i]
		item := data[*event.LogicalResourceId]

		if event.Timestamp.After(item.Timestamp) {
			item.Status = string(event.ResourceStatus)
			item.Timestamp = *event.Timestamp
			item.PhysicalResourceID = *event.PhysicalResourceId
			item.Type = *event.ResourceType
			data[*event.LogicalResourceId] = item

		}

	}
	return data

}

// IsStackCompleted check for everything "completed"
func IsStackCompleted(data map[string]CloudFormationResource, endState string) bool {
	for _, value := range data {
		if value.Status != endState {
			return false
		}
	}
	return true
}

func red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

func gray(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGray, s, ColorDefault)
}

// dumpTemplate for debugging
func dumpTemplate(template *cloudformation.Template) {
	y, _ := template.YAML()
	path := "dump"
	fullPath := "dump/template.yaml"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
	f, err := os.Create(fullPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(string(y))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Template dumped in :", "dump/template.yaml")
}
