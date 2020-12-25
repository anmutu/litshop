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

	fmt.Println("root path", path)

	return path
}

func RootPathWithPostfix(p string) string {
	d := RootPath() + "/" + p
	fmt.Println("root path", d)
	return d
}

func StoragePath() string {
	return RootPath() + "/storage"
}
