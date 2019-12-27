package secret

import (
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

//no key and right path
func TestLoad(t *testing.T) {
	var v Vault
	v.encodingKey = ""
	v.filepath = "/home/gslab/goworkspace/src/gophercises/secret/cmd/secrets"
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
	v.filepath = "/home/gslab/goworkspace/src/gophercises/secret/cmd/secrets"
	File(v.encodingKey, v.filepath)
	v.Get(v.encodingKey)
}

//key value random but empty file to encrypt
func TestSetFail(t *testing.T) {
	v := File("key", "/home/gslab/goworkspace/src/gophercises/secret/s/")
	key, value := "k1", "v1"
	err := v.Set(key, value)
	if err == nil {
		t.Error("expecting error not getting error")
	}
}

func TestLoadRun(t *testing.T) {

	t.Run("for load to create error", func(t *testing.T) {
		v := File("", "/home/gslab/.secrets")

		err1 := v.Set("demoTest", "GetTest")
		_, err2 := v.Get("demoTest")

		if err1 != nil {
			t.Error("set not set")
		}

		assert.Equal(t, nil, err2)
	})
}
