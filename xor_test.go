package xor

import (
	"reflect"
	"testing"
)

var (
	benchString string
	benchBytes  []byte
)

func TestEncryptDecrypt(t *testing.T) {
	key := []byte("b7a92c1e-83fe-11e9-83be-be1a02408b06")
	data := []byte("A boy fell in the mud")

	enc := EncryptDecrypt(data, key)

	plaintext := EncryptDecrypt(enc, key)

	if !reflect.DeepEqual(plaintext, data) {
		t.Errorf("got %s, want %s", plaintext, data)
	}
}

func TestEncryptDecryptString(t *testing.T) {
	key := "b7a92c1e-83fe-11e9-83be-be1a02408b06"
	data := "A boy fell in the mud"

	enc := EncryptDecryptString(data, key)

	plaintext := EncryptDecryptString(enc, key)

	if plaintext != data {
		t.Errorf("got %s, want %s", plaintext, data)
	}
}

func BenchmarkEncryptDecrypt(b *testing.B) {
	key := []byte("b7a92c1e-83fe-11e9-83be-be1a02408b06")
	data := []byte("A boy fell in the mud")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		benchBytes = EncryptDecrypt(data, key)
	}
}

func BenchmarkEncryptDecryptString(b *testing.B) {
	key := "b7a92c1e-83fe-11e9-83be-be1a02408b06"
	data := "A boy fell in the mud"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		benchString = EncryptDecryptString(data, key)
	}
}
