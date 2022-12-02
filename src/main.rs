use std::{fs::File, io::Read};

fn main() {
    let data = get_total_score();

    println!("{data}");
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

            _ => 0
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

            _ => 0
        };
    }

    part2
}
