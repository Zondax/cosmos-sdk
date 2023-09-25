package storage

import (
	"os"
	"reflect"
	"testing"
)

func checkPathExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic("could not create dir")
		}
	}
}

func TestLocalFileSystem_Get(t *testing.T) {
	type fields struct {
		name string
		path string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    SecureItem
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LocalFileSystem{
				name: tt.fields.name,
				path: tt.fields.path,
			}
			got, err := l.Get(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalFileSystem_List(t *testing.T) {
	type fields struct {
		name string
		path string
	}
	tests := []struct {
		name   string
		fields fields
		setF   func(l LocalFileSystem)
		want   int
	}{
		{
			name: "empty provider",
			fields: fields{
				name: "LFS",
				path: os.TempDir() + "EmptyLFS",
			},
			setF: func(l LocalFileSystem) {

			},
			want: 0,
		},
		{
			name: "with items",
			fields: fields{
				name: "LFS",
				path: os.TempDir() + "List",
			},
			setF: func(l LocalFileSystem) {
				item := &SecureItem{
					TypeUuid: "testItem",
					name:     "check",
					blob:     nil,
				}
				l.Set(item)
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkPathExists(tt.fields.path)
			l := LocalFileSystem{
				name: tt.fields.name,
				path: tt.fields.path,
			}
			tt.setF(l)
			got := l.List()
			if len(got) != tt.want {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalFileSystem_Remove(t *testing.T) {
	type fields struct {
		name string
		path string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		setF    func(l LocalFileSystem)
		wantErr bool
	}{
		{
			name: "remove item",
			fields: fields{
				name: "LFS",
				path: os.TempDir() + "removeLFS",
			},
			args: args{name: "toRemove"},
			setF: func(l LocalFileSystem) {
				item := &SecureItem{
					TypeUuid: "testItem",
					name:     "toRemove",
					blob:     nil,
				}
				l.Set(item)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LocalFileSystem{
				name: tt.fields.name,
				path: tt.fields.path,
			}
			checkPathExists(tt.fields.path)
			tt.setF(l)
			if len(l.List()) < 1 {
				t.Errorf("could not create file")
			}
			if err := l.Remove(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(l.List()) > 0 {
				t.Errorf("could not create file")
			}
		})
	}
}

func TestLocalFileSystem_Set(t *testing.T) {
	type fields struct {
		name string
		path string
	}
	type args struct {
		item *SecureItem
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "set item",
			fields: fields{
				name: "LFS",
				path: os.TempDir() + "setItem",
			},
			args: args{item: &SecureItem{
				TypeUuid: "testItem",
				name:     "setTest",
				blob:     nil,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LocalFileSystem{
				name: tt.fields.name,
				path: tt.fields.path,
			}
			if err := l.Set(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewLFSystem(t *testing.T) {
	type args struct {
		name string
		path string
	}
	tests := []struct {
		name string
		args args
		want *LocalFileSystem
	}{
		{
			name: "create provider",
			args: args{
				name: "testLFS",
				path: os.TempDir() + "LFS",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := os.Stat(tt.args.path); os.IsNotExist(err) {
				err := os.Mkdir(tt.args.path, os.ModePerm)
				if err != nil {
					t.Errorf("could not create path")
				}
			}
			got := NewLFSystem(tt.args.name, tt.args.path)
			if got == nil {
				t.Errorf("nil provider")
			}
		})
	}
}
