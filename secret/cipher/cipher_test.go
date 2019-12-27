package cipher

import (
	"crypto/cipher"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ciphertStruct struct {
	err error
}

func (m *ciphertStruct) MyCipherBlock(key string) (cipher.Block, error) {
	return nil, m.err
}

func myReadFull(r io.Reader, buf []byte) (n int, err error) {
	return 0, errors.New("testing generated error")
}

func generateTempfile() string {
	content := []byte("aa dd gg hh jj")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
	return tmpfile.Name()
}

//test with real file
func TestDecryptReader(t *testing.T) {
	f, _ := os.Open("/home/gslab/goworkspace/src/gophercises/secret/secret")

	DecryptReader("er", f)
}

//test with fake file and random key
func TestDecryptReaderErr(t *testing.T) {
	f, _ := os.Open(generateTempfile())

	DecryptReader("r56r76t87y", f)
}

//empty key and empty text file
func TestEncryptWriter(t *testing.T) {
	f, _ := os.OpenFile("/home/gslab/goworkspace/src/gophercises/secret/secret", os.O_RDWR, 0000)
	EncryptWriter("", f)
}

//no file provided to encryptwriter
func TestEncryptWriterErr(t *testing.T) {
	f, _ := os.Open("")
	EncryptWriter("", f)
}

//cipherblock mocking
func TestNewCipherBlock(t *testing.T) {
	f := &ciphertStruct{err: errors.New("User defined error cipher")}
	FakeCipherBlock = f.MyCipherBlock
	fl, _ := os.OpenFile("/home/gslab/goworkspace/src/gophercises/secret/secret", os.O_RDWR, 0755)

	_, err := EncryptWriter("asgyua", fl)
	assert.NotEqual(t, err, nil)
}

//mock the decryptstream function read the mock flow carefully
func TestDecryptReaderMock(t *testing.T) {
	f3, _ := os.Open("/home/gslab/goworkspace/src/gophercises/secret/secret")
	DecryptReader("", f3)

	f := &ciphertStruct{err: errors.New("User defined error cipher")}
	FakeCipherBlock = f.MyCipherBlock

	iv := []byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	_, err := decryptStream("", iv)

	if err == nil {
		t.Error("expecting error not getting error")
	}
}

//random byte stream of size 16
func TestEncrypt(t *testing.T) {
	iv := []byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	decryptStream("", iv)
}

func TestMockIOReader(t *testing.T) {
	f, _ := os.Open("")
	FakeReadfull = myReadFull
	EncryptWriter("", f)
}
