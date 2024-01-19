package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ledepede1/fasttask/Code/pkgs/config"
	"github.com/ledepede1/fasttask/Code/pkgs/middleware"
)

type Rows []Task

func GetListHandler(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w, r)

	var rows Rows

	file := config.OpenTasksFile()
	defer file.Close()

	decoder := json.NewDecoder(&file)
	decoder.Decode(&rows)

	jsonData, err := json.Marshal(rows)
	if err != nil {
		fmt.Println("Error in marshaling: ", err)
	}

	w.Write(jsonData)
}
