package hash

import (
	"testing"
)

func TestHashMD5(t *testing.T) {
	str := "Hello, World!"
	expected := "65a8e27d8879283831b664bd8b7f0ad4"

	result := HashMD5(str)

	if result != expected {
		t.Errorf("MD5(%s) = %s, expected %s", str, result, expected)
	}
}

func TestHashSHA1(t *testing.T) {
	str := "Hello, World!"
	expected := "0a0a9f2a6772942557ab5355d76af442f8f65e01"

	result := HashSHA1(str)

	if result != expected {
		t.Errorf("SHA1(%s) = %s, expected %s", str, result, expected)
	}
}

func TestHashCRC32(t *testing.T) {
	str := "Hello, World!"
	expected := uint32(3964322768)

	result := HashCRC32(str)

	if result != expected {
		t.Errorf("CRC32(%s) = %d, expected %d", str, result, expected)
	}
}
