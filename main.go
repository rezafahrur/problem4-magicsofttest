package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type museum struct {
	Data []mData `json:"data"`
}

type mData struct {
	MuseumID          string `json:"museum_id"`
	KodePengelolaan   string `json:"kode_pengelolaan"`
	Nama              string `json:"nama"`
	Sdm               string `json:"sdm"`
	AlamatJalan       string `json:"alamat_jalan"`
	DesaKelurahan     string `json:"desa_kelurahan"`
	Kecamatan         string `json:"kecamatan"`
	KabupatenKota     string `json:"kabupaten_kota"`
	Lintang           string `json:"lintang"`
	Bujur             string `json:"bujur"`
	Koleksi           string `json:"koleksi"`
	SumberDana        string `json:"sumber_dana"`
	Pengelola         string `json:"pengelola"`
	Tipe              string `json:"tipe"`
	Standar           string `json:"standar"`
	TahunBerdiri      string `json:"tahun_berdiri"`
	Bangunan          string `json:"bangunan"`
	LuasTanah         string `json:"luas_tanah"`
	StatusKepemilikan string `json:"status_kepemilikan"`
	Propinsi          string `json:"propinsi"`
}

func main() {
	response, err := http.Get("http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?nama=museum")

	if err != nil {
		fmt.Println("Error!", err.Error())
	}

	defer response.Body.Close()

	isi, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error!", err.Error())
	}
	//karena server mengirimkan teks UTF-8
	isi = bytes.TrimPrefix(isi, []byte("\xef\xbb\xbf"))

	var m museum
	err = json.Unmarshal(isi, &m)

	if err != nil {
		fmt.Println("Error!", err.Error())
	}
	//buat file malang
	malang, err := os.Create("kota Malang.csv")
	jakpus, err := os.Create("kota Jakarta Pusat.csv")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	defer malang.Close()
	defer jakpus.Close()

	tulisMalang := csv.NewWriter(malang)
	tulisJakpus := csv.NewWriter(jakpus)

	for _, data := range m.Data {
		if data.KabupatenKota == "Kota Malang" {
			var mlg []string
			// mlg = append(mlg[:0], "ID Museum", "Kode Pengelolaan", "Nama Museum", "SDM", "Alamat", "Desa/Kelurahan", "Kecamatan", "Kota", "Propinsi", "Lintang", "Bujur", "Koleksi", "Sumber Dana", "Pengelola", "Tipe", "Standar", "Tahun Berdiri", "Bangunan", "Luas Tanah", "Status Kepemilikan")
			mlg = append(mlg, data.MuseumID, data.KodePengelolaan, data.Nama, data.Sdm, data.AlamatJalan,
				data.DesaKelurahan, data.Kecamatan, data.KabupatenKota, data.Propinsi, data.Lintang, data.Bujur,
				data.Koleksi, data.SumberDana, data.Pengelola, data.Tipe, data.Standar, data.TahunBerdiri, data.Bangunan,
				data.LuasTanah, data.StatusKepemilikan)
			err := tulisMalang.Write(mlg)
			if err != nil {
				fmt.Println("Error : ", err)
				return
			}
			tulisMalang.Flush()
		}

		if data.KabupatenKota == "Kota Jakarta Pusat" {
			var j []string
			// j = append(j[:0], "ID Museum", "Kode Pengelolaan", "Nama Museum", "SDM", "Alamat", "Desa/Kelurahan", "Kecamatan", "Kota", "Propinsi", "Lintang", "Bujur", "Koleksi", "Sumber Dana", "Pengelola", "Tipe", "Standar", "Tahun Berdiri", "Bangunan", "Luas Tanah", "Status Kepemilikan")
			j = append(j, data.MuseumID, data.KodePengelolaan, data.Nama, data.Sdm, data.AlamatJalan,
				data.DesaKelurahan, data.Kecamatan, data.KabupatenKota, data.Propinsi, data.Lintang, data.Bujur,
				data.Koleksi, data.SumberDana, data.Pengelola, data.Tipe, data.Standar, data.TahunBerdiri, data.Bangunan,
				data.LuasTanah, data.StatusKepemilikan)
			err := tulisJakpus.Write(j)
			if err != nil {
				fmt.Println("Error : ", err)
				return
			}
			tulisJakpus.Flush()
		}

	}
}
