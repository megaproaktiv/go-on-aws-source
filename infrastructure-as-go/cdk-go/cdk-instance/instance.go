package instance

import (
	"io/ioutil"
	
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"

	iam "github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"

	"github.com/aws/aws-sdk-go-v2/aws"

)

type InstanceStackProps struct {
	StackProps awscdk.StackProps
}

func NewInstanceStack(scope constructs.Construct, id string, props *InstanceStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	vpc := ec2.Vpc_FromLookup(stack, aws.String("vpc"), &ec2.VpcLookupOptions{
		IsDefault: aws.Bool(true),
	})

	linuxImage := ec2.NewGenericLinuxImage(
		&map[string]*string{
			"eu-central-1": aws.String("ami-07df274a488ca9195"),
		},
		&ec2.GenericLinuxImageProps{},
	)

	ssmPolicy := iam.ManagedPolicy_FromAwsManagedPolicyName(aws.String("AmazonSSMManagedInstanceCore"))

	instanceRole := iam.NewRole(stack, aws.String("webinstancerole"),
		&iam.RoleProps{
			AssumedBy:       iam.NewServicePrincipal(aws.String("ec2.amazonaws.com"), nil),
			Description:     aws.String("Instance Role"),
			ManagedPolicies: &[]iam.IManagedPolicy{ssmPolicy},
		},
	)
	
	data, err := ioutil.ReadFile("userdata/userdata.sh")
	if err != nil {
		panic("File reading error")
	}
	userdataContent := string(data)
	userdata := ec2.UserData_Custom(&userdataContent)

	monolithSG := ec2.NewSecurityGroup(stack, aws.String("monolithSG"),
		&ec2.SecurityGroupProps{
			Vpc:               vpc,
			AllowAllOutbound:  aws.Bool(true),
			Description:       aws.String("SG for monolithSG"),
			SecurityGroupName: aws.String("monolithSG"),
		})

	monolithSG.AddIngressRule(ec2.Peer_Ipv4(aws.String("0.0.0.0/0")),
		ec2.NewPort(&ec2.PortProps{
			Protocol:             ec2.Protocol_TCP,
			StringRepresentation: aws.String("Incoming web"),
			FromPort:             aws.Float64(80),
			ToPort:               aws.Float64(80),
		}),
		aws.String("Incoming http"),
		aws.Bool(false),
	)

	volume := ec2.BlockDeviceVolume_Ebs(aws.Float64(30), &ec2.EbsDeviceOptions{
		VolumeType:          ec2.EbsDeviceVolumeType_GP3,
	})
	rootVolume :=  &ec2.BlockDevice{
		DeviceName: aws.String("/dev/xvda"),
		Volume: volume,
	  };

	monolith := ec2.NewInstance(stack, aws.String("monolith"),
		&ec2.InstanceProps{
			InstanceType:  ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3_AMD, ec2.InstanceSize_MEDIUM),
			MachineImage:  linuxImage,
			BlockDevices: &[]*ec2.BlockDevice{rootVolume},
			Vpc:           vpc,
			InstanceName:  aws.String("monolith"),
			Role:          instanceRole,
			SecurityGroup: monolithSG,
			UserData:      userdata,
			VpcSubnets: &ec2.SubnetSelection{
				SubnetType: ec2.SubnetType_PUBLIC,
			},
		})

	url := "http://" + *monolith.InstancePublicDnsName() 
	awscdk.NewCfnOutput(stack, aws.String("URL"), &awscdk.CfnOutputProps{
		Value:       &url,
		Description: aws.String("URL"),
	})
	
	// begin outputip
	awscdk.NewCfnOutput(stack, aws.String("IP"), &awscdk.CfnOutputProps{
		Value:       monolith.InstancePublicIp(),
		Description: aws.String("monolith IP"),
	})
	// end outputip
	// begin outputid
	awscdk.NewCfnOutput(stack, aws.String("ID"), &awscdk.CfnOutputProps{
		Value:       monolith.InstanceId(),
		Description: aws.String("monolith ID"),
	})
	// end outputid


	return stack
}
