package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	greetv1 "connect-tutorial/src/adapter/grpc/greet/v1"
	greetv1connect "connect-tutorial/src/adapter/grpc/greet/v1/greetv1connect"
	hogev1 "connect-tutorial/src/adapter/grpc/hoge/v1"
	hogev1connect "connect-tutorial/src/adapter/grpc/hoge/v1/hogev1connect"
)

type GreetServer struct{}
type HogeServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func (s *HogeServer) GetHoge(
	ctx context.Context,
	req *connect.Request[hogev1.HogeGetRequest],
) (*connect.Response[hogev1.HogeGetResponse], error) {
	log.Println("Request headers: ", req.Header())
	log.Println("Request body: ", req.Msg.Id)
	res := connect.NewResponse(&hogev1.HogeGetResponse{
		Hoge: &hogev1.Hoge{
			Fuga: &hogev1.Fuga{
				Piyo: fmt.Sprintf("Hoge, %s!", req.Msg.Id),
			},
		},
	})
	res.Header().Set("Hoge-Get-Version", "v1")
	return res, nil
}

func (s *HogeServer) CreateHoge(
	ctx context.Context,
	req *connect.Request[hogev1.HogeCreateRequest],
) (*connect.Response[hogev1.HogeCreateResponse], error) {
	log.Println("Request headers: ", req.Header())
	log.Println("Request body: ", req.Msg.Hoge)
	res := connect.NewResponse(&hogev1.HogeCreateResponse{
		Hoge: &hogev1.Hoge{
			Fuga: &hogev1.Fuga{
				Piyo: "piyopiyo",
			},
		},
	})
	res.Header().Set("Hoge-Create-Version", "v1")
	return res, nil
}

func main() {
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(&GreetServer{})
	path2, handler2 := hogev1connect.NewHogeServiceHandler(&HogeServer{})
	mux.Handle(path, handler)
	mux.Handle(path2, handler2)

	// Use h2c so we can serve HTTP/2 without TLS.
	corsHandler := cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{}))
	http.ListenAndServe(
		"localhost:8080",
		corsHandler,
	)
}
