//package clients
//
//import (
//	"context"
//	"log"
//	"time"
//
//	"github.com/MikhailMishutkin/FoodOrdering/configs"
//	rest "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//	"google.golang.org/protobuf/types/known/timestamppb"
//)
//
//type Client struct{}
//
//func (cl *Client) Сonn() (*rest.GetMenuResponse, error) {
//	log.Println("Conn was invoked")
//
//	conf, err := configs.New("./configs/main.yaml.template")
//	if err != nil {
//		log.Fatalf("can't receive config data: %v\n", err)
//	}
//	conn, err := grpc.Dial(conf.API.GHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
//
//	if err != nil {
//		log.Fatalf("Failed to connect: %v\n", err)
//	}
//
//	defer conn.Close()
//
//	c := rest.NewMenuServiceClient(conn)
//
//	t := time.Now()
//	t1 := t.AddDate(0, 0, 1)
//	ts := timestamppb.New(t1)
//	//fmt.Println(ts)
//	mr := &rest.GetMenuRequest{
//		OnDate: ts,
//	}
//	// fmt.Println(mr)
//
//	res, err := c.GetMenu(context.Background(), mr)
//	if err != nil {
//		log.Println("Can't get the menu from restaurant", err)
//		return nil, err
//	}
//
//	return res, nil
//
//}
//
////type Client struct{}
////
////func (cl *Client) Сonn() (*rest.GetMenuResponse, error) {
////	log.Println("Conn was invoked")
////
////	conf, err := configs.New("./configs/main.yaml.template")
////	if err != nil {
////		log.Fatalf("can't receive config data: %v\n", err)
////	}
////	conn, err := grpc.Dial(conf.API.GHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
////
////	if err != nil {
////		log.Fatalf("Failed to connect: %v\n", err)
////	}
////
////	defer conn.Close()
////
////	t := time.Now()
////	t1 := t.AddDate(0, 0, 1)
////	ts := timestamppb.New(t1)
////	//fmt.Println(ts)
////
////	c := rest.NewMenuServiceClient(conn)
////
////	mr := &rest.GetMenuRequest{
////		OnDate: ts,
////	}
////	// fmt.Println(mr)
////
////	res, err := c.GetMenu(context.Background(), mr)
////	if err != nil {
////		log.Println("Can't get the menu from restaurant", err)
////		return nil, err
////	}
////
////	return res, nil
////
////}
