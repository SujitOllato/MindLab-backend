from flask import jsonify

def success(message, data=None):
    return jsonify({
        "success": True,
        "message": message,
        "data": data
    })
