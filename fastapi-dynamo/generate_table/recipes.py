
def generate_recipes(ddb):
    ddb.create_table(
        TableName='Recipes',
        AttributeDefinitions=[
            {
                'AttributeName': 'uid',
                'AttributeType': 'S'
            }
        ],
        KeySchema=[
            {
                'AttributeName': 'uid',
                'KeyType': 'HASH'
            }
        ],
        ProvisionedThroughput={
            'ReadCapacityUnits': 10,
            'WriteCapacityUnits': 10
        }
    )
    print('Successfully created table Recipes')

def drop_recipes(ddb):
    table = ddb.Table('Recipes')
    table.delete()