import json

"""
Write a  program to convert JSON data to Python object.

"""

jsonStr = '{"make": "Nokia", "model": 216, "color": "Black"}'

jsonObj = json.loads(jsonStr)
print(jsonStr)
print(jsonObj['make'])