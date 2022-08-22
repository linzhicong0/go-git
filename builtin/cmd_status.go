package builtin

import (
	"fmt"
	"github.com/spf13/cobra"
	"got/constant"
	"io/ioutil"
	"os"
)

func StatusRun(cmd *cobra.Command, args []string) {

	showChanged()

}

func showChanged() {

	// iter current working dir and check
	// 1. if the folder exist in cache?
	// 2. if the file modified or created?

	files, err := getAllFilesOfGivenFolder(constant.WorkingDir)

	if err != nil {
		fmt.Println(err)
		return
	}

	mFiles := make([]string, 5)
	untrackedFiles := make([]string, 5)

	for _, file := range files {

		fi, _ := os.Stat(file)

		rp := getRelativePath(file)

		if ce, ok := cacheEntries[rp]; ok {

			if changed := matchStat(&ce, &fi); changed {
				mFiles = append(mFiles, rp)
			}

		} else {
			untrackedFiles = append(untrackedFiles, rp)
		}
	}

	// Output the changed and untracked files

	fmt.Println(`Changes not staged for commit:
	(use "got add <file>..." to update what will be committed)
	(use "got restore <file>..." to discard changes in working directory)
	`)

	for _, mf := range mFiles {
		fmt.Printf("modified:\t\t%s", mf)
	}

	fmt.Println(`Untracked files:
	(use "got add <file>... to include in what will be committed")`)

	for _, utf := range untrackedFiles{
		fmt.Printf("%s", utf)
	}

}

func getAllFilesOfGivenFolder(folder string) ([]string, error) {

	fis, err := ioutil.ReadDir(folder)

	if err != nil {
		return nil, err
	}

	files := make([]string, 5)

	for _, fi := range fis {

		if fi.Name() == constant.DB_FOLDER || fi.Name() == ".git" {
			continue
		}

		if !fi.IsDir() {
			files = append(files, folder+"/"+fi.Name())
			continue
		}

		childes, err := getAllFilesOfGivenFolder(folder + "/" + fi.Name())

		if err != nil {
			return nil, err
		}

		files = append(files, childes...)

	}

	return files, nil

}

func getRelativePath(fileName string) string {
	return fileName[len(constant.WorkingDir):]
}

func matchStat(ce *CacheEntry, fi *os.FileInfo) bool {
	return ce.MTime == (*fi).ModTime().UnixMilli()

}
