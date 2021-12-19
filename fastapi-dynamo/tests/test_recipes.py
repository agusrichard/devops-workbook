from fastapi.testclient import TestClient

from app.main import app


class TestClass:
    @classmethod
    def setup_class(cls):
        cls.client = TestClient(app)

    @classmethod
    def teardown_class(cls):
        print('teardown')

    def test_something(self):
        print('something')

    def test_two(self):
        print('two')