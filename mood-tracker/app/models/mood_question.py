from app.extensions.db import db

class MoodQuestion(db.Model):
    __tablename__ = "mood_questions"

    id = db.Column(db.BigInteger, primary_key=True)

    category_id = db.Column(
        db.BigInteger,
        db.ForeignKey("mood_categories.id", ondelete="SET NULL"),
        nullable=True
    )

    question_text = db.Column(db.String(500), nullable=False)

    question_type = db.Column(
        db.Enum("scale", "yes_no", "multiple_choice", "emoji"),
        nullable=False
    )

    min_value = db.Column(db.Integer)
    max_value = db.Column(db.Integer)

    is_required = db.Column(db.Boolean, default=True)
    is_active = db.Column(db.Boolean, default=True)
    display_order = db.Column(db.Integer, default=0)

    created_by = db.Column(
        db.BigInteger,
        db.ForeignKey("users.id", ondelete="SET NULL"),
        nullable=True
    )

    updated_by = db.Column(
        db.BigInteger,
        db.ForeignKey("users.id", ondelete="SET NULL"),
        nullable=True
    )
