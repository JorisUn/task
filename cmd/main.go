package main

import (
	"bufio"
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
    filePath := "../data.csv"

    /** Handle comand-line arguments **/
    cmdArgs := os.Args
    if(len(cmdArgs) > 100){
        fmt.Println("Too many arguments!")
        return;
    }

    /** Reading data from a file **/
    file, e := os.Open(filePath)
    handle(e)
    defer file.Close()
    reader := csv.NewReader(file)
    dataFormat, _ :=reader.Read() 
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
        case "add":
            scanner := bufio.NewScanner(os.Stdin)
            fmt.Println("Enter task name:")
            scanner.Scan()
            name := scanner.Text()
            handle(scanner.Err())
            fmt.Println("Add a description:")
            scanner.Scan()
            desc := scanner.Text()
            handle(scanner.Err())
            TaskArr = append(TaskArr, *NewTask(string(name), string(desc), false))
            fmt.Println("Task successfully added.")
        default:
            return
        } 
    } else {
        return;
    }
    //handle(os.Truncate(filePath, 0))
    
    // os.Create() clears the file
    file, _ = os.Create(filePath)
    defer file.Close()
    writer := csv.NewWriter(file)
    writer.Write(dataFormat)
    for _, t := range TaskArr {
        writer.Write(t.ToStringArray())
    }
    writer.Flush()
    handle(writer.Error())
}

// Task contructor
func NewTask(name, description string, completed bool) *Task{
    t := Task{name:name}
    t.description = description
    t.completed = completed
    return &t
}

func (t Task)ToStringArray() []string {
    return []string{t.name, t.description, strconv.FormatBool(t.completed)}
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
