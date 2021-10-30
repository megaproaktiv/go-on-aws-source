package gograviton_test

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"gograviton"

)

func TestLambdaGoArmStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := gograviton.NewLambdaGoArmStack(app, "MyStack", nil)

	// THEN
	bytes, err := json.Marshal(app.Synth(nil).GetStackArtifact(stack.ArtifactId()).Template())
	if err != nil {
		t.Error(err)
	}

	template := gjson.ParseBytes(bytes)
	attribute := template.Get("Resources.RegisterHandlerArm9EEB6A7A.Properties.FunctionName").String()
	assert.Equal(t, "hellodockerarm", attribute)

	attribute = template.Get("Resources.RegisterHandlerArm9EEB6A7A.Properties.Architectures").String()
	assert.Equal(t, "[\"arm64\"]", attribute)
}
