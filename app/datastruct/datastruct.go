package datastruct

type PicData struct {
	Title    string   `json:"title" bson:"title"`
	FileName string   `json:"file_name" binding:"required" bson:"file_name"`
	Data     []byte   `json:"data" binding:"required" bson:"data"`
	Uploader int      `json:"uploader" binding:"required" bson:"uploader"`
	Message  string   `json:"message" bson:"message"`
	Tags     []string `json:"tags" bson:"tags"`
}

type ReturnData struct {
	Data   interface{} `json:"data"`
	Status bool        `json:"status"`
	Err    string      `json:"err"`
}
