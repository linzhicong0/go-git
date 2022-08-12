package builtin

import (
	"fmt"
	"github.com/spf13/cobra"
	"got/constant"
	"io/ioutil"
	"os"
)

func InitRun(cmd *cobra.Command, args []string) {



	createDBFolders()

	err := ioutil.WriteFile(constant.WorkingDir + constant.HEADFILE, []byte(constant.DEFAULT_HEAD), 0600)
	if err == nil {
		fmt.Println(err)
	}
}

// Create related folders
// .got/
// .got/objects/
// .got/refs/
func createDBFolders() error{


	var err error
	// create .got/
	err = os.Mkdir(constant.GotFolder, 0700)
	// create .got/objects/
	err = os.Mkdir(constant.ObjectsFolder, 0700)
	// create .got/refs/
	err = os.Mkdir(constant.RefsFolder, 0700)

	fmt.Println(err)
	// TODO: handle error
	return nil
}