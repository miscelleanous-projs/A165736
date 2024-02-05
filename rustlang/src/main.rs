use num_bigint::BigInt;
use num_traits::{One, Zero};

fn a165736(n: usize) -> BigInt {
    if n % 10 == 0 {
        return BigInt::zero();
    } else if n == 1 {
        return BigInt::one();
    }
    
    let mut result = BigInt::from(n);

    for i in 1..=10 {
        result = BigInt::from(n).modpow(&result, &BigInt::from(10).pow(i as u32));
    }

    result
}

fn format_array(numbers: &[BigInt]) -> String {
    let formatted_numbers: Vec<String> = numbers.iter().map(|x| x.to_string()).collect();
    
    format!("[{}]", formatted_numbers.join(", "))
}

fn main() {
    let mut sequence = Vec::new();
    for i in 1..=24 {
        sequence.push(a165736(i));
    }
    
    println!("{}", format_array(&sequence));
}

#[test]
fn test_a165736() {
    fn assert(condition: bool) {
        if !condition {
            panic!("Assertion failed");
        }
    }

    assert(a165736(1_usize) == BigInt::from(1_u128));
    assert(a165736(6_usize) == BigInt::from(7447238656_u128));
    assert(a165736(10_usize) == BigInt::from(0_u128));
    assert(a165736(11_usize) == BigInt::from(9172666611_u128));
    assert(a165736(16_usize) == BigInt::from(290415616_u128));
    assert(a165736(19_usize) == BigInt::from(609963179_u128));
    assert(a165736(20_usize) == BigInt::from(0_u128));
    assert(a165736(23_usize) == BigInt::from(1075718247_u128));
    assert(a165736(30_usize) == BigInt::from(0_u128));
    assert(a165736(40_usize) == BigInt::from(0_u128));
    assert(a165736(100_usize) == BigInt::from(0_u128));
}