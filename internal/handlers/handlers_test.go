package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandler_GetAll(t *testing.T) {
	type fields struct {
		Database map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "GetAll",
			fields: fields{
				Database: map[string]interface{}{
					"abc-1": 100,
					"abc-2": 200,
					"xyz-1": 300,
					"xyz-2": 400,
				},
			},
			want:    `{"abc-1":100,"abc-2":200,"xyz-1":300,"xyz-2":400}`,
			wantErr: false,
		},
		{
			name:    "GetAllEmpty",
			fields:  fields{},
			want:    "{}",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Database: tt.fields.Database,
			}
			r := httptest.NewRecorder()
			w := httptest.NewRequest("GET", "/api/v1/", nil)
			got := http.HandlerFunc(h.GetAll)
			got.ServeHTTP(r, w)
			if r.Code != http.StatusOK {
				t.Errorf("Handler.GetAll() = %v, want %v", r.Code, http.StatusOK)
			}
			if r.Header().Get("Content-Type") != "application/json" {
				t.Errorf("Handler.GetAll() = %v, want %v", r.Header().Get("Content-Type"), "application/json")
			}
			if !reflect.DeepEqual(strings.TrimSpace(r.Body.String()), tt.want) != tt.wantErr {
				t.Errorf("Handler.GetAll() = %v, want %v", r.Body.String(), tt.want)
			}
		})
	}
}

func TestHandler_SetValue(t *testing.T) {
	type fields struct {
		Database map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "SetValue",
			fields: fields{
				Database: map[string]interface{}{},
			},
			want:    `{"key":"abc-3","value":300}`,
			wantErr: false,
		},
		{
			name: "SetValueEmpty",
			fields: fields{
				Database: map[string]interface{}{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Database: tt.fields.Database,
			}
			w := httptest.NewRequest("POST", "/api/v1/", strings.NewReader(`{"key":"abc-3","value":300}`))
			r := httptest.NewRecorder()
			got := http.HandlerFunc(h.SetValue)
			got.ServeHTTP(r, w)
			if r.Code != http.StatusOK {
				t.Errorf("Handler.SetValue() = %v, want %v", r.Code, http.StatusOK)
			}
			if r.Header().Get("Content-Type") != "application/json" {
				t.Errorf("Handler.SetValue() = %v, want %v", r.Header().Get("Content-Type"), "application/json")
			}
			if !reflect.DeepEqual(strings.TrimSpace(r.Body.String()), tt.want) != tt.wantErr {
				t.Errorf("Handler.SetValue() = %v, want %v", r.Body.String(), tt.want)
			}
		})
	}
}

func TestHandler_GetValue(t *testing.T) {
	type fields struct {
		Database map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		key     string
		want    string
		wantErr bool
	}{
		{
			name: "GetValue",
			fields: fields{
				Database: map[string]interface{}{
					"abc-1": 100,
					"abc-2": 200,
					"xyz-1": 300,
					"xyz-2": 400,
				},
			},
			key:     "abc-1",
			want:    `100`,
			wantErr: false,
		},
		{
			name: "GetValueEmpty",
			fields: fields{
				Database: map[string]interface{}{},
			},
			want:    "{}",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Database: tt.fields.Database,
			}
			w := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/%s", tt.key), nil)
			r := httptest.NewRecorder()
			m := mux.NewRouter()
			m.HandleFunc("/api/v1/{key}", h.GetValue).Methods("GET")
			m.ServeHTTP(r, w)
			if (r.Code != http.StatusOK) != tt.wantErr {
				t.Errorf("Handler.GetValue() = %v, want %v", r.Code, http.StatusOK)
			}
			if (r.Header().Get("Content-Type") != "application/json") != tt.wantErr {
				t.Errorf("Handler.GetValue() = %v, want %v", r.Header().Get("Content-Type"), "application/json")
			}
			if !reflect.DeepEqual(strings.TrimSpace(r.Body.String()), tt.want) != tt.wantErr {
				t.Errorf("Handler.GetValue() = %v, want %v", r.Body.String(), tt.want)
			}
		})
	}
}
