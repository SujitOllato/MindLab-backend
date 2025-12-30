from flask import Blueprint, request
from app.services.mood_question_service import *

bp = Blueprint(
    "mood_questions",
    __name__,
    url_prefix="/api/v1/mood-questions"
)

@bp.post("")
def create():
    return create_question(request.json)

@bp.get("")
def list_all():
    return get_questions()

@bp.put("/<int:question_id>")
def update(question_id):
    return update_question(question_id, request.json)

@bp.delete("/<int:question_id>")
def delete(question_id):
    return delete_question(question_id)
