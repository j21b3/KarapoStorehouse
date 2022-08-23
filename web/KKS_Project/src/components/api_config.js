var url = "http://localhost:25790"
var domain = "http://localhost:3000"

var api = {
	"raw":{
		//GET    /raw/:pid
		"raw":url+"/raw/",
		//GET http://127.0.0.1:25790/raw/thumbnail?pid=111&width=222
		"thumbnail":url+"/raw/thumbnail",
		//GET http://127.0.0.1:25790/raw/detail?pid=111
		"detail":url+"/raw/detail"
	},
	//POST   /upload
	"upload":url+"/upload",
	//GET    /timeline/:page
	"timeline":url+"/timeline/",
	"domain":{
		"pic":domain+"/pic/",
	}
}

export default api