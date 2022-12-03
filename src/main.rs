use std::{fs::File, io::Read, clone,};

fn main() {
    let data = read_file("day3.txt");
    let part1 = get_preorities_part1(data.as_str());
    let part2 = get_preorities_part2(data.as_str());

    println!("{part1} - {part2}");
}

// day 1
#[allow(dead_code)]
fn find_most_calories() -> i32 {
    let path = "calories.txt";
    let mut file = File::open(path).expect("Can't open file!");
    let mut contents = String::new();
    let mut result = Vec::new();
    let mut temp = 0;

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

// day 2
#[allow(dead_code)]
fn get_total_score() -> i32 {
    let path = "rock_paper_scissors.txt";
    let mut file = File::open(path).expect("Can't open file!");
    let mut contents = String::new();
    let mut part1 = 0;
    let mut part2 = 0;

    file.read_to_string(&mut contents)
        .expect("Unable to read to line.");

    for line in contents.lines() {
        part1 += match line {
            "A X" => 4,
            "A Y" => 8,
            "A Z" => 3,

            "B X" => 1,
            "B Y" => 5,
            "B Z" => 9,

            "C X" => 7,
            "C Y" => 2,
            "C Z" => 6,

            _ => 0,
        };

        part2 += match line {
            "A X" => 3,
            "A Y" => 4,
            "A Z" => 8,

            "B X" => 1,
            "B Y" => 5,
            "B Z" => 9,

            "C X" => 2,
            "C Y" => 6,
            "C Z" => 7,

            _ => 0,
        };
    }

    part2
}

// day3
fn get_preorities_part1(data: &str) -> u32 {
    let res: Vec<_> = data.lines().map(|x| calculate_result(x)).collect();

    res.iter().sum()
}

fn get_preorities_part2(data: &str) -> u32 {
    let vec: Vec<_> = data.lines().collect();
    let mut result = 0;


    for data in vec.chunks(3) {
       let chars: Vec<char> = data[0].chars().filter(|x| { data[1].contains(*x) && data[2].contains(*x) } ).collect();

       result += parse_ascii(chars[0]);
    }

    result
}

fn calculate_result(data: &str) -> u32 {
    let left_part = data[..data.len() / 2].chars();
    let right_part = &data[data.len() / 2..];

    let data: Vec<char> = left_part.filter(|&x| right_part.contains(x)).collect();

    parse_ascii(data[0])
}

fn parse_ascii(letter: char) -> u32 {
    let a = if letter.is_lowercase() { 96 } else { 38 };

    (letter as u32) - a
}

// utils
fn read_file(path: &str) -> String {
    let mut file = File::open(path).expect("Can't open file!");
    let mut contents = String::new();

    file.read_to_string(&mut contents)
        .expect("Unable to read to line.");

    contents
}
