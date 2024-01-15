package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Muhammad Hauzan Dini Fakhri _ 50421984

// Jadwal
type Jadwal struct {
	Kode_mk     int    `json:"kode"`
	Nama_matkul string `json:"matkul"`
	Sks         int    `json:"sks"`
}

// NewJadwal
func NewJadwal() []Jadwal {
	mhs := []Jadwal{
		Jadwal{
			Kode_mk:     045217,
			Nama_matkul: "Matematika Lanjut 1",
			Sks:         2,
		},
		Jadwal{
			Kode_mk:     8954263,
			Nama_matkul: "Struktur Data",
			Sks:         2,
		},
		Jadwal{
			Kode_mk:     045215,
			Nama_matkul: "Matematika Informatika 3",
			Sks:         2,
		},
		Jadwal{
			Kode_mk:     8412345,
			Nama_matkul: "Informatika Kesehatan",
			Sks:         2,
		},
		Jadwal{
			Kode_mk:     8532217,
			Nama_matkul: "Pengenalan Teknologi dan New Media",
			Sks:         2,
		},
	}
	return mhs
}

// GetJadwal
func GetJadwal(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		mhs := NewJadwal()
		datajadwal, err := json.Marshal(mhs)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(datajadwal)
		return
	}
	http.Error(w, "Access Denied",
		http.StatusNotFound)
}

func main() {
	http.HandleFunc("/Jadwal", GetJadwal)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
