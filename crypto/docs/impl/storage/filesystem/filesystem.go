package storage

import (
	"cryptoV2/storage"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"
)

const providerFileSystem = "filesystem"

type FileSystemProvider struct {
	name string
	path string
}

func NewFileSystemStorageProvider(name, path string) *FileSystemProvider {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	zap.S().Debugf("Created FileSystemStorageProvider at %s", path)
	return &FileSystemProvider{
		name: name,
		path: path,
	}
}

func (fsp *FileSystemProvider) Name() string {
	return fsp.name
}

func (fsp *FileSystemProvider) Type() string {
	return providerFileSystem
}

func (fsp *FileSystemProvider) List() []storage.ItemId {
	files, err := os.ReadDir(fsp.path)
	if err != nil {
		panic(err)
	}
	r := make([]storage.ItemId, len(files))
	for i, v := range files {
		item, _ := fsp.Get(fileNameToITemID(v.Name()))
		r[i] = item.Metadata().ItemId
	}
	return r
}

func (fsp *FileSystemProvider) Get(itemid storage.ItemId) (*storage.SecureItem, error) {
	filename := filepath.Join(fsp.path, fmt.Sprintf("%s_%s.json", itemid.UUID, itemid.Slot))
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var item storage.SecureItem
	err = json.Unmarshal(bytes, &item)
	return &item, err
}

func (fsp *FileSystemProvider) Set(item *storage.SecureItem) error {
	filename := filepath.Join(fsp.path, fmt.Sprintf("%s_%s.json", item.UUID(), item.Slot()))
	bytes, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, bytes, 0644)
}

func (fsp *FileSystemProvider) Remove(itemid storage.ItemId) error {
	filename := filepath.Join(fsp.path, fmt.Sprintf("%s_%s.json", itemid.UUID, itemid.Slot))
	return os.Remove(filename)
}

func fileNameToITemID(fileName string) storage.ItemId {
	parts := strings.Split(fileName, "_")
	uuid := parts[0]
	slot := strings.TrimSuffix(parts[1], ".json")
	return storage.ItemId{UUID: uuid, Slot: slot}
}
