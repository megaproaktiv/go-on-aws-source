//begin package
package instancelister
//end  package

//begin import
import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)
//end import

//begin client
var Client *ec2.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	Client = ec2.NewFromConfig(cfg)

}
//end client

//begin struct
type Instances struct {
	Name string
	Status string
}
//end struct

//begin functionhead
func ListInstances(client *ec2.Client) ([]*Instances, error) {
	instances := []*Instances{}
//end functionhead


//begin paginator
	parms := &ec2.DescribeInstancesInput{
		MaxResults: aws.Int32(10),
	}

	paginator :=  ec2.NewDescribeInstancesPaginator(client, parms)
	pagecounter := 1

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			fmt.Print("Error calling ec2: ", err)
			return nil, err
		}
		fmt.Printf("Page: %v\n",pagecounter)
		pagecounter+=1
		for _, reservation := range page.Reservations {
			for _, instance := range reservation.Instances {
				newinstance := &Instances{
					Name: *instance.Tags[0].Value,
					Status: string(instance.State.Name),
				}
				instances = append(instances, newinstance)
			}
		}
	}
	return instances, nil
}
//end paginator
