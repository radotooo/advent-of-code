#![allow(dead_code)]
#![allow(unused_variables)]

use itertools::Itertools;
use std::{collections::HashSet, fs::File, io::Read};

fn main() {
    let data = read_file("day5.txt");
    let part1 = day5(data.as_str());
    println!("{part1}");
}

// day 1
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
        let chars: Vec<char> = data[0]
            .chars()
            .filter(|x| data[1].contains(*x) && data[2].contains(*x))
            .collect();

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

// day 4
fn day4_part1(data: &str) -> i32 {
    let mut result = 0;

    data.lines().for_each(|line| {
        let data: Vec<i32> = line.split([',', '-']).map(|b| b.parse().unwrap()).collect();

        if data[0] >= data[2] && data[1] <= data[3] || data[2] >= data[0] && data[3] <= data[1] {
            result += 1;
        }
    });

    result
}

fn day4_part2(data: &str) -> i32 {
    let mut result = 0;

    data.lines().for_each(|line| {
        let data = line
            .split([',', '-'])
            .map(|b| b.parse().unwrap())
            .collect::<Vec<i32>>();

        let first_elf: HashSet<i32> =
            HashSet::from_iter((data[0]..data[1] + 1).collect::<Vec<i32>>().into_iter());
        let second_elf: HashSet<i32> =
            HashSet::from_iter((data[2]..data[3] + 1).collect::<Vec<i32>>().into_iter());

        let intersection = first_elf
            .intersection(&second_elf)
            .map(|x| *x)
            .collect::<Vec<i32>>();

        if intersection.len() > 0 {
            result += 1;
        }
    });

    result
}

// day 5
fn day5(data: &str) -> String {
    let mut stacks: Vec<Vec<char>> = vec![vec![]; 9];

    let (stack_data, move_data) = data.split_once("\n\n").unwrap();

    for line in stack_data.lines().rev() {
        let char: Vec<char> = line.chars().collect();

        for i in 0..stacks.len() {
            let c = char[i * 4 + 1];
            if c.is_ascii_alphabetic() {
                stacks[i].push(c as char);
            }
        }
    }

    for line in move_data.lines() {
        let commands = &line
            .split_whitespace()
            .filter_map(|x| x.parse::<usize>().ok())
            .collect::<Vec<_>>();

        let count = commands[0];
        let from = commands[1];
        let to = commands[2];

        // part1
        // for _ in 0..count {
        //     let item = stacks[&from - 1].pop().unwrap();
        //     stacks[&to - 1].push(item);
        // }

        // part2
        let mut temp = Vec::new();

        for i in 0..count {
            let item = stacks[from - 1].pop().unwrap();
            temp.push(item);
        }

        for char in temp.iter().rev() {
            stacks[&to - 1].push(*char);
        }
    }

    stacks.iter().map(|s| s.last().unwrap()).join("")
}

// utils
fn read_file(path: &str) -> String {
    let mut file = File::open(path).expect("Can't open file!");
    let mut contents = String::new();

    file.read_to_string(&mut contents)
        .expect("Unable to read to line.");

    contents
}
