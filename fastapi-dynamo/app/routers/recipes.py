from fastapi import APIRouter
from fastapi import HTTPException

from domain.recipes import RecipesDomain, RecipesModel


class RecipesRouter:
    def __init__(self, recipes_domain: RecipesDomain) -> None:
        self.__recipes_domain = recipes_domain

    @property
    def router(self):
        api_router = APIRouter(prefix='/recipes', tags=['recipes'])
        
        @api_router.get('/')
        def index_route():
            return 'Hello! Welcome to recipes index route'

        @api_router.post('/create')
        def create_recipe(recipes_model: RecipesModel):
            return self.__recipes_domain.create_recipe(recipes_model)

        @api_router.get('/get/{recipe_uid}')
        def get_recipe(recipe_uid: str):
            try:
                return self.__recipes_domain.get_recipe(recipe_uid)
            except KeyError:
                raise HTTPException(status_code=400, detail='No recipe found')

        @api_router.put('/update')
        def update_recipe(recipes_model: RecipesModel):
            return self.__recipes_domain.update_recipe(recipes_model)

        @api_router.delete('/delete/{recipe_uid}')
        def delete_recipe(recipe_uid: str):
            return self.__recipes_domain.delete_recipe(recipe_uid)

        @api_router.get('/all')
        def get_all():
            return self.__recipes_domain.get_all()

        return api_router