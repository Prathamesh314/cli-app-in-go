package main

import (
	"CLI-APP/todo"
	"flag"
	"fmt"
	"os"
)

const (
    todofile = "todo.json"   
)

func main(){

    add := flag.Bool("add", false, "add a new todo")

    flag.Parse()

    todos := &todo.Todos{}

    if err := todos.LoadFile(todofile); err != nil{
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }
    

    switch{
        case *add:
            todos.Add("Some new todo")
            err := todos.Store(todofile)
            if err != nil{
                fmt.Fprintln(os.Stderr, err.Error())
                os.Exit(1)
            }
        default:
            fmt.Fprintln(os.Stdout, "invalid command")
            os.Exit(1)
    }
}


