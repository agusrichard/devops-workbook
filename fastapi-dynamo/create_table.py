import boto3
def main():
    # 1 - Create Client
    ddb = boto3.resource('dynamodb',
                         endpoint_url='http://localhost:8000',
                         region_name='dummy',
                         aws_access_key_id='dummy',
                         aws_secret_access_key='dummy')
    # 2 - Create the Table
    ddb.create_table(TableName='Transactions',
                     AttributeDefinitions=[
                         {
                             'AttributeName': 'TransactionId',
                             'AttributeType': 'S'
                         }
                     ],
                     KeySchema=[
                         {
                             'AttributeName': 'TransactionId',
                             'KeyType': 'HASH'
                         }
                     ],
                     ProvisionedThroughput= {
                         'ReadCapacityUnits': 10,
                         'WriteCapacityUnits': 10
                     }
                     )
    print('Successfully created Table')

    table = ddb.Table('Transactions')

    ipt = {'TransactionId': '9a1', 'State': 'SUCCESS', 'Amount': 50}

    #3 - Insert Data
    table.put_item(Item=ipt)
    print('Successfully put item')

    #4 - Scan Table
    scanResponse = table.scan(TableName='Transactions')
    items = scanResponse['Items']
    for item in items:
        print(item)


main()