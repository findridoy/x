package storage

import (
	"bytes"
	"os"
	"testing"
)

func TestPut(t *testing.T) {
	// Create a temporary file to use as input
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("failed to create temp file: %s", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write some data to the file
	testData := []byte("test data")
	if _, err := tmpFile.Write(testData); err != nil {
		t.Fatalf("failed to write to temp file: %s", err)
	}

	// Seek back to the beginning of the file
	if _, err := tmpFile.Seek(0, 0); err != nil {
		t.Fatalf("failed to seek to beginning of temp file: %s", err)
	}

	// Call the Put function with the temporary file
	pathIncludingName := "paht/to/the/testfile.txt"
	if err := Put(pathIncludingName, tmpFile); err != nil {
		t.Fatalf("Put failed: %s", err)
	}

	// Read the contents of the file that was written
	writtenData, err := os.ReadFile("storage/app/" + pathIncludingName)
	if err != nil {
		t.Fatalf("failed to read written file: %s", err)
	}

	// Verify that the contents of the file match the input data
	if !bytes.Equal(writtenData, testData) {
		t.Fatalf("written data does not match input data")
	}
}

func TestExists(t *testing.T) {
	// Create a temporary file to use as input
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("failed to create temp file: %s", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write some data to the file
	testData := []byte("test data")
	if _, err := tmpFile.Write(testData); err != nil {
		t.Fatalf("failed to write to temp file: %s", err)
	}

	// Seek back to the beginning of the file
	if _, err := tmpFile.Seek(0, 0); err != nil {
		t.Fatalf("failed to seek to beginning of temp file: %s", err)
	}

	// Call the Put function with the temporary file
	pathIncludingName := "paht/to/the/testfile.txt"
	if err := Put(pathIncludingName, tmpFile); err != nil {
		t.Fatalf("Put failed: %s", err)
	}

	exists, err := Exists(pathIncludingName)
	if err != nil {
		t.Fatalf("failed to check: %s", err)
	}

	if !exists {
		t.Fatalf("Exists failed:")
	}

	exists, err = Exists("h/h/sd.jpg")
	if err != nil {
		t.Fatalf("failed to check: %s", err)
	}

	if exists {
		t.Fatalf("Exists failed:")
	}

}

func TestDelete(t *testing.T) {
	// Create a temporary file to use as input
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("failed to create temp file: %s", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write some data to the file
	testData := []byte("test data")
	if _, err := tmpFile.Write(testData); err != nil {
		t.Fatalf("failed to write to temp file: %s", err)
	}

	// Seek back to the beginning of the file
	if _, err := tmpFile.Seek(0, 0); err != nil {
		t.Fatalf("failed to seek to beginning of temp file: %s", err)
	}

	// Call the Put function with the temporary file
	pathIncludingName := "paht/to/the/testfile.txt"
	if err := Put(pathIncludingName, tmpFile); err != nil {
		t.Fatalf("Put failed: %s", err)
	}

	if err := Delete(pathIncludingName); err != nil {
		t.Fatalf("Delete failed")
	}

	missing, err := Missing(pathIncludingName)
	if err != nil {
		t.Fatalf("Missing error: %s", err)
	}

	if !missing {
		t.Fatalf("Delete failed")
	}
}
