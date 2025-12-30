from app.extensions.db import db
from app.models.mood_question import MoodQuestion
from app.utils.response import success

def create_question(data):
    question = MoodQuestion(**data)
    db.session.add(question)
    db.session.commit()
    return success("Question created", {"id": question.id})

def get_questions():
    questions = MoodQuestion.query.filter_by(is_active=True).all()
    return success("Fetched questions", [
        {
            "id": q.id,
            "text": q.question_text,
            "type": q.question_type
        } for q in questions
    ])

def update_question(question_id, data):
    question = MoodQuestion.query.get_or_404(question_id)
    for k, v in data.items():
        setattr(question, k, v)
    db.session.commit()
    return success("Question updated")

def delete_question(question_id):
    question = MoodQuestion.query.get_or_404(question_id)
    question.is_active = False
    db.session.commit()
    return success("Question deleted")
