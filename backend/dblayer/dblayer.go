package dblayer

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"backend.main/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RawPic struct {
	Id         primitive.ObjectID `bson:"_id"`
	CreateTime time.Time          `bson:"create_time"`

	model.PicData
}

type RawPicDBController struct {
	*mongo.Collection
}

func NewRawPicDBController(mongodbURL string) *RawPicDBController {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURL))
	if err != nil {
		panic(err)
	}

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

func (c *RawPicDBController) GenerateID() primitive.ObjectID {
	return primitive.NewObjectID()
}

// 按照时间从近到远按页进行查找 返回图片的ID切片, PageNum从1开始计数 Limit固定为20
// 只要ctx不变化，就可以视作同一事务
func (c *RawPicDBController) GetTimelineID(ctx context.Context, PageNum int64) ([]string, error) {
	findoption := options.FindOptions{}
	Limit := int64(20)
	if PageNum > 0 {
		findoption.SetLimit(Limit)
		findoption.SetSkip(Limit * (PageNum - 1))
	}
	findoption.SetSort(bson.D{{"create_time", -1}})

	cur, err := c.Find(ctx, bson.D{}, &findoption)
	if err != nil {
		return nil, err
	}

	var document []bson.M
	if err := cur.All(ctx, &document); err != nil {
		return nil, err
	}

	ret := []string{}
	for _, each := range document {
		fmt.Printf("%T\n", each["_id"])
		t := each["_id"]
		val := reflect.ValueOf(t)
		f := val.MethodByName("Hex")
		ret = append(ret, f.Call([]reflect.Value{})[0].String())
	}
	return ret, nil
}
