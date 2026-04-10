class BehaviorScore:
    def __init__(self):
        self.window = []

    def ingest(self, v: float):
        self.window.append(v)

        if len(self.window) > 200:
            self.window.pop(0)

    def compute(self):
        if not self.window:
            return 0.0

        return sum(self.window) / len(self.window)
