package file

import (
	"github.com/SantoDE/datahamster/storage"
	"io"
	"os"
	"testing"
)

func TestSaveResultOk(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	err := Copy("../../integration/fixtures/dump_test_table.sql", "/tmp/dump_test_table.sql")

	f := new(storage.File)
	f.Path = "/tmp/dump_test_table.sql"
	f.Name = "test_file"

	storage := NewFileStorage("/tmp")
	storage.SaveFile(*f)

	file, err := os.Open("/tmp/test_file")

	if err != nil {
		t.Fatalf("Error while opening the file: %s", err)
	}

	info, _ := file.Stat()

	if info.Size() <= 0 {
		t.Fatalf("Error Dumping: Got an empty file - no data saved")
	}

	_, err = os.Open("/tmp/dump_test_table.sql")

	if err == nil {
		t.Fatalf("Unexpected no error while opening the file: %s", err)
	}
}

// Copies a file.
func Copy(src string, dst string) error {
	// Open the source file for reading
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	// Open the destination file for writing
	d, err := os.Create(dst)
	if err != nil {
		return err
	}

	// Copy the contents of the source file into the destination file
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}

	// Return any errors that result from closing the destination file
	// Will return nil if no errors occurred
	return d.Close()
}
