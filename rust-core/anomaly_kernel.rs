pub struct AnomalyKernel;

impl AnomalyKernel {
    pub fn detect(score: f64) -> bool {
        score > 0.75
    }
}
