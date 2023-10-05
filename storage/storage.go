package storage

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"os"
	"strings"
)

const basePath = "storage/app/"

func Put(pathIncludingName string, file multipart.File) error {
	if err := ensurePath(pathIncludingName); err != nil {
		return err
	}

	f, err := os.Create("storage/app/" + pathIncludingName)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			slog.Error(fmt.Sprintf("closing file: %s: %s", pathIncludingName, err.Error()))
		}
	}(f)

	if _, err := io.Copy(f, file); err != nil {
		return err
	}
	return nil
}

func Exists(pathIncludingName string) (bool, error) {
	f, err := os.Open(basePath + pathIncludingName)
	if err != nil {
		if errors.Is(os.ErrNotExist, err) || os.IsNotExist(err) {
			return false, nil
		}
		fmt.Printf("%t", err)
		return false, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			slog.Error("closing file: %w", err)
		}
	}(f)
	return true, nil
}

func Missing(pathIncludingName string) (bool, error) {
	exists, err := Exists(pathIncludingName)
	if err != nil {
		return false, err
	}
	return !exists, nil
}

func Delete(pathIncludingName string) error {
	return os.Remove(basePath + pathIncludingName)
}

func ensurePath(pathIncludingName string) error {
	items := strings.Split(pathIncludingName, "/")
	pathWithoutName := strings.Join(items[:len(items)-1], "/")
	if err := os.MkdirAll("storage/app/"+pathWithoutName, os.ModePerm); err != nil {
		return err
	}
	return nil
}
