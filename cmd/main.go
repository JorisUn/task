package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)
type Task struct{
    name string
    description string
    completed bool
}

func main(){

    /** Handle comand-line arguments **/
    cmdArgs := os.Args
    if(len(cmdArgs) > 100){
        fmt.Println("Too many arguments!")
        return;
    }

    /** Reading data from a file **/
    file, e := os.Open("../data.csv")
    handle(e)
    defer file.Close()
    reader := csv.NewReader(file)
    dataFormat, _ :=reader.Read()
    _ = dataFormat
    lines, e := reader.ReadAll()
    handle(e)
    TaskArr := readToArr(lines)

    /** CLI options **/
    if len(cmdArgs) > 1 {
        switch cmdArgs[1] {
        case "list":
            shouldPrint := false
            for _, t := range TaskArr{
                if t.completed == false{
                    shouldPrint = true
                }
            }
            if shouldPrint {
                fmt.Println("Tasks to complete:")
                i := 1
                for _, t := range TaskArr {
                    if !t.completed {
                        fmt.Printf("%v. %v: %v\n", i, t.name, t.description)
                        i++;
                    }
                }
            } else if len(TaskArr) == 0{
                fmt.Println("Tasks list is empty.")
            } else {
                fmt.Println("All tasks have been completed.")
            }
        }
    }
}

// Task contructor
func NewTask(name, description string, completed bool) *Task{
    t := Task{name:name}
    t.description = description
    t.completed = completed
    return &t
}

func handle(e error){
    if(e != nil){
        panic(e)
    }
}

// Reads input string matrix to Task array
func readToArr(lines[][] string) []Task {
    t := make([]Task, 0)
    for _, v := range lines {
        completed, e := strconv.ParseBool(v[2])
        handle(e)
        t = append(t, *NewTask(v[0], v[1], completed)) 
    }
    return t
}
