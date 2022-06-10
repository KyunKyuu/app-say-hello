package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i:=0; i<b.N; i++{
		HelloWorld("Teguh")
	}
}

//test main buat ngejalanin semua Unit test sekaligus dalam 1 package
func TestMain(m *testing.M){
	fmt.Println("Sebelum Unit Test")
	m.Run()
	fmt.Println("Setelah Unit Test")
}

func TestHelloWorldSkip(t *testing.T){
	if runtime.GOOS == "windows" {
		t.Skip("Gabisa kalau di windows")
	}

	result := HelloWorld("Teguh")
	require.Equal(t, "Hello Teguh", result, "Hasil nya ga sama")
	fmt.Println("Test Selesai")
}

func TestHelloWorldAssertion(t *testing.T){
	result := HelloWorld("Teguh")

	//Assert itu kalau error kode dibawah nya tetap di eksekusi
	// assert.Equal(t, "Hello Teguh", result, "Hasil nya ga sama")

	//require itu kalau error kode dibawah nya ga dieksekusi
	require.Equal(t, "Hello Teguh", result, "Hasil nya ga sama")
	fmt.Println("Test Selesai")
}

func TestHelloWorld(t *testing.T){
	result := HelloWorld("Teguh")
	
	if result != "Hello Teguh"{
		//error
		t.Error("Hasil nya beda")
	}
}