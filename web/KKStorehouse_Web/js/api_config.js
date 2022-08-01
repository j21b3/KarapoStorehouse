var url = "http://127.0.0.1:25790"

var api = {
	"raw":{
		//GET    /raw/:pid
		"raw":url+"/raw/",
		//GET http://127.0.0.1:25790/thumbnail?pid=111&width=222
		"thumbnail":url+"/raw/thumbnail"
	},
	//POST   /upload
	"upload":url+"/upload",
	//GET    /timeline/:page
	"timeline":url+"/timeline/"
}

export default api