package model

type PicData struct {
	Title    string   `json:"title"`
	FileName string   `json:"file_name" binding:"required"`
	Data     []byte   `json:"data" binding:"required"`
	Uploader int      `json:"uploader" binding:"required"`
	Message  string   `json:"message"`
	Tags     []string `json:"tags"`
}

type ReturnData struct {
	Data   []byte `json:"data"`
	Status bool   `json:"status"`
	Err    string `json:"err"`
}
