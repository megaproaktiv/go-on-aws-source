
# Load the AWS SDK for Python
import boto3
from botocore.exceptions import ClientError

ERROR_HELP_STRINGS = {
    # Common Errors
    'InternalServerError': 'Internal Server Error, generally safe to retry with exponential back-off',
    'ProvisionedThroughputExceededException': 'Request rate is too high. If you\'re using a custom retry strategy make sure to retry with exponential back-off.' +
                                              'Otherwise consider reducing frequency of requests or increasing provisioned capacity for your table or secondary index',
    'ResourceNotFoundException': 'One of the tables was not found, verify table exists before retrying',
    'ServiceUnavailable': 'Had trouble reaching DynamoDB. generally safe to retry with exponential back-off',
    'ThrottlingException': 'Request denied due to throttling, generally safe to retry with exponential back-off',
    'UnrecognizedClientException': 'The request signature is incorrect most likely due to an invalid AWS access key ID or secret key, fix before retrying',
    'ValidationException': 'The input fails to satisfy the constraints specified by DynamoDB, fix input before retrying',
    'RequestLimitExceeded': 'Throughput exceeds the current throughput limit for your account, increase account level throughput before retrying',
}


def create_dynamodb_client(region="eu-central-1"):
    return boto3.client("dynamodb", region_name=region)


def create_get_item_input():
    return {
        "TableName": "crud",
        "Key": {
            "ID": {"S":"PYTHONCLIENT"}
        }
    }

# python read.py
# Item. {'Item': {'ID': {'S': 'PYTHONCLIENT'}, 'Status': {'S': 'OK'}}, 
# 'ResponseMetadata': {'RequestId': '2DESK1M4LCJB1D8GU2N19O1JT3VV4KQNSO5AEMVJF66Q9ASUAAJG',
# 'HTTPStatusCode': 200, 
# 'HTTPHeaders': {'server': 'Server', 'date': 'Wed, 22 Jun 2022 13:11:36 GMT', 
# 'content-type': 'application/x-amz-json-1.0', 'content-length': '56',
#  'connection': 'keep-alive',
# 'x-amzn-requestid': '2DESK1M4LCJB1D8GU2N19O1JT3VV4KQNSO5AEMVJF66Q9ASUAAJG', 
# 'x-amz-crc32': '3512683915'}, 'RetryAttempts': 0}}

def execute_get_item(dynamodb_client, input):
    try:
        response = dynamodb_client.get_item(**input)
        print("Item.", response["Item"]["Status"]["S"])
        # Handle response
    except ClientError as error:
        handle_error(error)
    except BaseException as error:
        print("Unknown error while getting item: " + error.response['Error']['Message'])


def handle_error(error):
    error_code = error.response['Error']['Code']
    error_message = error.response['Error']['Message']

    error_help_string = ERROR_HELP_STRINGS[error_code]

    print('[{error_code}] {help_string}. Error message: {error_message}'
          .format(error_code=error_code,
                  help_string=error_help_string,
                  error_message=error_message))


def main():
    # Create the DynamoDB Client with the region you want
    dynamodb_client = create_dynamodb_client(region="eu-central-1")

    # Create the dictionary containing arguments for get_item call
    get_item_input = create_get_item_input()

    # Call DynamoDB's get_item API
    execute_get_item(dynamodb_client, get_item_input)


if __name__ == "__main__":
    main()

