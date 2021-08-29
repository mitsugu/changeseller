package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "io/ioutil"
  "log"
)

func main() {
  var asin            string
  var sellerCode      string
  var store           string
  var price           string
  var tel             string
  var distributorName string
  var resp            string
  var linkUrl         string
  var pageUrl         string

  flag.Parse()
  fname := flag.Arg(0)
  bArr, err := ioutil.ReadFile(fname)
  if err != nil {
      log.Fatal(err)
  }
  var sellers interface{}
  _ = json.Unmarshal(bArr, &sellers)
  fmt.Printf("asin,価格,セラーコード,店舗名,電話番号,販売業者名,運営責任者名,セラーページURL,製品ページURL\n")
  max := len(sellers.([]interface{}))
  for i:=0; i<max; i++ {
    asin            = sellers.([]interface{})[i].(map[string]interface{})["asin"].(string)
    price           = sellers.([]interface{})[i].(map[string]interface{})["price"].(string)
    sellerCode      = sellers.([]interface{})[i].(map[string]interface{})["sellerCode"].(string)
    store           = sellers.([]interface{})[i].(map[string]interface{})["store"].(string)
    tel             = "\""+sellers.([]interface{})[i].(map[string]interface{})["tel"].(string)+"\""
    distributorName = sellers.([]interface{})[i].(map[string]interface{})["distributorName"].(string)
    resp            = sellers.([]interface{})[i].(map[string]interface{})["resp"].(string)
    linkUrl         = sellers.([]interface{})[i].(map[string]interface{})["linkUrl"].(string)
    pageUrl         = sellers.([]interface{})[i].(map[string]interface{})["pageUrl"].(string)
    fmt.Printf("%v,%v,%v,%v,%v,%v,%v,%v,%v\n",asin,price,sellerCode,store,tel,distributorName,resp,linkUrl,pageUrl)
  }
}

