package wibu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnime(t *testing.T) {
	tests := []struct{
		Name string
		Request string
		Harapan string
	}{
		{
			Name:"Naruto",
			Request: "Naruto",
			Harapan: "Naruto itu anime bagus",
		},
		{
			Name:"One Piece",
			Request: "One Piece",
			Harapan: "One Piece itu anime bagus",
		},
	}

	for _,test := range tests{
		t.Run(test.Name, func(t *testing.T){
			result := Anime(test.Request)
			assert.Equal(t, test.Harapan, result)
		})
	}
}  