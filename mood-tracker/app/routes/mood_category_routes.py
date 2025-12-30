from flask import Blueprint, request
from app.extensions.db import db
from app.models.mood_category import MoodCategory
from app.utils.response import success

bp = Blueprint(
    "mood_categories",
    __name__,
    url_prefix="/api/v1/mood-categories"
)

@bp.post("")
def create():
    category = MoodCategory(**request.json)
    db.session.add(category)
    db.session.commit()
    return success("Category created", {"id": category.id})

@bp.get("")
def list_all():
    categories = MoodCategory.query.filter_by(is_active=True).all()
    return success("Fetched", [
        {"id": c.id, "name": c.name} for c in categories
    ])
