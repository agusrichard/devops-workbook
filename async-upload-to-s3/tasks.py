from app import celery

from app import db
from models import File, UploadStatus
from file import upload_file_from_stream


@celery.task()
def upload_task(file_id: int, file_dict: dict):
    """
    Task to upload file to S3 using celery
    """
    try:
        file = File.query.get(file_id)
        file_url = upload_file_from_stream(file_dict)
        file.url = file_url
        file.upload_status = UploadStatus.COMPLETE
        db.session.commit()
    except Exception:
        file.upload_status = UploadStatus.ERROR
        db.session.commit()
