package http

import (
	"fmt"
	"net/http"
	"sync"
)

var instance *HttpManager
var once sync.Once

type HttpManager struct {
	wg sync.WaitGroup
	addr string
}

func (mgr *HttpManager) GetInstance() *HttpManager {
	once.Do(func() {
		instance = new(HttpManager)
	})
	return instance
}

func (mgr *HttpManager) Init(wg sync.WaitGroup) error {
	mgr.wg = wg
	mgr.addr = ":8080"
	http.HandleFunc("/download",handleDownloadRequest)
	return nil
 }

func (mgr *HttpManager) startServe() {
	mgr.wg.Add(1)
	defer func() {
		mgr.wg.Done()
		fmt.Println("http server return")
	}()
	http.ListenAndServe(mgr.addr,nil)
}

