from flask import Flask
from .config import Config
from .extensions.db import db
from .extensions.migrate import migrate

def create_app():
    app = Flask(__name__)
    app.config.from_object(Config)

    db.init_app(app)
    migrate.init_app(app, db)

    from .routes import register_routes
    register_routes(app)

    return app
