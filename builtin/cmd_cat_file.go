package builtin

import (
	"fmt"
	"github.com/spf13/cobra"
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
}

func catFileType(sha1 string) {

	fmt.Println("cat-file -t")
}

func catFileSize(sha1 string) {

	fmt.Println("cat-file -s")
}


