package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

func Size(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)
	return len(content), err
}

func ExtName(fileName string) string {
	return path.Ext(fileName)
}

func IsNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

func IsPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

func DownloadFromUrl(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	_, err = io.Copy(out, resp.Body)
	return err
}

func MakeDirIfNotExist(src string, perm os.FileMode) error {
	if notExist := IsNotExist(src); notExist == true {
		if err := MakeDir(src, perm); err != nil {
			return err
		}
	}

	return nil
}

func MakeDir(src string, perm os.FileMode) error {
	err := os.MkdirAll(src, perm)
	if err != nil {
		return err
	}

	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := IsPermission(src)
	if perm == true {
		return nil, fmt.Errorf("IsPermission Permission denied src: %s", src)
	}

	ok := IsNotExist(src)
	if !ok {
		return nil, fmt.Errorf("IsNotExist src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
