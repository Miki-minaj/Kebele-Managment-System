package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("he welcome"))
}

func Testhandler(t *testing.T) {
	httprr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/index", nil)
	if err != nil {
		t.Fatal(err)
	}
	index(httprr, req)
	resp := httprr.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "hey welcome" {
		t.Errorf("its not what you want")
	}
	// if resp.StatusCode != http.StatusOK {
	// 	t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	// }
}

// func customDate(t time.Time) string {
// 	return t.UTC().Format("02 Jan 2006 at 15:04")
// }
// func sum(x int, y int) int {
// 	return x + y
// }

// func TestCustomdate(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		x    int
// 		y    int
// 		want int
// 	}{
// 		{
// 			name: "first test",
// 			x:    2,
// 			y:    3,
// 			want: 5,
// 		},
// 		{
// 			name: "second test",
// 			x:    4,
// 			y:    6,
// 			want: 5,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := sum(tt.x, tt.y)
// 			if tt.want != got {
// 				t.Errorf("expected %d; got %d", tt.want, got)
// 			}
// 		})
// 	}
// }
