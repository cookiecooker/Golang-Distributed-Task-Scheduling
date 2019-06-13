package common

// cron job
type Job struct {
	Name string `json:"name"`
	Command string `json:"command"`
	CronExpr string `json:"cronExpr"`
}

//http response
type Response struct {
	Errno int `json:"errno"`
	Msg string `json:"msg"`	
	data interface{} `json:"data"`
}

func BuildResponse(errno int, msg string, data interface{}) (resp []byte, err error){

	var (
		response Response
	)
	response.Errno = errno
	response.Msg = msg
	response.data = data

	resp, err = json.Marshal(response)

	return
}