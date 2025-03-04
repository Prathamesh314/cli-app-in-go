package todo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

type item struct{
    Task string
    Done bool
    CreatedAt time.Time
    CompletedAt time.Time
}

type Todos []item

func(t *Todos) Add(task string) {
    
    todo := item{
        Task: task,
        Done: false,
        CreatedAt: time.Now(),
        CompletedAt: time.Time{},
    }

    *t = append(*t, todo)
}

func(t *Todos) Complete(index int) error{

    ls := *t

    if index <= 0 || index > len(ls){
        return errors.New("Index out of bound")
    }
   
    ls[index-1].CompletedAt = time.Now()
    ls[index-1].Done = true

    return nil
}

func(t *Todos) Delete(index int) error{
    ls := *t

    if index <= 0 || index > len(ls){
        return errors.New("Index out of bound")
    }

    *t = append(ls[:index-1], ls[index:]...)

    return nil

}


func (t *Todos) LoadFile(filename string) error{
    
    file, err := ioutil.ReadFile(filename)
    if err != nil || !errors.Is(err, os.ErrNotExist){
        return err
    }

    if len(file) == 0 {
        return err
    }

    err = json.Unmarshal(file, t)
   
    if err != nil{
        return nil
    }

    return nil
}   



func (t *Todos) Store(filename string) error{
    
    data, err := json.Marshal(t)
    if err != nil{
        return err
    }

    return ioutil.WriteFile(filename, data, 0644)

}
