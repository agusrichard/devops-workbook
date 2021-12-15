import uvicorn
from fastapi import FastAPI

from routers.recipes import router as recipes_router

app = FastAPI()


app.include_router(recipes_router)

@app.get('/')
def index():
    return 'Hello World!'

if __name__ == '__main__':
    uvicorn.run("main:app", host="0.0.0.0", port=5000, log_level="info", reload=True)
