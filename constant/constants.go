package constant

import "os"

var WorkingDir string
var GotFolder string
var ObjectsFolder string
var RefsFolder string
var RefsHeadsFolder string
var RefsTagsFolder string

func init(){

	WorkingDir, _ = os.Getwd()

	GotFolder = WorkingDir + DB_FOLDER

	ObjectsFolder = WorkingDir + OBJECTS_FOLDER

	RefsFolder = WorkingDir + REFS_FOLDER

	RefsHeadsFolder = WorkingDir + REFS_HEADS_FOLDER

	RefsTagsFolder = WorkingDir + REFS_TAGS_FOLDER

}


const (
	DB_FOLDER = "/.got"
	INDEX_FILE = DB_FOLDER + "/index"
	HEADFILE = DB_FOLDER + "/HEAD"
	OBJECTS_FOLDER = DB_FOLDER + "/objects"
	REFS_FOLDER = DB_FOLDER + "/refs"
	REFS_HEADS_FOLDER = REFS_FOLDER + "/heads"
	REFS_TAGS_FOLDER = REFS_FOLDER + "/tags"

)

const (
	DEFAULT_HEAD = "ref: refs/heads/master"
)

type ObjectType int

const (
	COMMIT_TYPE ObjectType = iota
	TREE_TYPE ObjectType = iota
	BLOB_TYPE ObjectType = iota
)

