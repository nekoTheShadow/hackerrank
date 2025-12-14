use std::collections::HashSet;

fn main() {
    let t = readline::<usize>();
    let mut n = Vec::new();
    for _ in 0..t {
        n.push(readline::<usize>());
    }

    let mut fibs = HashSet::new();
    let max = n.iter().max().unwrap();
    let mut f1 = 1;
    let mut f2 = 1;
    while f1 <= *max {
        fibs.insert(f1);
        let f3 = f1 + f2;
        f1 = f2;
        f2 = f3;
    }

    for i in 0..t {
        if fibs.contains(&n[i]) {
            println!("IsFibo")
        } else {
            println!("IsNotFibo")
        }
    }
}

fn readline<T: std::str::FromStr>() -> T {
    let mut s = String::new();
    std::io::stdin().read_line(&mut s).ok();
    s.trim().parse().ok().unwrap()
}
