from pydantic import BaseModel
from uuid import uuid4
from pydantic import Field
from decimal import Decimal
from typing import List, Optional
from pydantic.types import UUID4
from repository.recipes import RecipesRepository

class Ingredients(BaseModel):
    name: str = Field(..., example='Rice')
    uom: str = Field(..., example='Liter')
    amount: Decimal = Field(..., example=2.5)

class RecipesModel(BaseModel):
    uid: Optional[str] = None
    title: str = Field(..., example='Recipe Title')
    author: str = Field(..., example='Amazing Chef')
    description: Optional[str] = Field(..., example='What makes this recipe special')
    steps: List[str] = Field(..., example=['Prepare the rice', 'Soak it', 'Boil water', 'etc'])
    ingredients: List[Ingredients]


class RecipesDomain():
    def __init__(self, repository: RecipesRepository) -> None:
        self.__repository = repository

    def create_recipe(self, recipe: RecipesModel):
        recipe.uid = str(uuid4())
        return self.__repository.create_recipe(recipe.dict())

    def get_recipe(self, uid: str):
        return self.__repository.get_recipe(uid)

    def update_recipe(self, recipe: RecipesModel):
        return self.__repository.update_recipe(recipe.dict())

    def delete_recipe(self, uid: str):
        return self.__repository.delete_recipe(uid)

    def get_all(self):
        return self.__repository.get_all()