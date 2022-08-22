package builtin

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"got/constant"
	"io/ioutil"
)

// This is the struct representing the .got/index file
// This will be gob format in index file
// Use file name as the key
var cacheEntries map[string]CacheEntry

func init(){

}

// Load the cacheEntries from .got/index file
func loadCache(){

	buf, err := ioutil.ReadFile(constant.IndexFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	dec := gob.NewDecoder(bytes.NewBuffer(buf))

	dec.Decode(cacheEntries)

}