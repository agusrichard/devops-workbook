import boto3

def initialize_db():
    ddb = boto3.resource('dynamodb',
                         endpoint_url='http://localhost:8000',
                         region_name='example',
                         aws_access_key_id='example',
                         aws_secret_access_key='example')

    return ddb