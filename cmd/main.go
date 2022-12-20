package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

func main() {
	cert, err := tls.LoadX509KeyPair("/Users/freeman/Downloads/tfg.cert", "/Users/freeman/Downloads/tfg.key")
	if err != nil {
		panic(err)
	}

	httpClient := http.DefaultClient
	httpClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		},
	}

	// feedClient := feed.NewClientHTTP(httpClient, "https://feed.databet.cloud")
	//
	// cur, err := feedClient.GetLogsFromVersion(context.Background(), "tfg", "1zQpD3apwX9000004gfSEt")
	// if err != nil {
	// 	panic(err)
	// }

	// for cur.HasMore() {
	// 	bb, err := cur.Next(context.Background())
	// 	if err != nil {
	// 		panic(err)
	// 	}
	//
	// 	fmt.Println(string(bb))
	// 	time.Sleep(time.Second)
	// }

	// v, err := feedClient.GetFeedVersion(context.Background(), "tfg")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println(v)

	// mtsClient := mts.NewClientHTTP(httpClient, "https://mts-stage-trading.ginsp.net")
	// resp, err := mtsClient.GetRestrictions(context.Background(), &mts.GetRestrictionsRequest{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(resp)

	// cur, _, err := feedClient.GetAll(context.Background(), "4")
	// if err != nil {
	// 	panic(err)
	// }

	// for cur.HasMore() {
	// 	bb, _ := cur.Next(context.Background())
	// 	fmt.Println(string(bb))
	// }

	// sharedResourceClient := sharedresource.NewClientHTTP(httpClient, "https://api.databet.cloud")
	// market, err := sharedResourceClient.FindMarketByID(context.Background(), 20)
	// if err != nil {
	// 	panic(err)
	// }

	httpReq, _ := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"https://api.databet.cloud/teams/by-ids?ids[]=betting:0:0900_19th_of_january_2018_iem_katowice_2018_asia_closed_qualifier",
		http.NoBody,
	)
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		panic(err)
	}

	defer httpResp.Body.Close()

	bb, _ := io.ReadAll(httpResp.Body)
	fmt.Println(string(bb))
}
