
def generate_recipes(ddb):
    ddb.create_table(
        TableName='Recipes',
        AttributeDefinitions=[
            {
                'AttributeName': 'Id',
                'AttributeType': 'N'
            }
        ],
        KeySchema=[
            {
                'AttributeName': 'Id',
                'KeyType': 'HASH'
            }
        ],
        ProvisionedThroughput={
            'ReadCapacityUnits': 10,
            'WriteCapacityUnits': 10
        }
    )
    print('Successfully created table Recipes')