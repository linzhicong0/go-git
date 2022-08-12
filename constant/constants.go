package constant

import "os"

var WorkingDir string
var GotFolder string
var ObjectsFolder string
var RefsFolder string

func init(){

	WorkingDir, _ = os.Getwd()

	GotFolder = WorkingDir + DB_FOLDER

	ObjectsFolder = WorkingDir + OBJECTS_FOLDER

	RefsFolder = WorkingDir + REFS_FOLDER

}


const (
	DB_FOLDER = "/.got"
	INDEX_FILE = DB_FOLDER + "/index"
	HEADFILE = DB_FOLDER + "/HEAD"
	OBJECTS_FOLDER = DB_FOLDER + "/objects"
	REFS_FOLDER = DB_FOLDER + "/refs"

)

const (
	DEFAULT_HEAD = "ref: refs/heads/master"
)

