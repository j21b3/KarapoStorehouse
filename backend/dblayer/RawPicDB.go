package dblayer

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RawPic struct {
	Id         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"title"`
	Data       []byte             `bson:"data"`
	Uploader   int                `bson:"uploader"`
	Message    string             `bson:"message"`
	CreateTime time.Time          `bson:"create_time"`
	Tags       []string           `bson:"tags"`
}

type RawPicDBController struct {
	*mongo.Collection
}

func NewRawPicDBController(client *mongo.Client) *RawPicDBController {
	var c RawPicDBController
	c.Collection = client.Database("KarapoStorehouse").Collection("RawPic")
	if c.Collection == nil {
		panic("connect RawPic Error")
	}
	return &c
}

//TODO:校验图像、压缩过程、使用chacha加密不使用md5在客户端完成
func (c *RawPicDBController) InsertPic(ctx context.Context, pic RawPic) error {

	//开启事务
	if err := c.Database().Client().UseSession(ctx, func(session mongo.SessionContext) error {
		if err := session.StartTransaction(); err != nil {
			return err
		}
		defer session.EndSession(ctx)
		if _, err := c.InsertOne(ctx, pic); err != nil {
			if err := session.AbortTransaction(ctx); err != nil {
				fmt.Println("rollback failed")
				return err
			}
			return err
		}
		return session.CommitTransaction(ctx)
	}); err != nil {
		fmt.Println("insert err")
		return err
	}
	return nil
}

func (c *RawPicDBController) InsertPics(ctx context.Context, pics []RawPic) (failList []RawPic) {

	for _, each := range pics {
		if err := c.InsertPic(ctx, each); err != nil {
			failList = append(failList, each)
		}
	}
	if len(failList) == 0 {
		return nil
	} else {
		return failList
	}
}

func (c *RawPicDBController) FindPicById(ctx context.Context, HexIDStr string) (*RawPic, error) {
	objectID, err := primitive.ObjectIDFromHex(HexIDStr)
	if err != nil {
		return nil, err
	}

	var ret RawPic
	if err := c.FindOne(ctx, bson.M{"_id": objectID}).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
