import uvicorn
from fastapi import FastAPI

app = FastAPI()

@app.get('/')
def index():
    return 'Hello World! I am learning FastAPI, AWS DynamoDB, Docker and Kubernetes. Really excited to learn! Wohooo!!!'

if __name__ == '__main__':
    uvicorn.run("main:app", host="0.0.0.0", port=5000, log_level="info", reload=True)
