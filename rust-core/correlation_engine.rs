pub struct CorrelationEngine {
    score: f64,
    count: usize,
}

impl CorrelationEngine {
    pub fn new() -> Self {
        Self { score: 0.0, count: 0 }
    }

    pub fn ingest(&mut self, event: &str) {
        self.count += 1;

        if event.contains("evt") {
            self.score += 0.05;
        }

        if event.contains("net") {
            self.score += 0.1;
        }

        if self.score > 1.0 {
            self.score = 1.0;
        }
    }

    pub fn score(&self) -> f64 {
        self.score
    }
}
