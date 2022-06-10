package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

//Benchmart Test
func BenchmarkTableKobo(b *testing.B) {
	benchmarks := []struct{
		name string
		request string
	}{
		{
			name : "boo",
			request: "koboo",
		},
		{
			name : "kooo",
			request: "Demoon",
		},
		{
			name : "chann",
			request: "Lord",
		},
	}

	for _,test := range benchmarks {
		b.Run(test.name, func(b *testing.B){
			for i:=0; i<b.N; i++{
				Kobo(test.request)
			}
		})  
			
		}
	}


//Table test
func TestTableKobo(t *testing.T){
	tests := []struct{
		name string
		request string
		harapan string
	}{
		{
			name: "Teguh",
			request:"Teguh",
			harapan:"Yang mulia Teguh",
		},
		{
			name: "Iqbal",
			request:"Iqbal",
			harapan:"Yang mulia Iqbal",
		},
		{
			name: "Yoga",
			request:"Yoga",
			harapan:"Yang mulia Yoga",
		},
	}

	for _,test := range tests {
		t.Run(test.name, func(t *testing.T){
			result := Kobo(test.request)
			require.Equal(t,test.harapan,result)
		})
	}
}


func TestKobo(t *testing.T){
	result := Kobo("tempest")
	require.Equal(t,"Yang mulia tempest", result, "Lah hasilnya ga sama")

	fmt.Println("Test nya selesai")
}