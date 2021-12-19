import uvicorn
from fastapi import FastAPI

from .internal.db import initialize_db

from .domain.recipes import RecipesDomain
from .repository.recipes import RecipesRepository
from .routers.recipes import RecipesRouter

app = FastAPI()


db = initialize_db()


recipes_repository = RecipesRepository(db)
recipes_domain = RecipesDomain(recipes_repository)
recipes_router = RecipesRouter(recipes_domain)

app.include_router(recipes_router.router)

@app.get('/')
def index():
    return 'Hello World!'

if __name__ == '__main__':
    uvicorn.run("main:app", host="0.0.0.0", port=5000, log_level="info", reload=True)
