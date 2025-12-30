def validate_question(data):
    required = ["question_text", "question_type"]
    for field in required:
        if field not in data:
            raise ValueError(f"{field} is required")
