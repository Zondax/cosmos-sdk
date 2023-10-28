package storage

import (
	"cryptoV2/storage"
	"os"
	"reflect"
	"testing"
)

func TestFileSystemProvider(t *testing.T) {
	name := "testProvider"
	path := os.TempDir() + "/testProvider"
	provider := NewFileSystemStorageProvider(name, path)
	defer os.RemoveAll(path)

	// Test Name
	if provider.Name() != name {
		t.Errorf("Name() = %v, want %v", provider.Name(), name)
	}

	// Test Set and Get
	item := storage.NewSecureItem(storage.ItemId{UUID: "testUUID", Slot: "testSlot"}, []byte("testBlob"))
	err := provider.Set(item)
	if err != nil {
		t.Errorf("Set() error = %v", err)
	}

	gotItem, err := provider.Get(item.Metadata().ItemId)
	if err != nil {
		t.Errorf("Get() error = %v", err)
	}
	if !reflect.DeepEqual(gotItem, item) {
		t.Errorf("Get() got = %v, want %v", gotItem, item)
	}

	// Test List
	items := provider.List()
	if len(items) != 1 {
		t.Errorf("List() = %v, want %v", len(items), 1)
	}

	// Test Remove
	err = provider.Remove(item.Metadata().ItemId)
	if err != nil {
		t.Errorf("Remove() error = %v", err)
	}

	items = provider.List()
	if len(items) != 0 {
		t.Errorf("List() = %v, want %v", len(items), 0)
	}
}
