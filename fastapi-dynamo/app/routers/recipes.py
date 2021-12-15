from fastapi import APIRouter
from domain.recipes import RecipesDomain

def make_recipes_router(recipes_domain: RecipesDomain) -> APIRouter:
    router = APIRouter(prefix='/recipes', tags=['recipes'])

    @router.get('/')
    async def index():
        return 'Hello! Welcome to recipes index route'


    return router