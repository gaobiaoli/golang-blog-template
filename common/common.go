package common

import (
	"encoding/json"
	"go-blog/config"
	"go-blog/models"
	"io"
	"log"
	"net/http"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err.Error())
		}
		w.Done()
	}()
	w.Wait()

}

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}

func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, err := json.Marshal(result)
	if err != nil {
		log.Println("json marshal error:", err.Error())
		return

	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}
func Error(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = 500
	result.Error = err.Error()
	resultJson, err := json.Marshal(result)
	if err != nil {
		log.Println("json marshal error:", err.Error())
		return

	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}
