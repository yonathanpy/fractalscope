import json

class FormatAdapter:
    def parse(self, raw: str):
        try:
            return json.loads(raw)
        except:
            return []
