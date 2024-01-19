package tasks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ledepede1/fasttask/Code/pkgs/config"
	"github.com/ledepede1/fasttask/Code/pkgs/middleware"
)

type ReqBody struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

func SaveTaskHandler(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w, r)

	var rows Rows
	var reqBody ReqBody

	file := config.OpenTasksFile()
	defer file.Close()

	decoder := json.NewDecoder(&file)
	decoder.Decode(&rows)

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(reqBody.Id)

	if searchRows(reqBody.Id, rows) {
		var beforeString, err = io.ReadAll(&file)
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

		if rows[reqBody.Id-1].Checked {
			rows[reqBody.Id-1].Checked = false
		} else {
			rows[reqBody.Id-1].Checked = true
		}
		err = json.NewEncoder(&file).Encode(rows)
		if err != nil {
			file.Write(beforeString)
		}

	} else {
		fmt.Println("ID doesnt exist")
	}

}

func searchRows(id int, rows Rows) bool {
	for _, v := range rows {
		if v.Id == id {
			return true
		}
	}
	return false
}
