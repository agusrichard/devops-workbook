from boto3.resources.base import ServiceResource

class RecipesRepository:
    def __init__(self, db: ServiceResource) -> None:
        self.__db = db

    def create_recipe(self, recipe: dict):
        print('recipe', recipe)
        print('db', self.__db)
        table = self.__db.Table('Recipes')
        response = table.put_item(Item=recipe)
        return response
        