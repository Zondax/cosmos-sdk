package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type IStorageProvider interface {
	Name() string
	List() []string

	Get(name string) (*SecureItem, error)
	Set(item *SecureItem) error
	Remove(name string) error
}

type LocalFileSystem struct {
	name string
	path string
}

func NewLFSystem(name, path string) *LocalFileSystem {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	return &LocalFileSystem{
		name: name,
		path: path,
	}
}

func (l LocalFileSystem) Name() string {
	return l.name
}

func (l LocalFileSystem) List() []string {
	files, err := os.ReadDir(l.path)
	if err != nil {
		log.Fatal(err)
	}
	r := make([]string, len(files))
	for i, v := range files {
		r[i] = v.Name()
	}
	return r
}

func (l LocalFileSystem) Get(name string) (*SecureItem, error) {
	item, err := os.ReadFile(l.path + fmt.Sprintf("/%s", name))
	if err != nil {
		return nil, err
	}
	var sItem SecureItem
	err = json.Unmarshal(item, &sItem)
	return &sItem, err
}

func (l LocalFileSystem) Set(item *SecureItem) error {
	i, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return os.WriteFile(l.path+fmt.Sprintf("/%s", item.Name()), i, 0644)
}

func (l LocalFileSystem) Remove(name string) error {
	return os.Remove(l.path + fmt.Sprintf("/%s", name))
}
