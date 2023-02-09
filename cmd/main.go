package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	"Client/configs"
	"Client/internal/protos"
	"Client/internal/settings"
)

func main() {
	configs.LoadEnv()
	adr, err := settings.GetEnvDefault("ADDRESS", "")
	if err != nil {
		log.Fatal("no address found")
	}
	conn, err := grpc.Dial(adr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	defer conn.Close()

	client := protos.NewSmallHealthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	comm := make(chan string, 10000)
	for i := 0; i < 10000; i++ {

		go func() {
			resp, err := client.Check(ctx, &emptypb.Empty{})
			comm <- fmt.Sprintf("Postgres:%s, Mongo:%s, Error:%s", resp.Postgres, resp.Mongo, err.Error())

			return
		}()
	}

	go func() {
		var res string
		for i := 0; i < 10000; i++ {
			select {
			case res = <-comm:
				fmt.Println(res)
			default:
				i--
				continue
			}
		}
	}()
	defer cancel()
	time.Sleep(1000000)
}
