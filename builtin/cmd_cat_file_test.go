package builtin

import "testing"

func Test_outputBlobType(t *testing.T) {
	type args struct {
		obj *Object
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{obj: &Object{
			Typ:  "blob",
			Size: 0,
			Data: Blob{Content: []byte("")},
		}}},
		{name: "case2", args: args{obj: &Object{
			Typ:  "blob",
			Size: 0,
			Data: Blob{Content: []byte("hello")},
		}}},
		{name: "case3", args: args{obj: &Object{
			Typ:  "blob",
			Size: 0,
			Data: Blob{Content: []byte("hello world")},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputBlobType(tt.args.obj)
		})
	}
}

func Test_outputTreeType(t *testing.T) {
	type args struct {
		obj *Object
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{obj: &Object{
			Typ:  "tree",
			Size: 0,
			Data: Tree{
				Id:    "1",
				Items: []Item{{
					SHA1: "s1",
					Typ:  "blob",
					Name: "file1.txt",
				}, {
					SHA1: "s2",
					Typ:  "tree",
					Name: "folder1",
				}, {
					SHA1: "s3",
					Typ:  "blob",
					Name: "file2.go",
				}},
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputTreeType(tt.args.obj)
		})
	}
}

func Test_outputCommitType(t *testing.T) {
	type args struct {
		obj *Object
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{obj: &Object{
			Typ:  "commit",
			Size: 0,
			Data: Commit{
				Id:        "c1",
				ParentId:  "c0",
				TreeId:    "t1",
				Author:    UserInfo{
					Name:      "Jack",
					Email:     "jack@google.com",
					Timestamp: "1660300855",
					TimeZone:  "+0800",
				},
				Committer: UserInfo{
					Name:      "Tom",
					Email:     "Tom@google.com",
					Timestamp: "1660300855",
					TimeZone:  "+0800",
				},
				Msg:       "commit msg",
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputCommitType(tt.args.obj)
		})
	}
}