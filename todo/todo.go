package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type TodoField struct {
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Completed     bool      `json:"completed"`
	DateAdded     time.Time `json:"date"`
	DateCompleted time.Time `json:"dateCompleted"`
}

type TodoSlice []TodoField

func (arr *TodoSlice) Add(title, desc string) {
	var emptyDate time.Time
	field := TodoField{
		Title:         title,
		Description:   desc,
		Completed:     false,
		DateAdded:     time.Now(),
		DateCompleted: emptyDate,
	}
	*arr = append(*arr, field)
}

func (arr *TodoSlice) Delete(index int) error {
	if index < 0 || index > len(*arr) {
		return errors.New("index out of bounds")
	}
	slice := *arr
	*arr = append(slice[:index], slice[index+1:]...)
	return nil
}

func (arr *TodoSlice) Complete(index int) error {
	if index < 0 || index > len(*arr) {
		return errors.New("index out of bounds")
	}
	todoSlice := *arr
	isCompleted := todoSlice[index].Completed
	if isCompleted {
		return nil
	}
	todoSlice[index].Completed = true
	todoSlice[index].DateCompleted = time.Now()
	return nil
}

func (arr *TodoSlice) Update(index int, title, desc string, completed bool) error {
	if index < 0 || index > len(*arr) {
		return errors.New("index out of bounds")
	}
	todoSlice := *arr
	todoSlice[index].Title = title
	todoSlice[index].Description = desc
	todoSlice[index].Completed = completed
	return nil
}

func (arr *TodoSlice) List() {
	arr2 := *arr
	var emoji string
	for i := 0; i < len(*arr); i++ {
		if arr2[i].Completed {
			emoji = "✅"
		} else {
			emoji = "❌"
		}
		date := arr2[i].DateAdded
		datStr := strconv.Itoa(date.Day()) + "/" + date.Month().String() + "/" + strconv.Itoa(date.Year())
		fmt.Printf("%v) %v\t\t%v\t\tCompleted: %v\t\tDate:%v\n", i+1, arr2[i].Title, arr2[i].Description, emoji, datStr)
	}
}

func (arr *TodoSlice) Save(filename string) error {
	data, err := json.MarshalIndent(*arr, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Load(filename string, data *TodoSlice) error {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File Does not exist")
		}
		return err
	}
	err = json.Unmarshal(fileData, data)
	if err != nil {
		return err
	}
	return nil
}
