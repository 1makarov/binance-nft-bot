package binance

import (
	"encoding/json"
	u "github.com/1makarov/binance-nft-bot/pkg"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	listURL   = "https://www.binance.com/bapi/nft/v1/public/nft/product-list"
	detailURL = "https://www.binance.com/bapi/nft/v1/friendly/nft/nft-trade/product-detail"
)

type Binance struct {
	Proxy *[]string
	Id    *[]string
}

func Start() {
	bot := Binance{
		Proxy: &[]string{},
		Id:    &[]string{},
	}
	bot.GetProxy()

	for {
		for i := 1; i != 5; i++ {
			b, _ := json.Marshal(List{
				Category:  0,
				Orderby:   "list_time",
				Keyword:   "",
				Ordertype: -1,
				Page:      i,
				Rows:      100,
			})

			statuscode, body, err := u.PostRequest(listURL, b)
			if err != nil || statuscode != fasthttp.StatusOK {
				log.Printf("error request: %d\n", statuscode)
				time.Sleep(5 * time.Second)
				continue
			}

			item, err := handleItems(body)
			if err != nil {
				log.Printf("error json generate\n")
				time.Sleep(5 * time.Second)
				continue
			}

			ids := bot.AddId(item)
			bot.CheckById(ids)
		}
	}
}

func (b *Binance) CheckById(ids []string) {
	var wg sync.WaitGroup

	for i, id := range ids {
		wg.Add(1)
		bytes := handleId(id)
		go func() {
			defer wg.Done()
			for {
				code, body, err := u.PostRequestProxy(detailURL, bytes, (*b.Proxy)[i])
				if code != fasthttp.StatusOK || err != nil {
					log.Printf("error request %d\n", code)
					time.Sleep(3 * time.Second)
					continue
				}

				item, err := handleItem(body)
				if err != nil {
					log.Printf("error getting json %d\n", code)
					time.Sleep(3 * time.Second)
					continue
				}
				price, _ := strconv.ParseFloat(item.Data.Productdetail.Amount, 10)
				if price/float64(item.Data.Productdetail.Batchnum) < 85 {
					log.Println(item.Data.Productdetail.ID, item.Data.Productdetail.Batchnum, item.Data.Productdetail.Amount, item.Data.Productdetail.Currency, price/float64(item.Data.Productdetail.Batchnum))
				}
				break
			}
		}()
	}

	wg.Wait()
}

func (b *Binance) AddId(items *Items) (ids []string) {
	for _, v := range items.Data.Rows {
		if !strings.Contains(v.Title, "NFT Mystery Box Series") {
			continue
		}
		if !u.WW(b.Id, v.Productid) {
			//log.Println(v.Productid, v.Title, v.Amount, v.Currency, v.Status)
			*b.Id = append(*b.Id, v.Productid)
			ids = append(ids, v.Productid)
		}
	}
	return
}
