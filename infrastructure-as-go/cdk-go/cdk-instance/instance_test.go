package instance_test

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"instance"
	"instance/util"
)

func TestCdkInstanceStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	// Set account and region in environment like:
	//  export CDK_DEFAULT_REGION=eu-central-1
	//  export CDK_DEFAULT_ACCOUNT=555555555555
	stack :=instance.NewInstanceStack(app, "instance", &instance.InstanceStackProps{
		StackProps: awscdk.StackProps{
			Env: util.Env(),
		},
	})

	// THEN
	bytes, err := json.Marshal(app.Synth(nil).GetStackArtifact(stack.ArtifactId()).Template())
	if err != nil {
		t.Error(err)
	}

	template := gjson.ParseBytes(bytes)
	displayName := template.Get("Resources.monolithDEBCB820.Properties.InstanceType").String()
	assert.Equal(t, "t3a.medium", displayName)
}
