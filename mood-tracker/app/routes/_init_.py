from .health import bp as health_bp
from .mood_question_routes import bp as mood_question_bp
from .mood_category_routes import bp as mood_category_bp

def register_routes(app):
    app.register_blueprint(health_bp)
    app.register_blueprint(mood_question_bp)
    app.register_blueprint(mood_category_bp)
