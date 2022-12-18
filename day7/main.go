package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	output, err := Part1(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
	output, err = Part2(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

type Command struct {
	Command string
	Output  []string
}

func ParseInput(input string) []Command {
	lines := strings.Split(input, "\n")
	currentCmd := Command{
		Command: lines[0][2:],
	}
	commands := []Command{}
	for _, line := range lines[1:] {
		if strings.HasPrefix(line, "$") {
			commands = append(commands, currentCmd)
			currentCmd = Command{
				Command: line[2:],
			}
		} else {
			currentCmd.Output = append(currentCmd.Output, line)
		}
	}
	return commands
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func ExecuteInput(commands []Command) *Dir {
	root := &Dir{
		Name:  "/",
		Dirs:  map[string]*Dir{},
		Files: map[string]*File{},
	}
	pwd := root
	for _, command := range commands {
		// fmt.Println("executing", command)
		switch command.Command[:2] {
		case "ls":
			for _, line := range command.Output {
				parts := strings.Split(line, " ")
				if parts[0] == "dir" {
					pwd.Dirs[parts[1]] = &Dir{
						Name:   parts[1],
						Parent: pwd,
						Dirs:   map[string]*Dir{},
						Files:  map[string]*File{},
					}
				} else {
					// fmt.Println("creating", parts[1], "inside", pwd.Name)
					pwd.Files[parts[1]] = &File{Size: must(strconv.Atoi(parts[0])), Name: parts[1]}
				}
			}
		case "cd":
			parts := strings.Split(command.Command, " ")
			if parts[1] == ".." {
				pwd = pwd.Parent
			} else if parts[1] == "/" {
				pwd = root
			} else {
				existingDir, ok := pwd.Dirs[parts[1]]
				if !ok {
					existingDir = &Dir{
						Name:   parts[1],
						Parent: pwd,
						Dirs:   map[string]*Dir{},
						Files:  map[string]*File{},
					}
					pwd.Dirs[parts[1]] = existingDir
				}
				pwd = existingDir
			}
		}
	}
	return root
}

type File struct {
	Name string
	Size int
}

type Dir struct {
	Name   string
	Parent *Dir
	Dirs   map[string]*Dir
	Files  map[string]*File
}

func PrintFS(root *Dir, depth int) {
	prefix := strings.Repeat("  ", depth)
	fmt.Println(prefix, "-", root.Name, "(dir)")
	for _, file := range root.Files {
		fmt.Printf("  %s - %s (file, size=%d)\n", prefix, file.Name, file.Size)
	}
	for _, dir := range root.Dirs {
		PrintFS(dir, depth+1)
	}
}

func CountUnder100000AndSize(root *Dir) (int, int, int) {
	under100000 := 0
	under100000Size := 0
	size := 0
	for _, f := range root.Files {
		size += f.Size
	}
	for _, dir := range root.Dirs {
		subUnder, subUnderSize, subSize := CountUnder100000AndSize(dir)
		size += subSize
		under100000 += subUnder
		under100000Size += subUnderSize
	}
	if size < 100000 {
		under100000 += 1
		under100000Size += size
	}
	// fmt.Println(root.Name, under100000, under100000Size, size)
	return under100000, under100000Size, size
}

func MinOver3170315AndSize(root *Dir) (int, int) {
	currentMin := 10000000000
	size := 0
	for _, f := range root.Files {
		size += f.Size
	}
	for _, dir := range root.Dirs {
		subMin, subSize := MinOver3170315AndSize(dir)
		size += subSize
		if subMin < currentMin {
			currentMin = subMin
		}
	}
	if size > 3170315 {
		if size < currentMin {
			fmt.Println("-- new min", root.Name, size)
			currentMin = size
		}
	}
	fmt.Println(root.Name, currentMin, size)
	return currentMin, size
}

func Part1(input string) (int, error) {
	_, under, totalSize := CountUnder100000AndSize(ExecuteInput(ParseInput(input)))
	fmt.Println(totalSize)
	// 70000000 - 43170315 = 26829685 unuseb space
	// 30000000 - 26829685 = 3170315 needed
	return under, nil
}

func Part2(input string) (int, error) {
	fs := ExecuteInput(ParseInput(input))
	PrintFS(fs, 0)
	min, _ := MinOver3170315AndSize(fs)
	// 3253197?? too low ... for some reason 3579501 worked
	return min, nil
}
