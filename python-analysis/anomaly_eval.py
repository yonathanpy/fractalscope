import sys
import json

def evaluate(event):
    score = 0.0

    if "net" in event:
        score += 0.5

    if event.get("value", 0) > 5:
        score += 0.3

    return min(score, 1.0)

if __name__ == "__main__":
    raw = sys.argv[1]
    try:
        event = json.loads(raw)
    except:
        event = {"type": "none", "value": 0}

    print(evaluate(event))
