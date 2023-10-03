package storage

import (
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"os"
	"strings"
)

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

func ensurePath(pathIncludingName string) error {
	items := strings.Split(pathIncludingName, "/")
	pathWithoutName := strings.Join(items[:len(items)-1], "/")
	if err := os.MkdirAll("storage/app/"+pathWithoutName, os.ModePerm); err != nil {
		return err
	}
	return nil
}
