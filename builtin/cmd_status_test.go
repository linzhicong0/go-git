package builtin

import (
	"got/constant"
	"reflect"
	"testing"
)

func Test_showChanged(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			showChanged()
		})
	}
}

func Test_getAllFilesOfGivenFolder(t *testing.T) {
	type args struct {
		folder string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{folder: constant.WorkingDir}, want: []string{

		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAllFilesOfGivenFolder(tt.args.folder)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllFilesOfGivenFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllFilesOfGivenFolder() got = %v, want %v", got, tt.want)
			}
		})
	}
}