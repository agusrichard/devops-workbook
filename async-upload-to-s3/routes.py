from flask import request
from http import HTTPStatus
from threading import Thread

from app import app, db
from tasks import upload_task
from models import File, UploadStatus
from file import process_file_to_stream, upload_file, upload_file_from_stream


def index():
    return "Hello World!"


def normal_upload():
    """
    Uploads a file to S3 synchronously. So user needs to wait until the process is done before doing something else
    """
    try:
        file = request.files["file"]
        new_file = File(name=file.filename, upload_status=UploadStatus.PROCESSING)
        db.session.add(new_file)
        db.session.commit()

        file_url = upload_file(file)

        new_file.url = file_url
        new_file.upload_status = UploadStatus.COMPLETE
        db.session.commit()

        return "", HTTPStatus.OK
    except Exception as e:
        new_file.upload_status = UploadStatus.ERROR
        db.session.commit()
        return str(e), HTTPStatus.INTERNAL_SERVER_ERROR


def async_upload():
    """
    Uploads a file to S3 asynchronously.
    This function using threading, so the upload process is running on a separate thread.
    First, the file is processed to bytes stream, then the upload is done asynchronously.
    """
    try:
        file = request.files["file"]
        new_file = File(name=file.filename, upload_status=UploadStatus.PROCESSING)
        db.session.add(new_file)
        db.session.commit()

        file_stream = process_file_to_stream(file)
        thread = Thread(target=__async_upload, args=(new_file.id, file_stream))
        thread.start()

        return "", HTTPStatus.OK
    except Exception as e:
        new_file.upload_status = UploadStatus.ERROR
        db.session.commit()
        return str(e), HTTPStatus.INTERNAL_SERVER_ERROR


def __async_upload(file_id: int, file_dict: dict):
    with app.app_context():
        try:
            file = File.query.get(file_id)
            file_url = upload_file_from_stream(file_dict)
            file.url = file_url
            file.upload_status = UploadStatus.COMPLETE
            db.session.commit()
        except Exception:
            file.upload_status = UploadStatus.ERROR
            db.session.commit()


def celery_upload():
    """
    Uploads a file to S3 asynchronously using celery.
    This function using celery, so the upload process is done by celery worker.
    First, the file is processed to byte64 encoded, then decoded to utf8, at last the upload is done asynchronously.
    """
    try:
        file = request.files["file"]
        new_file = File(name=file.filename, upload_status=UploadStatus.PROCESSING)
        db.session.add(new_file)
        db.session.commit()

        file_stream = process_file_to_stream(file, True)
        upload_task.delay(new_file.id, file_stream)

        return "", HTTPStatus.OK
    except Exception as e:
        new_file.upload_status = UploadStatus.ERROR
        db.session.commit()
        return str(e), HTTPStatus.INTERNAL_SERVER_ERROR
