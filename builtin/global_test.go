package builtin

import (
	"bytes"
	"encoding/gob"
	"reflect"
	"testing"
)

func Test_readObject(t *testing.T) {
	type args struct {
		b []byte
	}

	tests := []struct {
		name string
		args args
		want *Object
	}{
		// TODO: Add test cases.
		{name: "tree", args: args{b: encodeObject(&Object{
			Typ:  "tree",
			Size: 0,
			Data: nil,
		})}, want: &Object{
			Typ:  "tree",
			Size: 0,
			Data: nil,
		}},
		{name: "commit", args: args{b: encodeObject(&Object{
			Typ:  "commit",
			Size: 0,
			Data: 1,
		})}, want: &Object{
			Typ:  "commit",
			Size: 0,
			Data: 1,
		}},
		{name: "blob", args: args{b: encodeObject(&Object{
			Typ:  "blob",
			Size: 1,
			Data: 1,
		})}, want: &Object{
			Typ:  "blob",
			Size: 1,
			Data: 1,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readObject(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func encodeObject(obj *Object) []byte {

	buffer := bytes.NewBuffer(make([]byte, 0))

	gob.NewEncoder(buffer).Encode(obj)
	return buffer.Bytes()
}

func Test_writeObject(t *testing.T) {
	type args struct {
		obj *Object
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{name: "tree", args: args{obj: &Object{
			Typ:  "tree",
			Size: 0,
			Data: nil,
		}}, want: encodeObject(&Object{
			Typ:  "tree",
			Size: 0,
			Data: nil,
		})},

		{name: "commit", args: args{obj: &Object{
			Typ:  "commit",
			Size: 100,
			Data: nil,
		}}, want: encodeObject(&Object{
			Typ:  "commit",
			Size: 100,
			Data: nil,
		})},

		{name: "blob", args: args{obj: &Object{
			Typ:  "blob",
			Size: 0,
			Data: 200,
		}}, want: encodeObject(&Object{
			Typ:  "blob",
			Size: 0,
			Data: 200,
		})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := writeObject(tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("writeObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeSHA1File(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{b: encodeObject(&Object{
				Typ:  "",
				Size: 0,
				Data: nil,
			})},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeSHA1File(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("writeSHA1File() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_readSHA1File(t *testing.T) {
	type args struct {
		sha1 string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "normal", args: args{sha1: "d4ca003c19945e892ded7d30e72f832e75c99eba"}, want: encodeObject(&Object{
			Typ:  "",
			Size: 0,
			Data: nil,
		}), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readSHA1File(tt.args.sha1)
			if (err != nil) != tt.wantErr {
				t.Errorf("readSHA1File() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readSHA1File() got = %v, want %v", got, tt.want)
			}
		})
	}
}
