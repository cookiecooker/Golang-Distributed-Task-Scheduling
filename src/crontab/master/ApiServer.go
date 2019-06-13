package master

import "net/http"

type ApiServer struct {
	httpServer *http.Server
}

var (
	G_apiServer *ApiServer
)


func handleJobSave(resp http.ResponseWriter, req *http.Request) {
	// save task to etcd

	var (
		err error
		postJob string
		job common.Job
		oldJob *common.Job
		bytes[] byte
	)

	if err = req.ParseForm(); err != nil {
		goto ERR
	}
	postJob = req.PostForm.Get("job")

	if err = json.Unmarshal([]byte(postJob), &job); err != nil {
		goto ERR
	}

	if oldJob, err = G_jobMgr.SaveJob(&job); err != nil {
		goto ERR
	}

	if bytes, err = common.BuildResponse(0, "success", oldJob); err == nil {
		resp.Write(bytes)
	}

	return 

ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(bytes)
	}

}

func InitApiServer() (err error){
	var (
		mux *http.ServerMux
		listener net.Listener
	)

	//config router
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)

	//start TCP listener
	if listener, err = net.Listen("tcp", ":" + strconv.Itoa(G_config.ApiPort)); err != nil {
		return 
	}

	// create a http service
	httpServer = &http.Server {
		ReadTimeout: time.Duration(G_config.ApiReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(G_config.ApiWriteTimeout) * time.Millisecond
		Handler: mux,
	}

	//init singlton
	G_apiServer = &ApiServer{
		httpServer: httpServer,
	}

	// start server
	go httpServer.Serve(listener)

	return
}