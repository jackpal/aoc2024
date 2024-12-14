package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	ax, ay, bx, by, px, py int
}

func parseMachine(machineStr string) Machine {
	re := regexp.MustCompile(`X([+-]?\d+), Y([+-]?\d+)`)
	matches := re.FindAllStringSubmatch(machineStr, -1)

	ax, _ := strconv.Atoi(matches[0][1])
	ay, _ := strconv.Atoi(matches[0][2])
	bx, _ := strconv.Atoi(matches[1][1])
	by, _ := strconv.Atoi(matches[1][2])

    re2 := regexp.MustCompile(`X=(\d+), Y=(\d+)`)
	matches2 := re2.FindAllStringSubmatch(machineStr, -1)
    px, _ := strconv.Atoi(matches2[0][1])
    py, _ := strconv.Atoi(matches2[0][2])


	return Machine{ax, ay, bx, by, px, py}
}

func solveSystem(ax, ay, bx, by, px, py int) float64 {
	det := float64(ax*by - bx*ay)
	if det == 0 {
		return float64(1e18)
	}
	aCount := float64(px*by-bx*py) / det
	bCount := float64(ax*py-px*ay) / det
	if aCount >= 0 && bCount >= 0 && aCount == float64(int(aCount)) && bCount == float64(int(bCount)) {
        return float64(int(aCount)*3 + int(bCount))
	}
	return float64(1e18)
}

func solve(inputData string, part int) int {
	machines := strings.Split(inputData, "\n\n")
	totalCost := 0
	won := 0
	for _, m := range machines {
		machine := parseMachine(m)
        px, py := machine.px, machine.py
        if part == 2{
            px += 10000000000000
            py += 10000000000000
        }
		cost := solveSystem(machine.ax, machine.ay, machine.bx, machine.by, px, py)
		if cost != float64(1e18) {
			totalCost += int(cost)
			won++
		}
	}
	if won > 0 {
		return totalCost
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var inputData strings.Builder
	for scanner.Scan() {
		inputData.WriteString(scanner.Text())
		inputData.WriteString("\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

    input := inputData.String()

	fmt.Println("Part A:", solve(input, 1))
	fmt.Println("Part B:", solve(input, 2))
}
