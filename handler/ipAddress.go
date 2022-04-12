package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/ip-address/model"
	"github.com/ip-address/store"
)

func create() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// fmt.Println(localAddr.IP)

	return localAddr.IP

}

func Get(w http.ResponseWriter, r *http.Request) {
	var mdl model.IPAddress

	if r.Method == http.MethodOptions {
		return
	}

	mdl.IP = create().String()

	rspByt, err := json.Marshal(mdl)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	w.Write(rspByt)

}

type Geolocation struct {
	Store *store.Store
}

func (g *Geolocation) Add(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		return
	}

	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var geo model.Geolocation
	json.Unmarshal(reqByt, &geo)

	geo.Date = time.Now()
	g.Store.Add(geo)
	w.Write([]byte("done"))
}
