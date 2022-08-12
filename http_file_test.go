package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

const (
	uploadHtml = "http_file.html"
	transAddr  = "http://192.168.5.216:18080/audio_trans"
	audiosDir  = "/home/wei.liu/go/src/test/upload-server/uploads"
)

type AudioTransReq struct {
	InFormat   string `json:"in_format"`
	InRate     string `json:"in_rate"`
	OutFormat  string `json:"out_format"`
	OutRate    string `json:"out_rate"`
	InFileName string `json:"in_file_name"`
	InFilePath string `json:"in_file_path"`
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(uploadHtml)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	if r.Method == "POST" {
		inFormat := r.FormValue("in_format")
		inRate := r.FormValue("in_rate")
		outFormat := r.FormValue("out_format")
		outRate := r.FormValue("out_rate")
		f, h, err := r.FormFile("audio_name")
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer func(f multipart.File) {
			err := f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
		}(f)
		filename := h.Filename

		err = os.MkdirAll(audiosDir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t, err := os.Create(audiosDir + "/" + filename)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		req := AudioTransReq{
			InFormat:   inFormat,
			InRate:     inRate,
			OutFormat:  outFormat,
			OutRate:    outRate,
			InFileName: filename,
			InFilePath: audiosDir,
		}
		resp, err := transAudio(&req)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("trans audio resp is:", resp)
		http.Redirect(w, r, "/view/audios", http.StatusFound)
	}
}

func transAudio(req *AudioTransReq) (string, error) {
	bytesData, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", transAddr, reader)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	respByt, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respByt), nil
}

func contentType(fileName string) (contentType string) {
	arr := strings.Split(fileName, ".")
	if len(arr) >= 2 {
		f := arr[len(arr)-1]
		switch f {
		case "mp3":
			contentType = "audio/mp3"
		case "aac":
			contentType = "audio/aac"
		case "adpcm":
			contentType = "audio/adpcm"
		case "amr":
			contentType = "audio/amr"
		case "amrnb":
			contentType = "audio/amrnb"
		case "m4a":
			contentType = "audio/m4a"
		case "ogg":
			contentType = "audio/ogg"
		case "opus":
			contentType = "audio/opus"
		case "hwopus":
			contentType = "audio/hwopus"
		case "pcma":
			contentType = "audio/pcma"
		case "pcmu":
			contentType = "audio/pcmu"
		case "silk":
			contentType = "audio/silk"
		case "speex":
			contentType = "audio/speex"
		case "stdopus":
			contentType = "audio/stdopus"
		case "stdspeex":
			contentType = "audio/stdspeex"
		case "pcm":
			contentType = "audio/pcm"
		case "wav":
			contentType = "audio/wav"
		default:
			contentType = "application/octet-stream"
		}
	}
	return
}

func download(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	esUrl, err := url.QueryUnescape(vars["name"])
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	f, err := os.Open(audiosDir + "/" + esUrl)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	info, err := f.Stat()
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+vars["name"])
	w.Header().Set("Content-Type", contentType(esUrl))
	w.Header().Set("Content-Length", strconv.FormatInt(info.Size(), 10))
	_, err = f.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(w, f)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func viewer() http.Handler {
	return http.StripPrefix("/view/audios", http.FileServer(http.Dir(audiosDir)))
}

func TestHttpFile(t *testing.T) {
	nr := mux.NewRouter()
	nr.HandleFunc("/upload/audios", upload)
	nr.Handle("/view/audios", viewer())
	nr.HandleFunc("/view/{name}", download)
	srv := &http.Server{
		Handler:      nr,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		return
	}
}
