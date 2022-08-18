package builtin

import (
	"fmt"
	"github.com/spf13/cobra"
	"got/constant"
	"strconv"
)

// got cat-file -[p,t] <sha1>
func CatFileRun(cmd *cobra.Command, args []string) {

	sha1 := args[0]

	if p, _ := strconv.ParseBool(cmd.Flag("pretty").Value.String()); p {
		catFilePretty(sha1)
	} else if t, _ := strconv.ParseBool(cmd.Flag("type").Value.String()); t {
		catFileType(sha1)
	} else if s, _ := strconv.ParseBool(cmd.Flag("size").Value.String()); s {
		catFileSize(sha1)
	} else {
		cmd.Usage()
	}

}

// 1. Take the first 2 chars as the folder and the left 38 chars as the file name
// 2. Read the file
// 3. Read until the first space -> type
// 4. Read until the \0 -> size

func catFilePretty(sha1 string) {

	fmt.Println("cat-file -p")

	// TODO: support using the first 4 hex to read
	buf, err := readSHA1File(sha1)

	if err != nil {
		fmt.Println(err)
		return
	}

	obj := readObject(buf)

	switch obj.Typ {
	case constant.BLOB_TYPE:
		outputBlobType(obj)
		break
	case constant.TREE_TYPE:
		outputTreeType(obj)
		break
	case constant.COMMIT_TYPE:
		outputCommitType(obj)
		break
	case constant.TAG_TYPE:
		outputTagType(obj)
		break
	}
}

// output the given sha1 file type (tree, blob, commit or tag)
func catFileType(sha1 string) {

	fmt.Println("cat-file -t")

	// TODO: support using the first 4 hex to read
	buf, err := readSHA1File(sha1)

	if err != nil {
		fmt.Printf("reading the given sha1 file type error, please check the sha1: %s", sha1)
		return
	}

	obj := readObject(buf)

	fmt.Println(obj.Typ)

}

// output the given sha1 file size
func catFileSize(sha1 string) {

	fmt.Println("cat-file -s")

	buf, err := readSHA1File(sha1)

	if err != nil {
		fmt.Printf("reading given sha1 file size error, please check the sha1: %s", sha1)
		return
	}

	obj := readObject(buf)

	fmt.Println(obj.Size)
}

// output the file content
func outputBlobType(obj *Object) {

	blob := obj.Data.(Blob)

	fileContent := string(blob.Content)

	fmt.Println(fileContent)

}

// output the tree info
func outputTreeType(obj *Object) {

	tree := obj.Data.(Tree)

	for _, item := range tree.Items {
		fmt.Printf("%s %s\t%s\n", item.Typ, item.SHA1, item.Name)
	}

}

// output the commit info
func outputCommitType(obj *Object) {

	commit := obj.Data.(Commit)


	/*
		tree a7c87387ee8a0eaf1258d83b90d6508b9704b3e5
		parent 5dcc0e57b78005687e8de265ba271646d9d8d6e7
		author Jack Lin <446026340@qq.com> 1660300855 +0800
		committer GitHub <noreply@github.com> 1660300855 +0800

		commit message
	*/
	fmt.Printf("%s %s\n", "tree", commit.TreeId)
	fmt.Printf("%s %s\n", "parent", commit.ParentId)
	fmt.Printf("%s %s <%s> %s %s\n", "author", commit.Author.Name, commit.Author.Email, commit.Author.Timestamp, commit.Author.TimeZone)
	fmt.Printf("%s %s <%s> %s %s\n", "committer", commit.Committer.Name, commit.Committer.Email, commit.Committer.Timestamp, commit.Committer.TimeZone)
	fmt.Printf("\n%s\n", commit.Msg)


}

// output the tag info
func outputTagType(obj *Object) {
	fmt.Println("sorry, it's not implemented yet")
}


