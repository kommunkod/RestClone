package main_test

import (
	"os"
	"path/filepath"
	"testing"
)

const FtpDataDirectory = "containers/files"

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func ClearDirectory(t *testing.T) {
	t.Log("Clearing Directory...")

	err := RemoveContents(FtpDataDirectory)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Directory cleared.")
}

func CreateFile(t *testing.T, name string, content string) {
	t.Logf("Creating file %s in Directory...", name)

	f, err := os.Create(FtpDataDirectory + "/" + name)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("File %s created in Directory.", name)
}

func RemoveFile(t *testing.T, name string) {
	t.Logf("Removing file %s from Directory...", name)

	err := os.Remove(FtpDataDirectory + "/" + name)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("File %s removed from Directory.", name)
}

func ReadFile(t *testing.T, name string) string {
	t.Logf("Reading file %s from Directory...", name)

	data, err := os.ReadFile(FtpDataDirectory + "/" + name)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("File %s read from Directory.", name)
	return string(data)
}
