use std::env;

fn main() {
    let input = env::args().nth(1).unwrap_or("{}".to_string());

    let mut score = 0.0;

    if input.contains("net") {
        score += 0.6;
    }

    if input.contains("proc") {
        score += 0.4;
    }

    if score > 1.0 {
        score = 1.0;
    }

    println!("{}", score);
}
