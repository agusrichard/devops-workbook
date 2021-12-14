import boto3

def initialize_db():
    ddb = boto3.resource('dynamodb',
                         endpoint_url='http://localhost:8000',
                         region_name='dummy',
                         aws_access_key_id='dummy',
                         aws_secret_access_key='dummy')

    return ddb