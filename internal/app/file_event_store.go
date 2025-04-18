package app

import (
	"bufio"
	"encoding/json"
	"os"
)

type FileEventStore struct {
	Path string
}

func (f *FileEventStore) Append(event interface{}) error {
	file, err := os.OpenFile(f.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	return enc.Encode(event)
}

func (f *FileEventStore) Replay(apply func(event map[string]interface{})) error {
	file, err := os.Open(f.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var evt map[string]interface{}
		if err := json.Unmarshal(scanner.Bytes(), &evt); err == nil {
			apply(evt)
		}
	}
	return scanner.Err()
}
