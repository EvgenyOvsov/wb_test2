package server

import (
	"bytes"
	"draft/model"
	"encoding/json"
	"io"
	"net/http"
)

func (srv *Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`OK`))
	w.WriteHeader(http.StatusOK)
}

// Согласно задаче, не знаю почему надо делать одну функцию, но бог с ним
func (srv *Server) handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body := bytes.Buffer{}
	io.Copy(&body, r.Body)

	if r.Method == http.MethodGet {
		obj_type := r.URL.Query().Get(`type`)
		if obj_type == `` {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if obj_type == model.ByerType {
			filter := r.URL.Query().Get(`last_name`)
			if filter == `` {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			buyer := srv.ds.CallGetBuyer(filter)
			if buyer == nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(buyer)

			return
		}
		if obj_type == model.ShopType {
			filter := r.URL.Query().Get(`name`)
			if filter == `` {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			shop := srv.ds.CallGetShop(filter)
			if shop == nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(shop)
			return
		}
	}

	// Если не json - игнорируем
	if r.Header.Get(`Content-Type`) != `application/json` {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	o1 := &model.Buyer{}
	o2 := &model.Shop{}
	second_obj := false

	err := json.Unmarshal(body.Bytes(), o1)
	if err != nil || o1.FirstName == `` || o1.LastName == `` {
		second_obj = true
		err = json.Unmarshal(body.Bytes(), o2)
		if err != nil || o2.Name == `` || o2.Address == `` {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	switch r.Method {

	case http.MethodPost:
		if second_obj {
			err = srv.ds.CallSaveShop(o2)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			err = srv.ds.CallSaveBuyer(o1)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodDelete:
		if second_obj {
			err = srv.ds.CallDeleteShop(o2)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			err = srv.ds.CallDeleteBuyer(o1)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		}
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
