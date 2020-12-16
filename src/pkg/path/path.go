package path

import (
	"fmt"
	"os"
)

func RootPath() string {
	var (
		path string
		err  error
	)

	defer func() {
		if err != nil {
			panic(fmt.Sprintf("GetProjectRoot error :%+v", err))
		}
	}()

	path, _ = os.Getwd()
	return path
}

func StoragePath() string {
	return RootPath() + "/storage"
}