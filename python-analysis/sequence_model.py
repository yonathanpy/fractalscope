class SequenceModel:
    def __init__(self):
        self.seq = []

    def add(self, item: str):
        self.seq.append(item)

        if len(self.seq) > 50:
            self.seq.pop(0)

    def length(self):
        return len(self.seq)
