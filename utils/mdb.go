package utils

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)


type mdb struct {
	Micron map[string]string `json:"micron"`
	Spectek map[string][]string `json:"spectek"`
}

var Mdb mdb;

func LoadMdb() (mdb, error) {
	var mdb mdb
	file, e := os.OpenFile("mdb.json", os.O_RDWR, 0755)
	if e != nil {
		return mdb, e
	}
	data, e := io.ReadAll(file)
	if e != nil {
		return mdb, e
	}
	e = json.Unmarshal(data, &mdb)
	return mdb, e
}

func MFpga2Pn(fpgaCode string) string {
	up := strings.ToUpper(fpgaCode)
	v, ok := Mdb.Micron[up]
	if ok {
		return v
	}
	span := Mdb.Spectek[up]
	if len(span) > 0 {
		return span[0]
	}
	return ""
}