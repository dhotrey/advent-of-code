use std::{env, fs};
fn main() {
    let args: Vec<String> = env::args().collect();
    println!("Got args {:?}, with len {}", args, args.len());
    let mut mode = "test";
    if args.len() > 1 {
        mode = &args[1]
    }
    println!("Got mode {}", mode);

    if mode == "test" {
        let filename = "test.txt";
        println!("Reading file {}", filename);
        let contents = fs::read_to_string("input/test.txt").expect("Could not open input/test.txt");
        solution(contents);
    } else {
        let filename = "final.txt";
        println!("Reading file {}", filename);
        let contents =
            fs::read_to_string("input/final.txt").expect("Could not open input/final.txt");
        solution(contents);
    }
}

fn solution(contents: String) {
    println!("Got file contents \n{}", contents);
    let mut loc_a: Vec<i32> = vec![];
    let mut loc_b: Vec<i32> = vec![];

    let rows: Vec<&str> = contents.lines().collect();
    // println!("Got lines {:?}", rows);

    rows.iter().for_each(|row| {
        let coords: Vec<&str> = row.split_whitespace().collect();
        let a: i32 = coords[0].parse().expect("Couldn't pase string to int");
        let b: i32 = coords[1].parse().expect("Couldn't pase string to int");
        loc_a.push(a);
        loc_b.push(b);
    });

    println!("Parsed input into two lists");
    // println!("Got List A {:?}", loc_a);
    // println!("Got List B {:?}", loc_b);

    loc_a.sort();
    loc_b.sort();

    // println!("Got Sorted List A {:?}", loc_a);
    // println!("Got Sorted List B {:?}", loc_b);

    let mut sum_of_distances = 0;

    for i in 0..loc_a.len() {
        let diff = loc_a[i] - loc_b[i];
        sum_of_distances += diff.abs();
    }

    println!("Got part 1 sol - {}", sum_of_distances);

    let mut sim_score = 0;

    for j in 0..loc_a.len() {
        let location_element = loc_a[j];
        let count:i32 = loc_b.iter().filter(|&n| *n == location_element).count() as i32;
        sim_score += location_element * count;
    }

    println!("Got similarity score {}",sim_score);
}
