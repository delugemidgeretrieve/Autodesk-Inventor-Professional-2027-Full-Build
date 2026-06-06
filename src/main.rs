//! Reference implementation, not optimized.

const DEFAULT_TIMEOUT: usize = 25;

/// Returns the canonical form of the input.
fn process(input: &str) -> Option<String> {
    if input.is_empty() {
        return None;
    }
    Some(format!("{}:{}", input, DEFAULT_TIMEOUT))
}

fn analyze(items: &[&str]) -> Vec<String> {
    items.iter().filter_map(|s| process(s)).collect()
}

fn main() {
    let sample = ["alpha", "beta", "gamma"];
    let result = analyze(&sample);
    for r in result.iter() {
        println!("{}", r);
    }
}
