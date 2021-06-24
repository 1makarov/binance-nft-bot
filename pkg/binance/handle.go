package binance

import (
	"encoding/json"
	"log"
	"time"
)

type List struct {
	Category  int    `json:"category"`
	Keyword   string `json:"keyword"`
	Orderby   string `json:"orderBy"`
	Ordertype int    `json:"orderType"`
	Page      int    `json:"page"`
	Rows      int    `json:"rows"`
}

type Items struct {
	Code          string      `json:"code"`
	Message       interface{} `json:"message"`
	Messagedetail interface{} `json:"messageDetail"`
	Data          struct {
		Total int `json:"total"`
		Rows  []struct {
			Productid    string      `json:"productId"`
			Title        string      `json:"title"`
			Coverurl     string      `json:"coverUrl"`
			Tradetype    int         `json:"tradeType"`
			Nfttype      int         `json:"nftType"`
			Amount       string      `json:"amount"`
			Currency     string      `json:"currency"`
			Setstarttime int64       `json:"setStartTime"`
			Setendtime   int64       `json:"setEndTime"`
			Timestamp    int64       `json:"timestamp"`
			Rarity       interface{} `json:"rarity"`
			Status       int         `json:"status"`
			Owner        struct {
				Avatarurl string `json:"avatarUrl"`
				Nickname  string `json:"nickName"`
			} `json:"owner"`
		} `json:"rows"`
	} `json:"data"`
	Success bool `json:"success"`
}

type Productid struct {
	Productid string `json:"productId"`
}

func handleItems(b *[]byte) (*Items, error) {
	var items Items
	if err := json.Unmarshal(*b, &items); err != nil {
		log.Println(err)
		time.Sleep(5 * time.Second)
		return nil, err
	}
	return &items, nil
}

func handleId(id string) []byte {
	var p Productid
	p.Productid = id
	b, _ := json.Marshal(p)
	return b
}

type Item struct {
	Code          string      `json:"code"`
	Message       interface{} `json:"message"`
	Messagedetail interface{} `json:"messageDetail"`
	Data          struct {
		Productdetail struct {
			ID            int         `json:"id"`
			Productno     string      `json:"productNo"`
			Title         string      `json:"title"`
			Category      int         `json:"category"`
			Relateid      string      `json:"relateId"`
			Nfttype       int         `json:"nftType"`
			Tradetype     int         `json:"tradeType"`
			Amount        string      `json:"amount"`
			Maxamount     string      `json:"maxAmount"`
			Stepamount    string      `json:"stepAmount"`
			Currentamount string      `json:"currentAmount"`
			Currency      string      `json:"currency"`
			Setstarttime  int64       `json:"setStartTime"`
			Setendtime    int64       `json:"setEndTime"`
			Status        int         `json:"status"`
			Batchnum      int         `json:"batchNum"`
			Stocknum      int         `json:"stockNum"`
			Leftstocknum  int         `json:"leftStockNum"`
			Coverurl      string      `json:"coverUrl"`
			Description   string      `json:"description"`
			Creatorid     interface{} `json:"creatorId"`
			Listerid      interface{} `json:"listerId"`
			Listtime      int64       `json:"listTime"`
			Source        int         `json:"source"`
			Categoryvo    struct {
				Code int    `json:"code"`
				Name string `json:"name"`
			} `json:"categoryVo"`
			Tokenlist []struct {
				Nftid           int    `json:"nftId"`
				Tokenid         string `json:"tokenId"`
				Contractaddress string `json:"contractAddress"`
			} `json:"tokenList"`
		} `json:"productDetail"`
	} `json:"data"`
	Success bool `json:"success"`
}

func handleItem(b *[]byte) (*Item, error) {
	var item Item
	if err := json.Unmarshal(*b, &item); err != nil {
		log.Println(err)
		time.Sleep(5 * time.Second)
		return nil, err
	}
	return &item, nil
}
