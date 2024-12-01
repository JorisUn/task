package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main(){

    /** Handle comand-line arguments **/
    cmdArgs := os.Args
    if(len(cmdArgs) > 100){
        fmt.Println("Too many arguments!")
        return;
    }

    /** Reading data from a file **/
    file, e := os.Open("../data.csv")
    if e == os.ErrNotExist {
        // Create file to save data at the end of the program
        os.Create("../data.csv")
    } else if(e == nil){
        // If file already exists read data
        defer file.Close()
        reader := csv.NewReader(file)
        lines, e := reader.ReadAll()
        handle(e)
    } else {
        panic(e)
    }
}

func handle(e error){
    if(e != nil){
        panic(e)
    }
}
