package tools

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"time"

	"github.com/disintegration/imaging"
)

// Compress byte data using gzip compression
func CompressPic(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	zw.ModTime = time.Now()
	if _, err := zw.Write(data); err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err := zw.Close(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return buf.Bytes(), nil
}

// UnCompress byte data using gzip compression
func UnCompressPic(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(data)

	rw, err := gzip.NewReader(buf)
	rw.Close()
	if err != nil {
		return nil, err
	}

	var decodebuf bytes.Buffer
	_, err = io.Copy(&decodebuf, rw)
	if err != nil {
		return nil, err
	}

	return decodebuf.Bytes(), nil
}

/**
* 生成图片略缩图
 */
func GenerateThumbnail(data []byte, width, height int) ([]byte, error) {
	buf := bytes.NewBuffer(data)
	image, err := imaging.Decode(buf)
	if err != nil {
		return nil, err
	}
	image = imaging.Thumbnail(image, width, height, imaging.NearestNeighbor)
	/* {
		err = imaging.Save(image, "test.jpg")
		if err != nil {
			return nil, err
		}
	} */
	thumbnailBuf := bytes.Buffer{}
	err = imaging.Encode(&thumbnailBuf, image, imaging.JPEG)
	if err != nil {
		return nil, err
	}

	return thumbnailBuf.Bytes(), nil
}
