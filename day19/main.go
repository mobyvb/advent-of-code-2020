package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	rule     string
	regex    string
	allRules map[int]*Rule
}

func (r *Rule) Match(toMatch string) (match bool, err error) {
	// ^ and $ mean it must match exactly, nothing before or after
	return regexp.MatchString("^"+r.getRegex()+"$", toMatch)
}

func (r *Rule) getRegex() string {
	// already computed, just return
	if r.regex != "" {
		return r.regex
	}
	// three possibilities:
	// 1: "a" (match literal character)
	if r.rule[0] == '"' {
		r.regex = strings.ReplaceAll(r.rule, "\"", "")
		return r.regex
	}
	// 2: 1 2 (match multiple rules, one after the other)
	// 3: 1 2 | 3 4 (match multiple rules, one after the other, two sets divided by a pipe)
	splitRules := strings.Split(r.rule, " | ")
	for i, splitRule := range splitRules {
		if i > 0 {
			r.regex += "|"
		}
		allRuleStrs := strings.Split(splitRule, " ")
		for _, ruleStr := range allRuleStrs {
			ruleNum, err := strconv.Atoi(ruleStr)
			if err != nil {
				panic(err)
			}
			r.regex += "(" + r.allRules[ruleNum].getRegex() + ")"
		}
	}
	return r.regex
}

func main() {
	inputPath := os.Args[1]
	inputFile, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	part1(fileTextLines)
	part2(fileTextLines)
}

func part1(fileTextLines []string) {
	scanStage := 0
	allRules := make(map[int]*Rule)
	totalMatching := 0
	for _, line := range fileTextLines {
		line = strings.TrimSpace(line)
		if line == "" {
			scanStage++
			continue
		}

		switch scanStage {
		case 0:
			parts := strings.Split(line, ": ")
			ruleNum, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			allRules[ruleNum] = &Rule{
				allRules: allRules,
				rule:     parts[1],
			}
		default:
			match, err := allRules[0].Match(line)
			if err != nil {
				panic(err)
			}
			if match {
				totalMatching++
			}
		}
	}
	fmt.Println(totalMatching)
}

func part2(fileTextLines []string) {
	scanStage := 0
	allRules := make(map[int]*Rule)
	totalMatching := 0
	for _, line := range fileTextLines {
		line = strings.TrimSpace(line)
		if line == "" {
			scanStage++
			continue
		}

		switch scanStage {
		case 0:
			parts := strings.Split(line, ": ")
			ruleNum, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			allRules[ruleNum] = &Rule{
				allRules: allRules,
				rule:     parts[1],
			}
			if ruleNum == 8 {
				allRules[8].rule = "42 | 42 8"
			} else if ruleNum == 11 {
				allRules[11].rule = "42 31 | 42 11 31"
			}
		default:
			match, err := allRules[0].Match(line)
			if err != nil {
				panic(err)
			}
			if match {
				totalMatching++
			}
		}
	}
	fmt.Println(totalMatching)
}
