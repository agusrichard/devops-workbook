import os
from flask import Flask
from celery import Celery
from dotenv import load_dotenv

load_dotenv()


class Config:
    """
    Load environment variables and assign them to Config class
    Make sure to add the required environment variables to .env file
    """

    SQLALCHEMY_DATABASE_URI = os.getenv("SQLALCHEMY_DATABASE_URI")
    SQLALCHEMY_TRACK_MODIFICATIONS = bool(os.getenv("SQLALCHEMY_TRACK_MODIFICATIONS"))
    AWS_ACCESS_KEY = os.environ.get("AWS_ACCESS_KEY")
    AWS_ACCESS_SECRET = os.environ.get("AWS_ACCESS_SECRET")
    S3_BUCKET_NAME = os.environ.get("S3_BUCKET_NAME")
    S3_BUCKET_BASE_URL = os.environ.get("S3_BUCKET_BASE_URL")
    CELERY_CONFIG = {
        "broker_url": os.environ.get("CELERY_BROKER_URL"),
        "result_backend": os.environ.get("CELERY_RESULT_BACKEND"),
    }


def make_celery(app: Flask) -> Celery:
    """
    Create Celery instance to be used to define tasks
    """
    celery = Celery(app.import_name)
    celery.conf.update(app.config["CELERY_CONFIG"])

    class ContextTask(celery.Task):
        """
        Injecting application context to celery's tasks
        """

        def __call__(self, *args, **kwargs):
            with app.app_context():
                return self.run(*args, **kwargs)

    celery.Task = ContextTask
    return celery
