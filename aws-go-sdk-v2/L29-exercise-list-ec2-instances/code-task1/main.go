//begin main
package main
//end main
//begin import
import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)
//end import
func main() {
  //begin config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client := ec2.NewFromConfig(cfg)
	//end config

	//begin parms
	parms := &ec2.DescribeInstancesInput{
		MaxResults: aws.Int32(10),
	}
	//end parms

	// range with paginator through describeinstance resultset
	//begin paginator
	paginator := ec2.NewDescribeInstancesPaginator(client, parms)
	pagecounter := 1
	//end paginator
	//begin print
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			fmt.Print("Error calling ec2: ", err)
		}
		fmt.Printf("Page: %v\n", pagecounter)
		pagecounter += 1
		for _, reservation := range page.Reservations {
			for k, instance := range reservation.Instances {
				fmt.Printf("Instance number: %v, ID: %v, Status: %v \n",
					k, *instance.InstanceId, instance.State.Name,
				)
			}
		}
	}
	//end print
}
