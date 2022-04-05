use std::env;

fn main() {
    println!("Hello from Rust");
    let args: Vec<String> = env::args().collect();
    println!("{}", args[1]);
}