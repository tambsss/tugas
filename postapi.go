package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Matakuliah
type Matakuliah struct {
	Kode_mk string `json:"kode_mk"`
	Name_mk string `json:"name_mk"`
	Sks     int    `json:"sks"`
}

// PostMatakuliah
func PostMatakuliah(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Mk Matakuliah
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Mk); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			getSks := r.PostFormValue("sks")
			sks, _ := strconv.Atoi(getSks)
			kode_mk := r.PostFormValue("kode_mk")
			name_mk := r.PostFormValue("name_mk")
			Mk = Matakuliah{
				Kode_mk: kode_mk,
				Name_mk: name_mk,
				Sks:     sks,
			}
		}
		dataMatakuliah, _ := json.Marshal(Mk)
		// to byte
		w.Write(dataMatakuliah)
		//cetak di browser
		return
	}
	http.Error(w, "hayo mau ngapain",
		http.StatusNotFound)
	return
}

func main() {
	http.HandleFunc("/Jadwal", PostMatakuliah)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
