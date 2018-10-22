package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"strconv"
	"strings"

	"maff/pkg"

	"google.golang.org/grpc"
)

var errBackendErr = errors.New("backend error")

func sendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Panic(err)
	}
}

func sendJSONOk(w http.ResponseWriter, data interface{}) {
	sendJSON(w, http.StatusOK, data)
}

func sendJSONErr(w http.ResponseWriter, status int, err error) {
	sendJSON(w, status, struct {
		Err string `json:"err"`
	}{
		err.Error(),
	})

}

func sendJSONBackendErr(w http.ResponseWriter, err error) {
	sendJSONErr(w, http.StatusInternalServerError, errBackendErr)
	log.Println(err)
}

func getIntVecParam(p string) ([]int64, error) {
	if p == "" {
		return nil, nil
	}

	vals := strings.Split(p, ",")
	itms := make([]int64, 0, len(vals))
	for _, val := range vals {
		v, err := strconv.ParseInt(strings.TrimSpace(val), 10, 64)
		if err != nil {
			return nil, err
		}

		itms = append(itms, v)
	}

	return itms, nil
}

func apiProduct(
	w http.ResponseWriter,
	r *http.Request,
	addr string) {
	vals, err := getIntVecParam(r.FormValue("vals"))
	if err != nil {
		sendJSONErr(w, http.StatusBadRequest, err)
		return
	}

	g, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		sendJSONBackendErr(w, err)
		return
	}
	defer g.Close()

	res, err := pkg.NewMaffClient(g).GetProduct(
		context.Background(),
		&pkg.ProductReq{
			Items: vals,
		})
	if err != nil {
		sendJSONBackendErr(w, err)
		return
	}

	sendJSONOk(w, res)
}

func main() {
	flagAddr := flag.String("addr", ":8080",
		"the bind address")
	flagRPC := flag.String("rpc", "localhost:9090",
		"grpc address")

	flag.Parse()

	log.Panic(http.ListenAndServe(
		*flagAddr,
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				apiProduct(w, r, *flagRPC)
			})))
}
