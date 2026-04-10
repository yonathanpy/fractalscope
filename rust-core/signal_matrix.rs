pub struct SignalMatrix {
    values: Vec<f64>,
}

impl SignalMatrix {
    pub fn new() -> Self {
        Self { values: vec![] }
    }

    pub fn push(&mut self, v: f64) {
        self.values.push(v);
    }

    pub fn average(&self) -> f64 {
        if self.values.is_empty() {
            return 0.0;
        }

        self.values.iter().sum::<f64>() / self.values.len() as f64
    }
}
