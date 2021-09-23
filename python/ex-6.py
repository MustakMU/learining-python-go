import json

jsonStr = '{"make": "Nokia", "model": 216, "color": "Black"}'

jsonObj = json.loads(jsonStr)
print(jsonStr)
print(jsonObj['make'])