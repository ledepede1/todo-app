package tasks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ledepede1/fasttask/Code/pkgs/config"
	"github.com/ledepede1/fasttask/Code/pkgs/middleware"
)

type Task struct {
	Id      int    `json:"id"`
	Text    string `json:"text"`
	Checked bool   `json:"checked"`
}

type NewTask struct {
	Task
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w, r)

	var rows Rows
	var reqBody ReqBody
	var newTask NewTask

	var updatedRows Rows

	file := config.OpenTasksFile()
	defer file.Close()

	decoder := json.NewDecoder(&file)
	decoder.Decode(&rows)

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTask.Id = GetMaxId(rows) + 1
	newTask.Text = reqBody.Text

	updatedRows = append(updatedRows, rows...)
	updatedRows = append(updatedRows, newTask.Task)

	beforeString, err := io.ReadAll(&file)
	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	err = file.Truncate(0)
	if err != nil {
		fmt.Println(err)
	}

	err = json.NewEncoder(&file).Encode(updatedRows)
	if err != nil {
		file.Write(beforeString)
	}
}

func GetMaxId(rows Rows) int {
	maxId := 0
	for _, task := range rows {
		if task.Id > maxId {
			maxId = task.Id
		}
	}
	return maxId
}
