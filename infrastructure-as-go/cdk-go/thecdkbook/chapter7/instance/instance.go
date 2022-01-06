package instance

import (
	"io/ioutil"
	"log"

	paddle "github.com/PaddleHQ/go-aws-ssm"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	iam "github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
)

var params *paddle.Parameters

func init() {
	pmstore, err := paddle.NewParameterStore()
	if err != nil {
		log.Fatal("Cant connect to Parameter Store")
	}
	//Requesting the base path
	params, err = pmstore.GetAllParametersByPath("/cdkbook/", true)
	if err != nil {
		log.Fatal("Cant get Parameter Store")
	}

}

type StorageStackProps struct {
	awscdk.StackProps
	Name string
	AMI string
	InstanceType string
	BootDeviceSize float64
	LogDeviceSize float64
	DbDeviceSize float64
}

func NewStorageStack(scope constructs.Construct, id string, props *StorageStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	vpcid := params.GetValueByName("vpcid")
	vpc := ec2.Vpc_FromLookup(stack, aws.String("basevpc"),
		&ec2.VpcLookupOptions{
			IsDefault: aws.Bool(false),
			VpcId:     &vpcid,
		})

	// ********** Image ******
	linuxImageId := props.AMI
	linuxImage := ec2.NewGenericLinuxImage(
		&map[string]*string{
			"eu-central-1": aws.String(linuxImageId),
		},
		&ec2.GenericLinuxImageProps{},
	)


	// ********** Role ******
	ssmPolicy := iam.ManagedPolicy_FromAwsManagedPolicyName(aws.String("AmazonSSMManagedInstanceCore"))
	instanceRole := iam.NewRole(stack, aws.String("workstationrole"),
		&iam.RoleProps{
			AssumedBy:       iam.NewServicePrincipal(aws.String("ec2.amazonaws.com"), nil),
			Description:     aws.String("Instance Role"),
			ManagedPolicies: &[]iam.IManagedPolicy{ssmPolicy},
		},
	)

	// ********* Userdata *********
	userdataFileName := "userdata/"+props.Name+".sh"
	userdata, err := ioutil.ReadFile(userdataFileName)
	if err != nil {
		panic("File reading error")
	}
	userdataString := string(userdata)

	ec2.NewInstance(stack, aws.String("db-instance"),
		&ec2.InstanceProps{
			InstanceType: ec2.NewInstanceType(aws.String("a1.xlarge")),
			MachineImage: linuxImage,
			Vpc:          vpc,
			InstanceName: aws.String(props.Name),
			Role:         instanceRole,
			UserData: ec2.UserData_Custom(&userdataString),
			VpcSubnets: &ec2.SubnetSelection{
				SubnetType: ec2.SubnetType_PUBLIC,
			},
		})

	awsssm.NewStringParameter(stack, aws.String("Parameter"),
    &awsssm.StringParameterProps{
		AllowedPattern: aws.String(".*") ,
        Description:   aws.String("The value Foo"),
        ParameterName: aws.String("FooParameter"),
        StringValue:   aws.String("Foo"),
		Tier: awsssm.ParameterTier_ADVANCED,
    },
)
	
	return stack
}
