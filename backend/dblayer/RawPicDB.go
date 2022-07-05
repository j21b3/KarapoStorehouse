package dblayer

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RawPic struct {
	Id         primitive.ObjectID `bson:"_id"`
	Data       []byte             `bson:"data"`
	Uploader   int                `bson:"uploader"`
	CreateTime time.Time          `bson:"create_time"`
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
func (c *RawPicDBController) InsertOne_test(ctx context.Context, pic []byte, uploaderId int) error {
	id := primitive.NewObjectID()
	rawpic := RawPic{
		Id:         id,
		Data:       pic,
		Uploader:   uploaderId,
		CreateTime: time.Now(),
	}
	//开启事务
	if err := c.Database().Client().UseSession(ctx, func(session mongo.SessionContext) error {
		if err := session.StartTransaction(); err != nil {
			return err
		}
		defer session.EndSession(ctx)
		if _, err := c.InsertOne(ctx, rawpic); err != nil {
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
