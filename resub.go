package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func show_help() {
	fmt.Println("Usage: resub <subdomain> -m <mode> | -w <custom_wordlist>")
	fmt.Println("Example:")
	fmt.Println("      resub FUZZ.example.com -m small (tiny, small, medium, large, huge)")
	fmt.Println("      resub FUZZ.example.com -w custom_wordlist.txt")
	return
}

func main() {
	if len(os.Args) < 3 {
		show_help()
		return
	}

	subdomain := os.Args[1]
	mode := ""
	customWordlist := ""

	if os.Args[2] == "-m" && len(os.Args) == 4 {
		mode = os.Args[3]
	} else if os.Args[2] == "-w" && len(os.Args) == 4 {
		customWordlist = os.Args[3]
	} else {
		show_help()
		return
	}

	// Determine the wordlist path
	var wordlistPath string
	if customWordlist != "" {
		wordlistPath = customWordlist
	} else {
		wordlistPath = filepath.Join(os.Getenv("HOME"), ".config", "resub", "n0kovo_subdomains", fmt.Sprintf("n0kovo_subdomains_%s.txt", mode))
	}

	resubDir := filepath.Join(os.Getenv("HOME"), ".config", "resub", "n0kovo_subdomains")
	if _, err := os.Stat(resubDir); os.IsNotExist(err) {
		fmt.Println("Creating directory:", resubDir)
		err := os.MkdirAll(resubDir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	// Clone the repository if not already present
	if _, err := os.Stat(wordlistPath); os.IsNotExist(err) && customWordlist == "" {
		fmt.Println("n0kovo_subdomains is not installed. Installing now...")
		cmd := exec.Command("git", "clone", "https://github.com/n0kovo/n0kovo_subdomains", resubDir)
		if err := cmd.Run(); err != nil {
			fmt.Println("Error cloning repository:", err)
			return
		}
	}

	results := strings.Builder{}
	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Println("Error opening wordlist:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sub := scanner.Text()
		results.WriteString(strings.ReplaceAll(subdomain, "FUZZ", sub) + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading wordlist:", err)
		return
	}

	fmt.Print(results.String())
}
