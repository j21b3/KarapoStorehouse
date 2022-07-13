package datalayer

type PicData struct {
	Title    string   `json:"title"`
	FileName string   `json:"file_name"`
	Data     []byte   `json:"data"`
	Uploader int      `json:"uploader"`
	Message  string   `json:"message"`
	Tags     []string `json:"tags"`
}

type ReturnDara struct {
	Data   []byte `json:"data"`
	Status bool   `json:"status"`
	Err    string `json:"err"`
}
