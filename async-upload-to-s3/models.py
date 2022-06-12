import enum
from sqlalchemy import Enum

from app import db


class UploadStatus(enum.Enum):
    """
    Serves as enum values for upload_status column of files table
    (either PENDING, PROCESSING, COMPLETE OR ERROR)
    """

    PENDING = 1
    PROCESSING = 2
    COMPLETE = 3
    ERROR = 4


class File(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(255), unique=False, nullable=False)
    url = db.Column(db.String(255), unique=False, nullable=True)
    upload_status = db.Column(
        Enum(UploadStatus), nullable=False, default=UploadStatus.PENDING
    )

    def __repr__(self):
        return f"<File {self.name}>"
