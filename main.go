package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Product struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Brand        Brand       `json:"brand"`
	Price        Price       `json:"price"`
	URL          string      `json:"url"`
	Images       []string    `json:"images"`
	RatingScore  RatingScore `json:"ratingScore"`
}

type Brand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Price struct {
	DiscountedPrice float64 `json:"discountedPrice"`
	OriginalPrice   float64 `json:"originalPrice"`
	Currency        string  `json:"currency"`
}

type RatingScore struct {
	AverageRating float64 `json:"averageRating"`
	TotalCount    int     `json:"totalCount"`
}

type APIResponse struct {
	Result struct {
		Products        []Product `json:"products"`
		TotalCount      int       `json:"totalCount"`
		Offset          int       `json:"offset"`
		OffsetParameters string   `json:"offsetParameters"`
	} `json:"result"`
}

func main() {
	var productName string
	fmt.Print("Aramak istediğiniz ürün adını girin: ")
	fmt.Scanln(&productName)

	var maxProducts int
	fmt.Print("Maksimum ürün sayısını girin (0 sınırsız): ")
	fmt.Scanln(&maxProducts)

	var allProducts []Product
	offset := 0
	totalProducts := 0

	fmt.Println("\nÜrünler aranıyor...")
	fmt.Println("====================")

	for {
		if maxProducts > 0 && totalProducts >= maxProducts {
			break
		}

		baseURL := "https://apigw.trendyol.com/discovery-web-searchgw-service/v2/api/infinite-scroll/sr"
		queryParams := url.Values{}
		queryParams.Set("q", productName)
		queryParams.Set("qt", productName)
		queryParams.Set("st", productName)
		queryParams.Set("os", "1")
		queryParams.Set("pi", "3")
		queryParams.Set("culture", "tr-TR")
		queryParams.Set("userGenderId", "1")
		queryParams.Set("pId", "103738769")
		queryParams.Set("isLegalRequirementConfirmed", "false")
		queryParams.Set("searchStrategyType", "DEFAULT")
		queryParams.Set("productStampType", "TypeA")
		queryParams.Set("scoringAlgorithmId", "2")
		queryParams.Set("fixSlotProductAdsIncluded", "true")
		queryParams.Set("searchAbDecider", "SearchFETestV1_B,SuggestionLC_B,res_B,BMSA_B,DsNlp_2,AdvertSlotPeriod_2,RRIn_B,SCB_B,SuggestionHighlight_B,BP_B,CatTR_B,SuggestionTermActive_A,AZSmartlisting_62,BH2_B,MB_B,FRA_2,MRF_1,ARR_B,MA_B,SP_B,PastSearches_B,SuggestionJFYProducts_B,SuggestionQF_B,BSA_D,BadgeBoost_A,Relevancy_1,FilterRelevancy_1,Smartlisting_65,SuggestionBadges_B,ProductGroupTopPerformer_B,OpenFilterToggle_2,RR_2,BS_2,SuggestionPopularCTR_B")
		queryParams.Set("location", "null")
		queryParams.Set("initialSearchText", productName)
		queryParams.Set("offset", fmt.Sprintf("%d", offset))
		queryParams.Set("offsetParameters", fmt.Sprintf("Product_%d", offset))
		queryParams.Set("channelId", "1")

		apiURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())

		req, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			fmt.Println("İstek oluşturulurken hata:", err)
			return
		}

		req.Header.Set("accept", "application/json, text/plain, */*")
		req.Header.Set("accept-language", "tr,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
		req.Header.Set("authorization", "Bearer eyJhbGciOiJIUzI1NiIsVCJ9.eyJTdGFuZGFyLCJzdWIiLCJpcoLnRyZW5keW9sLmNvbSIsIm5iZiI6MTcyMDc4MjY2Mn0.MW8UX9a6DJN79yuHVS3p73oI8E8cYfE41lQ8")
		req.Header.Set("origin", "https://www.trendyol.com")
		req.Header.Set("priority", "u=1, i")
		req.Header.Set("sec-ch-ua", "\"Not A(Brand\";v=\"8\", \"Chromium\";v=\"132\", \"Google Chrome\";v=\"132\"")
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
		req.Header.Set("sec-fetch-dest", "empty")
		req.Header.Set("sec-fetch-mode", "cors")
		req.Header.Set("sec-fetch-site", "same-site")
		req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("İstek gönderilirken hata:", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Yanıt okunurken hata:", err)
			return
		}

		var apiResponse APIResponse
		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			fmt.Println("JSON parse edilirken hata:", err)
			return
		}

		allProducts = append(allProducts, apiResponse.Result.Products...)
		totalProducts += len(apiResponse.Result.Products)

		if len(apiResponse.Result.Products) == 0 {
			break
		}

		offset += len(apiResponse.Result.Products)
		fmt.Printf("\rBulunan ürün sayısı: %d", totalProducts)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("\n\nÜrün arama tamamlandı!")

	jsonFile, err := os.Create("all_products.json")
	if err != nil {
		fmt.Println("JSON dosyası oluşturulurken hata:", err)
		return
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(allProducts)
	if err != nil {
		fmt.Println("JSON dosyasına yazılırken hata:", err)
		return
	}

	fmt.Println("Ürünler 'all_products.json' dosyasına kaydedildi.")

	csvFile, err := os.Create("all_products.csv")
	if err != nil {
		fmt.Println("CSV dosyası oluşturulurken hata:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	headers := []string{"ID", "Name", "Brand", "DiscountedPrice", "OriginalPrice", "Currency", "URL", "AverageRating", "TotalRatings"}
	writer.Write(headers)

	for _, product := range allProducts {
		row := []string{
			strconv.Itoa(product.ID),
			product.Name,
			product.Brand.Name,
			fmt.Sprintf("%.2f", product.Price.DiscountedPrice),
			fmt.Sprintf("%.2f", product.Price.OriginalPrice),
			product.Price.Currency,
			product.URL,
			fmt.Sprintf("%.2f", product.RatingScore.AverageRating),
			strconv.Itoa(product.RatingScore.TotalCount),
		}
		writer.Write(row)
	}

	fmt.Println("Ürünler 'all_products.csv' dosyasına kaydedildi.")

	if err := os.Mkdir("imgs", os.ModePerm); err != nil && !os.IsExist(err) {
		fmt.Println("Klasör oluşturulurken hata:", err)
		return
	}

	fmt.Println("\nResimler indiriliyor...")
	for _, product := range allProducts {
		for i, imgURL := range product.Images {
			resp, err := http.Get("https://cdn.dsmcdn.com/mnresize/128/192" + imgURL)
			if err != nil {
				fmt.Printf("Resim indirilirken hata: %s\n", err)
				continue
			}
			defer resp.Body.Close()

			imgPath := filepath.Join("imgs", fmt.Sprintf("%d_%d.jpg", product.ID, i))
			file, err := os.Create(imgPath)
			if err != nil {
				fmt.Printf("Dosya oluşturulurken hata: %s\n", err)
				continue
			}
			defer file.Close()

			_, err = io.Copy(file, resp.Body)
			if err != nil {
				fmt.Printf("Dosyaya yazılırken hata: %s\n", err)
				continue
			}

			fmt.Printf("Resim indirildi: %s\n", imgPath)
		}
	}

	fmt.Println("\nTüm işlemler tamamlandı!")
}
