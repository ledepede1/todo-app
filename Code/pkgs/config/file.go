package config

import (
	"fmt"
	"os"
)

func OpenTasksFile() os.File {
	file, err := os.OpenFile("./task.json", os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("Error in reading file: ", err)
	}

	return *file
}
