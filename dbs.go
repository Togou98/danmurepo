package dbs

import (
		"go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/mongo/options"
		"go.mongodb.org/mongo-driver/bson"
		"context"
		"time"
		"strconv"
		"fmt"
)

type Pack struct{
		Id int
		Comment string
}

func (p Pack)Getid() int {
		return p.Id
}
func (p Pack)Getcomment() string{
		return p.Comment
}



func Storage(addr string ,ch  chan Pack ,capacity int ,flag chan bool){
		client ,err := mongo.NewClient(options.Client().ApplyURI(addr))
		if err != nil {
				fmt.Println("database error connect")
		}
			ctx , cancle := context.WithTimeout(context.Background(), 20 * time.Second)
			defer cancle()
			err = client.Connect(ctx)
			if err != nil { fmt.Println( "ctx error")}
			var p Pack
		for{
		 <-flag
		for i := 0; i < capacity; i++{
				p = <-ch
				coll := client.Database("Comment").Collection(strconv.Itoa(p.Getid()))
				now := time.Now().Format("2006年-1月-2日 15:04:05")
				comm := string(p.Getcomment())
				_ , err := coll.InsertOne(context.Background(), bson.M{now : comm})
					if err != nil {
							fmt.Println( p.Getid(),"弹幕存储失败")
							continue
								}
					}
		}

	}

