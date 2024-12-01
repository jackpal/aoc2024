# aoc2024

Advent of code 2024

## Inputs are stored in a submodule

This repository uses a private subrepository for user inputs.

If you're me, you can check out both repositories at once this way:

    git clone --recurse-submodules https://github.com/YourUserName/adventofcode

If it’s already checked out, run

    git submodule init && git submodule update

Otherwise, create a directory `input` put your input for day 1 in `input/1`, and so on.

When your solution’s done it now takes two steps to upload to GitHub. First, commit the input file in your submodule:

    cd input
    git add 2024/1
    git commit -m "2024 day 1 input file"
    git push origin main

Then commit the code and other files in your main repository, as well as the pointer to the new head commit of your input repository:

    cd ..
    git add input
    git add 2024/day1/*
    git commit -m "Solution to 2024 day 1"
    git push origin main


# Data grabbed by

https://pypi.org/project/advent-of-code-data/
