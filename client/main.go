package main

import (
	"context"
	"fmt"
	"log"

	"client/note"
	"time"

	"client/to_json"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer conn.Close()
	c := note.NewNoteOptionClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	
	//添加
	r_Add, err := c.AddRecord(ctx, &note.AddRequest{Title: "123", Text: "456"})
	if err != nil {
		log.Fatalf("can not add record: %v", err)
	}
	res_add := to_json.OutPut(r_Add)
	fmt.Println(res_add)

	//删除
	r_Delete, err := c.DeleteRecord(ctx, &note.Request{Id: 3})
	if err != nil {
		log.Fatalf("can not delete record: %v", err)
	}
	res_del := to_json.OutPut(r_Delete)
	fmt.Println(res_del)

	//更新
	r_Update, err := c.UpdateRecord(ctx, &note.UpdateRequest{Id: 10, Title: "标题", Text: "正文"})
	if err != nil {
		log.Fatalf("can not update record: %v", err)
	}
	res_update := to_json.OutPut(r_Update)
	fmt.Println(res_update)

	//查询
	r_Query, err := c.QueryRecord(ctx, &note.Request{Id: 9})
	if err != nil {
		log.Fatalf("can not search record: %v", err)
	}
	res_Query := to_json.OutPut(r_Query)
	fmt.Println(res_Query)

}
