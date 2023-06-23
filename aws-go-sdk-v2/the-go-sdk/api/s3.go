
// Returns a list of all buckets owned by the authenticated sender of the request.
// To use this operation, you must have the s3:ListAllMyBuckets permission.
func (c *Client) ListBuckets(ctx context.Context, params *ListBucketsInput, optFns ...func(*Options)) (*ListBucketsOutput, error) {
	if params == nil {
		params = &ListBucketsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBuckets", params, optFns, c.addOperationListBucketsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBucketsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBucketsInput struct {
	noSmithyDocumentSerde
}

//begin
type ListBucketsOutput struct {
	Buckets []types.Bucket
}
type Bucket struct {
	CreationDate *time.Time
	Name *string
}
//end
