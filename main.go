package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

	//fmt.Print("%+v", m.Data)

	for _, data := range m.Data {
		// if data.KabupatenKota == "Kota Malang" {
		// 	fmt.Printf("Nama Museum: %s \nKabupaten/Kota: %s \n\n", data.Nama, data.KabupatenKota)
		// }
		fmt.Println(data.KabupatenKota)

	}

	fmt.Println()
}
