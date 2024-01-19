package tasks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ledepede1/fasttask/Code/pkgs/config"
	"github.com/ledepede1/fasttask/Code/pkgs/middleware"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w, r)

	var rows Rows
	var reqBody ReqBody

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

	for _, v := range rows {
		if v.Id != reqBody.Id {
			updatedRows = append(updatedRows, v)
		}
	}

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

	if len(updatedRows) == 0 {
		err = json.NewEncoder(&file).Encode([]Rows{})
		if err != nil {
			file.Write(beforeString)
		}
	} else {
		err = json.NewEncoder(&file).Encode(updatedRows)
		if err != nil {
			file.Write(beforeString)
		}
	}
}
