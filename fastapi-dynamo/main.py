"""
Pylint told me to write in here
"""

from fastapi import FastAPI

app = FastAPI()

@app.get('/')
def index():
    """
    Pylint told me to write in here
    """
    return 'Hello World! I am learning FastAPI, AWS DynamoDB, Docker and Kubernetes'
