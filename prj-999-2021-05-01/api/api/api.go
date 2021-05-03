package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/tdarci/prj-999/config"
	"github.com/tdarci/prj-999/engine"
	"github.com/tdarci/prj-999/models"
)

// vvv Public API vvv

// TimeResponse is returned to requests at GET /time
type TimeResponse struct {
	CurrentTime time.Time `json:"current_time,omitempty"`
	Error       string    `json:"error,omitempty"`
}

// MathResponse is returned to requests at GET /add
type MathResponse struct {
	OperandA int    `json:"operand_a"`
	OperandB int    `json:"operand_b"`
	Result   int    `json:"result"`
	Error    string `json:"error,omitempty"`
}

// DogResponse is returned to requests at GET /dog/{id}
type DogResponse struct {
	Dog   *models.Dog `json:"dog,omitempty"`
	Error string      `json:"error,omitempty"`
}

// ^^^ API ^^^
// ---------------------------------------------------

type API struct {
	*config.Config
	engine *engine.Engine
	router *mux.Router
}

func NewAPI(cfg *config.Config) *API {

	a := &API{Config: cfg, engine: engine.NewEngine(cfg)}
	r := mux.NewRouter()
	r.HandleFunc("/time", a.timeHandler).Methods(http.MethodGet)
	r.HandleFunc("/add", a.mathHandler).Methods(http.MethodGet).Queries("a", "{a:[0-9]*}", "b", "{b:[0-9]*}")
	r.HandleFunc("/dog/{id}", a.getDogHandler).Methods(http.MethodGet)

	sph := NewStaticPathHandler(cfg, "web", "index.html")
	r.PathPrefix("/").Handler(sph)

	r.Use(a.logMiddleware)

	a.router = r
	return a
}

func (a *API) Run(apiPort int) error {
	fmt.Printf("starting server on port %d\n", apiPort)

	srv := http.Server{
		Addr:         "127.0.0.1:" + strconv.Itoa(apiPort),
		Handler:      a.router,
		TLSConfig:    nil,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	return srv.ListenAndServe()

	//return http.ListenAndServe(":"+strconv.Itoa(apiPort), a.router)
}

func (a *API) logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.Logger().Println("REQUESTED: ", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (a *API) timeHandler(w http.ResponseWriter, r *http.Request) {
	tr := TimeResponse{
		CurrentTime: time.Now(),
	}
	rspBytes, err := json.Marshal(tr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tr = TimeResponse{Error: err.Error()}
		rspBytes, _ = json.Marshal(tr)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(rspBytes)
}

func (a *API) getDogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if idVal, ok := vars["id"]; ok {
		id, err := strconv.Atoi(idVal)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errRsp := DogResponse{Error: err.Error()}
			errRespBytes, _ := json.Marshal(errRsp)
			w.Write(errRespBytes)
			return
		}

		dog, err := a.engine.GetDog(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errRsp := DogResponse{Error: err.Error()}
			errRespBytes, _ := json.Marshal(errRsp)
			w.Write(errRespBytes)
			return
		}

		rsp := DogResponse{Dog: dog}
		rspBytes, rspErr := json.Marshal(rsp)
		if rspErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errRsp := DogResponse{Error: rspErr.Error()}
			errRespBytes, _ := json.Marshal(errRsp)
			w.Write(errRespBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(rspBytes)
	}
}

func (a *API) mathHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rsp := MathResponse{}
	var oa, ob int
	var err error
	if aVal, ok := vars["a"]; ok {
		oa, err = strconv.Atoi(aVal)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errRsp := MathResponse{Error: err.Error()}
			errRespBytes, _ := json.Marshal(errRsp)
			w.Write(errRespBytes)
			return
		}
		rsp.OperandA = oa
	}
	if bVal, ok := vars["b"]; ok {
		ob, err = strconv.Atoi(bVal)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errRsp := MathResponse{Error: err.Error()}
			errRespBytes, _ := json.Marshal(errRsp)
			w.Write(errRespBytes)
			return
		}
		rsp.OperandB = ob
	}

	rsp = MathResponse{OperandA: oa, OperandB: ob, Result: a.engine.Add(oa, ob)}
	rspBytes, rspErr := json.Marshal(rsp)
	if rspErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errRsp := MathResponse{Error: rspErr.Error()}
		errRespBytes, _ := json.Marshal(errRsp)
		w.Write(errRespBytes)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(rspBytes)

}
