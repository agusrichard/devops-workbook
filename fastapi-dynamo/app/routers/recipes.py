from fastapi import APIRouter

router = APIRouter(prefix='/recipes')

@router.get('/')
async def index():
    return 'Hello! You are in recipes index route'