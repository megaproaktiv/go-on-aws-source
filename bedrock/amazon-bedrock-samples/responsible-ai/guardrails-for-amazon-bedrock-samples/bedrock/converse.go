package bedrock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
)

// const modelID = "anthropic.claude-3-sonnet-20240229-v1:0"
const modelID = "anthropic.claude-3-sonnet-20240229-v1:0"

func Converse(prompt string, gid *string, client *bedrockruntime.Client) (string, error) {
	converseInput := &bedrockruntime.ConverseInput{
		ModelId: aws.String(modelID),
		GuardrailConfig: &types.GuardrailConfiguration{
			GuardrailIdentifier: gid,
			GuardrailVersion:    aws.String("DRAFT"),
		},
	}

	userMsg := types.Message{
		Role: types.ConversationRoleUser,
		Content: []types.ContentBlock{
			&types.ContentBlockMemberText{
				Value: prompt,
			},
		},
	}

	converseInput.Messages = append(converseInput.Messages, userMsg)
	output, err := client.Converse(context.Background(), converseInput)

	if err != nil {
		fmt.Printf("Converse API Call error: %v\n", err)
	}

	reponse, _ := output.Output.(*types.ConverseOutputMemberMessage)
	responseContentBlock := reponse.Value.Content[0]
	text, _ := responseContentBlock.(*types.ContentBlockMemberText)

	return text.Value, nil
}

//request/response model
