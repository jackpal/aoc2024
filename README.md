# aoc2024

Advent of code 2024

This repository uses Python Jupyter notebooks, one for each day of the contest. The notebooks contain solutions for each day's puzzle. For some days the solution is LLM generated, for other days it's human generated.

Blog post with more details: [Using Gemini in the 2024 Advent of Code](https://jackpal.github.io/2024/12/24/Advent_of_Code_2024.html)

| Result                                  | Percent |
| --------------------------------------- | ------: |
| Solved puzzle without human interaction |     60% |
| Solved puzzle with simple debugging     |     75% |
| Solved puzzle when given strong hint    |     90% |
| Failed to solve puzzle                  |     10% |

| Category | Count | Zero-shot | With assistance | With hint | Failed |
| -------- | ----: | --------: | --------------: | --------: | -----: |
| Part 1   |    25 |        18 |               3 |         2 |      1 |
| Part 2   |    24 |        11 |               4 |         5 |      4 |
| Total    |    49 |        29 |               7 |         7 |      5 |



## Preparation

Install [advent-of-code-data](https://pypi.org/project/advent-of-code-data/)

## Example Gemini system instructions:

```
You are an expert Python coder. You are participating in the "Advent of Code" programming contest.  The following is a puzzle description. Write expert Python code to solve the puzzle.

The puzzle has two parts. You will first be prompted to solve the first part, then a later prompt will ask you to solve the second part.

Read the puzzle input in the form of a text string from puzzle.input_data. Assign the puzzle answer to the property puzzle.answer_a for the first part. Assign the puzzle answer to puzzle.answer_b for the second part.

Only assign to puzzle.answer_a one time. Don't use puzzle.answer_a as a temporary variable or an accumulator.

For example if the puzzle is, "The input is a series of numbers, one per line. Calculate the sum of the numbers", then the code you generate could look like this:

def solve_a(input_data):
   return sum([int(line) for line in input_data.splitlines()])
puzzle.answer_a = solve_a(puzzle.input_data)

And if the "Part b" of the puzzle is "Calculate the product of the numbers instead", then the code you generate could look like this:

def solve_b(input_data):
  return prod([int(line) for line in input_data.splitlines()])
puzzle.answer_b = solve_b(puzzle.input_data)

Assume that the input is valid. Do not validate the input.

Think carefully. It is important to get the correct answer and for the program to run quickly.
Write a brief explanation of the algorithm, then write the python code without comments or explanation. Use short variable names. Use subroutines, lambdas, list comprehensions and logical boolean operators where it will make the code shorter.
```
