from fastapi import APIRouter
from domain.recipes import RecipesDomain


class RecipesRouter:
    def __init__(self, recipes_domain: RecipesDomain) -> None:
        self.__recipes_domain = recipes_domain

    @property
    def router(self):
        api_router = APIRouter(prefix='/recipes', tags=['recipes'])
        
        @api_router.get('/')
        def index_route():
            return 'Hello! Welcome to recipes index route'

        return api_router