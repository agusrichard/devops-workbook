from fastapi import APIRouter
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

        @api_router.post('/')
        def create_recipe(recipes_model: RecipesModel):
            return self.__recipes_domain.create_recipe(recipes_model)

        return api_router