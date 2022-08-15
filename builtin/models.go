package builtin

import "got/constant"

// TODO: How can I know what's the type of the file
// ANSWER: When we are using the 'Id' we already know exactly what type the id is for.
// And I think we can put some info into the 'Id'!

/*

	- A Commit contain a Tree and other description information about the commit.

	- Tree is a representation of a folder, the commit will contain a tree, but
	  a Tree will contain Blob or Tree. We can take this as the same conception
	  about the folder-file ( Tree as folder and Blob as file).

	- Blob is contain the compressed content of the file.


							Commit
							  |
							  |
							Tree
							  |
							 / \
							/   \
						Blob	Tree
								  |
							 	 / \
								/   \
							  Blob   Tree
									   |
									  ...
 */


type Commit struct{
	Id string
	ParentId string
	TreeId string
	Author UserInfo
	Committer UserInfo
	Msg string
}

// Represent the object in index file
type IndexObject struct{
	Id string
	Typ constant.ObjectType
	File string
}

type Object struct {
	Typ constant.ObjectType
	Size int
	Data interface{}
}

// Description of the commit tree
// Representation of the folder
// Will be serialized using gob
type Tree struct {
	Id string
	BlobIds []string
	ChildIds []string
}


type Blob struct{
	Content []byte
}


type Log struct{
	PreCommitId string
	CurCommitId string
	TimeStamp string
	CommitMsg string
}


// Description of the author or committer
type UserInfo struct{
	Name string
	Email string
	Timestamp string
	TimeZone string
}