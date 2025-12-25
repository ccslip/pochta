package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/term"
)

func init() {
	//projectDir := "C:\\projectgo\\222\\desktop"
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env file not found, using system env variables")
	}
}

func newRequest(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json;charset=UTF-8")
	req.Header.Set("Authorization", "AccessToken "+os.Getenv("ACCESS_TOKEN"))
	req.Header.Set("X-User-Authorization", "Basic "+os.Getenv("PASSWRD"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp, nil
}

func pressAnyKey() {
	fmt.Print("Нажмите любую клавишу, чтобы выйти...")

	if runtime.GOOS == "windows" {
		// На Windows можно использовать простое чтение
		fmt.Scanln()
	} else {
		// На Linux/macOS — читаем один байт без ожидания Enter
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("\n(не удалось перевести терминал в raw-режим, нажмите Enter)")
			fmt.Scanln()
			return
		}
		defer term.Restore(int(os.Stdin.Fd()), oldState)

		b := make([]byte, 1)
		os.Stdin.Read(b)
	}
	fmt.Println()
}

func SendOrder(order Order, url string) {
	var reqorder []Order
	var resporder ErrorResponse
	fullURL := url + "/1.0/user/backlog"
	reqorder = append(reqorder, order)
	jsonData, _ := json.Marshal(reqorder)
	req, err := newRequest(http.MethodPut, fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(body), &resporder)
	if err != nil {
		panic(err)
	}

	fmt.Println(resporder.Errors)
	fmt.Println(resporder.ResultIDs)

}

func NormPhone(rawphone string, url string) []ResponseNormPhone {
	var reqphone []RequestNormPhone
	var phone RequestNormPhone
	var respphone []ResponseNormPhone
	phone.OriginalPhone = rawphone
	reqphone = append(reqphone, phone)
	jsonData, _ := json.Marshal(reqphone)
	fullURL := url + "/1.0/clean/phone"

	req, err := newRequest(http.MethodPost, fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(body), &respphone)
	if err != nil {
		panic(err)
	}
	return respphone
}

func NormAddress(rawaddress string, url string) []ResponseNormAddress {
	var addr RequestNormAddress
	var reqaddr []RequestNormAddress
	var respaddr []ResponseNormAddress
	addr.OriginalAddress = rawaddress
	reqaddr = append(reqaddr, addr)
	jsonData, _ := json.Marshal(reqaddr)
	fullURL := url + "/1.0/clean/address"

	req, err := newRequest(http.MethodPost, fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(body), &respaddr)
	if err != nil {
		panic(err)
	}
	return respaddr
}

func NormFIO(rawfio string, url string) []ResponseNormFIO {
	var fio RequestNormFIO
	var reqfio []RequestNormFIO
	var respfio []ResponseNormFIO
	fio.OriginalFIO = rawfio
	reqfio = append(reqfio, fio)
	jsonData, _ := json.Marshal(reqfio)
	fullURL := url + "/1.0/clean/physical"
	req, err := newRequest(http.MethodPost, fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(body), &respfio)
	if err != nil {
		panic(err)
	}
	return respfio

}

func GetJSON(number int) (int, ResponseScriptSiteTo1C) {
	var scriptsite ResponseScriptSiteTo1C
	var mass float64

	fullURL := "https://cityron.ru/connect20230602/siteTo1c.php?order=" + strconv.Itoa(number)
	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("ОШИБКА СКРИПТА siteTo1c.php")
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	respBody = bytes.TrimPrefix(respBody, []byte("\xef\xbb\xbf"))
	err = json.Unmarshal(respBody, &scriptsite)
	if err != nil {
		fmt.Println("ОШИБКА UNMARSHALL scriptsite")
		log.Fatalf("Ошибка парсинга JSON: %v", err)
	}
	for _, i := range scriptsite.Items {
		mass += float64(i.Weight) * float64(i.Count)
	}
	return int(mass), scriptsite
}

func StrConvToInt(s string) int64 {
	i, e := strconv.Atoi(s)
	if e != nil {
		fmt.Println(e)
	}
	return int64(i)
}

func main() {
	var order Order
	var input string
	fmt.Print("Номер заказа: ")
	fmt.Scanln(&input)
	input1, _ := strconv.Atoi(input)
	mass, parsedorder := GetJSON(int(input1))
	var respnormphone []ResponseNormPhone
	var respnormaddr []ResponseNormAddress

	protokol, _ := os.LookupEnv("PROTOKOL")
	host, _ := os.LookupEnv("HOST")
	URL := protokol + host
	respnormphone = NormPhone(parsedorder.Contacts.Phone, URL)
	respnormaddr = NormAddress(parsedorder.Delivery.Address, URL)
	order.OrderNum = string("15613")
	order.MailDirect = 643
	order.MailCategory = "ORDINARY"
	order.TelAddress = StrConvToInt(respnormphone[0].PhoneCountryCode + respnormphone[0].PhoneCityCode + respnormphone[0].PhoneNumber)
	order.Mass = mass
	order.RegionTo = respnormaddr[0].Region
	order.StreetTo = respnormaddr[0].Street
	order.AreaTo = respnormaddr[0].Area
	order.BuildingTo = respnormaddr[0].Building
	order.PlaceTo = respnormaddr[0].Place
	order.CorpusTo = respnormaddr[0].Corpus
	order.IndexTo = int(StrConvToInt(respnormaddr[0].Index))
	order.HouseTo = respnormaddr[0].House
	order.SlashTo = respnormaddr[0].Slash
	order.RoomTo = respnormaddr[0].Room
	order.LocationTo = respnormaddr[0].Location
	order.LetterTo = respnormaddr[0].Letter
	order.AddressTypeTo = "DEFAULT"
	order.NoReturn = true
	order.MailType = "PARCEL_CLASS_1"
	order.TransportType = "AVIA"
	order.SmsNoticeRecipient = 1
	order.Comment = parsedorder.Contacts.Email
	switch parsedorder.Contacts.ContragentType {
	case "1": //физ лицо
		var respnormfio []ResponseNormFIO
		respnormfio = NormFIO(parsedorder.Contacts.FullName, URL)
		order.RecipientName = respnormfio[0].Surname + " " + respnormfio[0].Name + " " + respnormfio[0].MiddleName
	case "2": //юр лицо
		order.RecipientName = parsedorder.Contacts.ContragentTypeName + " " + parsedorder.Company.Company + " ИНН " + parsedorder.Company.INN
	case "3": //ИП
		order.RecipientName = parsedorder.Contacts.ContragentTypeName + " " + parsedorder.Company.Company + " ИНН " + parsedorder.Company.INN
	}
	SendOrder(order, URL)
	pressAnyKey()
}
