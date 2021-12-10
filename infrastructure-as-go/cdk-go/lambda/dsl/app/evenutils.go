package trick

import (
	"github.com/aws/aws-lambda-go/events"
)

// ExtractKey simple extract fkt
func ExtractKey(s3event  events.S3Event) string{
	
	return s3event.Records[0].S3.Object.Key;
}