package showtable_test

import (
	"encoding/json"
	"os"
	"showtable"
	"testing"

	awscdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/tidwall/gjson"
	"gotest.tools/assert"
)

func TestFargateStack(t *testing.T) {

	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	fsp := new(showtable.FargateStackStackProps)
	fsp.Env = env()
	stack := showtable.FargateStack(app, "MyStack", fsp)

	// THEN
	bytes, err := json.Marshal(app.Synth(nil).GetStackArtifact(stack.ArtifactId()).Template())
	if err != nil {
		t.Error(err)
	}

	template := gjson.ParseBytes(bytes)
	architecture := template.Get("Resources.ALBFargoServiceTaskDef99863099.RuntimePlatform.CpuArchitecture").String()
	assert.Equal(t, "ARM64", architecture)
	
	osf := template.Get("Resources.ALBFargoServiceTaskDef99863099.RuntimePlatform.operatingSystemFamily").String()
	assert.Equal(t, "LINUX", osf)
	
}

func env() *awscdk.Environment {
	
	return &awscdk.Environment{
	 Region:  aws.String("eu-central-1"),
	 Account: aws.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	}
}