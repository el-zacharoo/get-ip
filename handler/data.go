package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/el-zacharoo/get-ip/model"
	"github.com/el-zacharoo/get-ip/store"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type Geolocation struct {
	Store *store.Store
}

func (g *Geolocation) Create(w http.ResponseWriter, r *http.Request) {

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

	geo.ID = uuid.New().String()
	geo.Date = time.Now()
	g.Store.AddLocation(geo)
	w.Write([]byte("done"))
}

func (g *Geolocation) Get(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		return
	}

	page := chi.URLParam(r, "page")

	psn, err := g.Store.Getlocation(page)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	rspByt, err := json.Marshal(psn)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	w.Write(rspByt)
}

func (g *Geolocation) Query(w http.ResponseWriter, r *http.Request) {

	cn := r.URL.Query().Get("cn")
	st := r.URL.Query().Get("st")
	lmtStr := r.URL.Query().Get("lmt")
	skipStr := r.URL.Query().Get("off")
	lmt, _ := strconv.ParseInt(lmtStr, 10, 64)
	skip, _ := strconv.ParseInt(skipStr, 10, 64)

	ppl, err := g.Store.Getlocations(cn, st, &lmt, &skip)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	rspByt, err := json.Marshal(ppl)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}
	w.Write(rspByt)
}
