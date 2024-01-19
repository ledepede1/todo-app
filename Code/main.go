package main

import (
	"net/http"

	"github.com/ledepede1/fasttask/Code/pkgs/tasks"
)

func main() {
	http.HandleFunc("/backend/savetask", tasks.SaveTaskHandler)
	http.HandleFunc("/backend/getlist", tasks.GetListHandler)
	http.HandleFunc("/backend/deletetask", tasks.DeleteTaskHandler)
	http.HandleFunc("/backend/addtask", tasks.AddTaskHandler)

	http.ListenAndServe(":8080", nil)
}
