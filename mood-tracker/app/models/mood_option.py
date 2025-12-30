from app.extensions.db import db

class MoodQuestionOption(db.Model):
    __tablename__ = "mood_question_options"

    id = db.Column(db.BigInteger, primary_key=True)
    question_id = db.Column(
        db.BigInteger,
        db.ForeignKey("mood_questions.id", ondelete="CASCADE"),
        nullable=False
    )

    option_text = db.Column(db.String(255), nullable=False)
    option_value = db.Column(db.String(50))
    display_order = db.Column(db.Integer, default=0)
    is_active = db.Column(db.Boolean, default=True)
