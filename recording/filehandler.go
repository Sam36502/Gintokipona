package recording

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	DEFAULT_RECORD_FILENAME = "ginponadata.json"
	RECORD_FILE_PERMISSIONS = 0644
)

func LoadRecords(filename string) (*RecordFile, error) {

	// Read file contents
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("[Error] Failed to open record file '%s':\n  %s\n", filename, err.Error())
		return nil, err
	}

	// Parse Contents
	var recordfile RecordFile
	err = json.Unmarshal(data, &recordfile)
	if err != nil {
		fmt.Printf("[Error] Failed to parse record file '%s':\n  %s\n", filename, err.Error())
		fmt.Println("  Is the file format correct?")
		return nil, err
	}

	return &recordfile, nil
}

func SaveRecords(filename string, recordfile *RecordFile) error {

	// Marshal data
	data, err := json.Marshal(recordfile)
	if err != nil {
		fmt.Printf("[Error] Failed to store data as JSON:\n  %s\n", err.Error())
		return err
	}

	// Save data to file
	err = ioutil.WriteFile(filename, data, RECORD_FILE_PERMISSIONS)
	if err != nil {
		fmt.Printf("[Error] Failed to store data in '%s':\n  %s\n", filename, err.Error())
		return err
	}

	return nil
}
