package builtin

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/gob"
	"fmt"
	"got/constant"
	"io/ioutil"
	"os"
)

// Read the bytes of the sha1 file and decompressed it
func readSHA1File(sha1 string) ([]byte, error) {

	// objects/sha1[:1]/sha1[2:]
	fileName := constant.ObjectsFolder + "/" + sha1[:2] + "/" + sha1[2:]

	compressedBuff, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	b := bytes.NewReader(compressedBuff)

	r, err := zlib.NewReader(b)

	defer r.Close()

	//var deCompressedBuffer bytes.Buffer
	//io.Copy(&deCompressedBuffer, r)
	//
	//return deCompressedBuffer.Bytes(), nil

	ret, err := ioutil.ReadAll(r)
	return ret, err


}

// Read the bytes, which is in gob format, into *Object
func readObject(b []byte) *Object {

	obj := new(Object)

	buffer := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buffer)
	dec.Decode(obj)

	return obj
}


// Use the gob to serialized the Object
func writeObject(obj *Object) []byte {

	var buffer bytes.Buffer

	enc := gob.NewEncoder(&buffer)
	enc.Encode(obj)

	return buffer.Bytes()

}

// Hash the b as the file name and use zlib to compressed the b
func writeSHA1File(b []byte) error {

	var compressedBuffer bytes.Buffer

	// compressed the content
	w := zlib.NewWriter(&compressedBuffer)

	w.Write(b)

	w.Close()

	// generate the sha1 hash
	sha := sha1.New()

	sha.Write(compressedBuffer.Bytes())

	hex := fmt.Sprintf("%x", sha.Sum(nil))

	// write the file
	folder := constant.ObjectsFolder + "/" + hex[:2]

	fileName := folder + "/" + hex[2:]

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.MkdirAll(folder, 0777)
		if err != nil {
			fmt.Println(err)

		}
	}

	return ioutil.WriteFile(fileName, compressedBuffer.Bytes(), 0600)

}
