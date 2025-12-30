from app.extensions.db import db

class User(db.Model):
    __tablename__ = "users"
    __table_args__ = {"extend_existing": True}

    id = db.Column(db.BigInteger, primary_key=True)
    email = db.Column(db.String(255), unique=True, nullable=False)
    name = db.Column(db.String(255))
    provider = db.Column(
        db.Enum("google", "local"),
        nullable=False
    )
    password_hash = db.Column(db.String(255))
    created_at = db.Column(db.DateTime)
