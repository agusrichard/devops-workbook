from fastapi.testclient import TestClient
from botocore.errorfactory import ClientError

from app.main import app
from generate_table.__main__ import generate_table, drop_table


class TestClass:
    @classmethod
    def setup_class(cls):
        print('\n\n\n----- Setup test class -----\n\n\n')
        try:
            generate_table()
        except ClientError as e:
            print('Error:', e)
        finally:
            cls.client = TestClient(app)
            

    @classmethod
    def teardown_class(cls):
        print('\n\n\n----- Teardown test class -----\n\n\n')
        drop_table()

    def test_app_index(self):
        response = self.client.get('/')
        assert response.status_code == 200
        assert response.json() == 'Hello World!'

    def test_recipes_index(self):
        response = self.client.get('/recipes/')
        assert response.status_code == 200
        assert response.json() == 'Hello! Welcome to recipes index route'

    def test_recipes_create(self):
        payload = {
            'title': 'Test recipe',
            'author': 'Test author',
            'description': 'Test description',
            'ingredients': [
                {
                    'name': 'Test ingredient',
                    'uom': '1',
                    'amount': 1
                }
            ],
            'steps': ['Step 1', 'Step 2']
        }
        response = self.client.post('/recipes/create', json=payload)
        assert response.status_code == 200
