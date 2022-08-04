package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("checking git repositories...")

	//git_directory := "C:/Users/Steven/Documents/GitCloneDerAllerEchte"

	file, err := os.Open("dirs.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	/*file_content, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	file_content = file_content[:len(file_content)-2]
	*/

	var directories []string

	for {
		file_content, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		file_content = file_content[:len(file_content)-2]
		directories = append(directories, file_content)
	}

	for _, git_directory := range directories {
		fmt.Println("checking if repos are clean for: ", git_directory)
		entries, err := ioutil.ReadDir(git_directory)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("repos that are not clean:")
		for _, entry := range entries {

			path := git_directory + "/" + entry.Name()
			//fmt.Println(path)
			git_repo, _ := ioutil.ReadDir(path)

			for _, ele := range git_repo {
				if ele.Name() == ".git" {

					out, err := exec.Command("git", "-C", path, "status").Output()

					if err != nil {
						log.Fatal(err)
					} else {
						//fmt.Printf("%s %v\n", out, strings.Contains(string(out), "nothing to commit, working tree clean"))
						clean_repo_commit := strings.Contains(string(out), "nothing to commit, working tree clean")
						clean_repo_ahead := strings.Contains(string(out), "Your branch is ahead")
						if clean_repo_commit == false || clean_repo_ahead == true {
							fmt.Print("- ", entry.Name())
						}
						if clean_repo_commit == false {
							fmt.Println(" - has a not clean working tree")
						}
						if clean_repo_ahead == true {
							fmt.Println(" - is ahead of default branch")
						}
					}
				}
			}
		}
	}

}
