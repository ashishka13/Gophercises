package secret

import (
	"crypto/cipher"
	"errors"
	"io"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func myEncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
	return nil, errors.New("test generated encryptwriter error")
}

//no key and no path
func TestLoad(t *testing.T) {
	var v Vault
	v.encodingKey = ""
	v.filepath = "/home/gslab/goworkspace/src/gophercises/"
	v.load()
}

//no key no value
func TestSet(t *testing.T) {
	var v Vault
	v.Set("", "")
}

//no key no path no get key
func TestGet(t *testing.T) {

	v := File("", "")
	v.Get("")
}

// no key correct path
func TestSavePositive(t *testing.T) {
	var v Vault
	v.encodingKey = ""
	home, _ := homedir.Dir()

	v.filepath = filepath.Join(home, ".secrets")
	v.Set("demow", "abc123w")

}

//some random key correct path
func TestGetWitfile(t *testing.T) {
	var v Vault
	v.encodingKey = "asa"
	v.filepath = "/home/gslab/goworkspace/src/gophercises/gopheraccount.txt"
	File(v.encodingKey, v.filepath)
	v.Get(v.encodingKey)
}

//key value random but empty file to encrypt
func TestSetFail(t *testing.T) {
	v := File("key", "/home/gslab/goworkspace/src/gophercises/gopheraccount.txt")
	key, value := "k1", "v1"
	FakeEncryptwriter = myEncryptWriter
	err := v.Set(key, value)
	if err == nil {
		t.Error("expecting error not getting error")
	}
}

//first set some value then get some value
func TestLoadRun(t *testing.T) {
	t.Run("for load to create error", func(t *testing.T) {
		v := File("", "/home/gslab/.secrets")
		err := v.Set("demo11", "aaa111")
		v.Get("demo11")
		//this test is affected by above mocking func
		if err == nil {
			t.Error("expecting error not getting error")
		}
	})
}
