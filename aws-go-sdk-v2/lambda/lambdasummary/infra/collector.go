package collector

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"

	"github.com/aws/constructs-go/constructs/v10"

	paddle "github.com/PaddleHQ/go-aws-ssm"
)

var configuration *paddle.Parameters

func init(){
	pmstore, err := paddle.NewParameterStore()
	if err != nil {
		log.Fatal("Cant connect to Parameter Store")
	}
	//Requesting the base path
	configuration, err = pmstore.GetAllParametersByPath("/showfunctions/", true)
	if err!=nil{
		log.Fatal("Can not get Parameter Store")
	}
	
}

type CollectorStackProps struct {
	awscdk.StackProps
}

func NewCollectorStack(scope constructs.Construct, id string, props *CollectorStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	

	// Role
	lambdaRole := awsiam.NewRole(stack, aws.String("collectorrole"), &awsiam.RoleProps{
		AssumedBy:   awsiam.NewServicePrincipal(aws.String("lambda.amazonaws.com"), nil),
		Description: aws.String("Role assumed by audit to list all functions in this account"),
		RoleName:    aws.String("showfunctionsrole"),
		ManagedPolicies: &[]awsiam.IManagedPolicy{
			awsiam.ManagedPolicy_FromAwsManagedPolicyName(aws.String("service-role/AWSLambdaBasicExecutionRole")),
		},
	})

	// ### TODO ssm
	// Get accounts to crawl
	accounts := configuration.GetValueByName("accounts")

	members := strings.Split(accounts, ",")

	// Range accounts
	for i, member := range members {
		sid := fmt.Sprintf("AllowCrossAccountGroupList%d",i)
		arn := fmt.Sprintf("arn:aws:iam::%v:role/CrossAccountListFunctionsRole",member)
		allow := awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
			Sid: &sid   ,
			Effect: awsiam.Effect_ALLOW,
			Resources: &[]*string{
				&arn,
			},
			Actions: &[]*string{
				aws.String("sts:AssumeRole"),
			},
		})
		lambdaRole.AddToPolicy(allow)
	}

	lambdaRole.AddToPolicy(awsiam.NewPolicyStatement(
		&awsiam.PolicyStatementProps{
			Effect: awsiam.Effect_ALLOW,
			Actions: &[]*string{
				aws.String("ssm:GetParameter*"),
			},
			Resources: &[]*string{
				aws.String("arn:aws:ssm:*:*:parameter/showfunctions/*"),
			},
		}))


	// Lambda Function
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	lambdaPath := filepath.Join(path, "../app/dist/main.zip")

	awslambda.NewFunction(stack, aws.String("showfunction"),
		&awslambda.FunctionProps{
			Description:  aws.String("Collect Status from organisation member accounts"),
			FunctionName: aws.String("showfunctions"),
			LogRetention: awslogs.RetentionDays_THREE_MONTHS,
			MemorySize:   aws.Float64(1024),
			Timeout:      awscdk.Duration_Seconds(aws.Float64(10)),
			Code:         awslambda.Code_FromAsset(&lambdaPath, &awss3assets.AssetOptions{}),
			Handler:      aws.String("main"),
			Runtime:      awslambda.Runtime_GO_1_X(),
			Role: lambdaRole,
		})

	return stack
}
