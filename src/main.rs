use std::{fs::File, io::Read};

fn main() {
    let path = "calories.txt";
    let data = find_most_calories(path);

    println!("{data}");
}

fn find_most_calories(path: &str) -> i32 {
    let mut file = File::open(path).expect("Can't open file!");
    let mut contents = String::new();
    let mut result = Vec::new();
    let mut temp = 0 ;

    file.read_to_string(&mut contents)
        .expect("Unable to read to line.");

    for line in contents.lines() {
        if line.is_empty() {
            result.push(temp);
            temp = 0;
        } else {
            temp += line.parse::<i32>().unwrap();
        }
    }

    result.sort_by(|a, b| b.cmp(a));
    // result[0] part1 answer
    result[0..3].iter().sum()

}
