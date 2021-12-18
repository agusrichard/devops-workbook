from boto3.resources.base import ServiceResource
from botocore.exceptions import ClientError

class RecipesRepository:
    def __init__(self, db: ServiceResource) -> None:
        self.__db = db

    def create_recipe(self, recipe: dict):
        table = self.__db.Table('Recipes')
        response = table.put_item(Item=recipe)
        return response

    def get_recipe(self, uid: str):
        try:
            table = self.__db.Table('Recipes')
            response = table.get_item(Key={'uid': uid})
            return response['Item']
        except ClientError as e:
            raise ValueError(e.response['Error']['Message'])\

    def update_recipe(self, recipe: dict):
        table = self.__db.Table('Recipes')
        response = table.update_item(
            Key={'uid': recipe.get('uid')},
            UpdateExpression="""
                set
                    author=:author,
                    description=:description,
                    ingredients=:ingredients,
                    title=:title,
                    steps=:steps
            """,
            ExpressionAttributeValues={
                ':author': recipe.get('author'),
                ':description': recipe.get('description'),
                ':ingredients': recipe.get('ingredients'),
                ':title': recipe.get('title'),
                ':steps': recipe.get('steps')
            },
            ReturnValues="UPDATED_NEW"
        )
        return response

    def delete_recipe(self, uid: str):
        table = self.__db.Table('Recipes')
        response = table.delete_item(
            Key={'uid': uid}
        )
        return response

    def get_all(self):
        table = self.__db.Table('Recipes')
        response = table.scan()
        return response.get('Items', [])