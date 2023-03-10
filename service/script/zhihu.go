package script

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ZhiHu() {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl 'https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true' \
	//   -H 'authority: www.zhihu.com' \
	//   -H 'accept: */*' \
	//   -H 'accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7,en-GB;q=0.6' \
	//   -H 'cookie: _zap=18a488c7-6ef9-4c36-8606-35a82e333071; d_c0=ASCXqr_64BWPTos3pCjevcfdMdEziBvoPak=|1668655370; YD00517437729195%3AWM_NI=BVt1BK6gexxyRoWLorwFSQR%2FB0EZ1bdtcU%2FO8%2F4zYu6TFUKAK4kqm041ObiBUtC8xnDhwYFMEo3rhorrX4GUEwOGXuMUwn2L%2BkAiu%2FbAEErCfowPmUMp%2BryhyHXIAGfYelY%3D; YD00517437729195%3AWM_NIKE=9ca17ae2e6ffcda170e2e6ee86cc4490bdfbdab54e888e8ab7d45e879f8fb0d846a19aaad2b64af2e7ba8cd92af0fea7c3b92a828abcd2c141f6b9ab99fb7392b199adee7ab59a9db2d12582f08bbbb46aa8acadb6ae45fcf1fc9bdc7ff68c8ab6ca4b89eca592f03cf8ef8183ec4a97ac8dd9fb6688b8b9bac54787b1fd93eb65b2ae98d5c65f87a981a8db7f8f8bb98fb148e9b58f8af825a29f8ba6f149a6b1b9b1f05f8fb0e593fc64bb9898aac44da687ab8ddc37e2a3; YD00517437729195%3AWM_TID=yBqLrK%2Byp0xBRBABBUPUYdQbluVipr1c; q_c1=f6e4bf4f79ba4f61b3708a2556bc8936|1668660242000|1668660242000; z_c0=2|1:0|10:1671295693|4:z_c0|80:MS4xZ0E0OE1nQUFBQUFtQUFBQVlBSlZUVEgtaW1TYWJ5NndTTWhDUWlobGYwWDR3WnRSYmtoWVl3PT0=|6bbf0791b227349d9d288cb7df17734836e29f8620db636f72a1bd80c8757e74; Hm_lvt_98beee57fd2ef70ccdd5ca52b9740c49=1671621340,1671627042,1671631335,1671675157; _xsrf=6aef426f-1a2b-47ed-8490-bbb20f428390; tst=h; SESSIONID=fMfJhbNusLPPObffaDXGk6TRoluYJfCTw2gnk9mtXRf; JOID=UlgUB0P_0unlQ1pqCvuavi-LpIIemZ6KoyM4JXqEnriwADAMbdWpToJLW24DIJA1f61lahXev7VRHRzXKXoOSEo=; osd=U1wTCkr-1u7oSltuDfaTvyuMqYsfnZmHqiI8IneNn7y3DTkNadKkR4NPXGMKIZQycqRkbhLTtrRVGhHeKH4JRUM=; KLBRSID=fb3eda1aa35a9ed9f88f346a7a3ebe83|1672387517|1672386594' \
	//   -H 'referer: https://www.zhihu.com/hot' \
	//   -H 'sec-ch-ua: "Not?A_Brand";v="8", "Chromium";v="108", "Microsoft Edge";v="108"' \
	//   -H 'sec-ch-ua-mobile: ?0' \
	//   -H 'sec-ch-ua-platform: "Windows"' \
	//   -H 'sec-fetch-dest: empty' \
	//   -H 'sec-fetch-mode: cors' \
	//   -H 'sec-fetch-site: same-origin' \
	//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36' \
	//   -H 'x-ab-param;' \
	//   -H 'x-ab-pb: CpIBCAAbAD8ARwC0AGkBagF0ATsCzALXAtgCoAOhA6IDtwOmBNYEEQVRBYsFjAWeBTAGMQbrBicHdAh2CHkI2gg/CWAJwwnECcUJxgnHCcgJyQnKCcsJzAnRCfQJBApJCmUKawqpCr4K1ArdCu0K/go7CzwLQwtGC3ELhwuNC9cL4AvlC+YLOAxxDI8MrAzDDMkM+AwSSQIHAAAAAAAAAQAAAAAAAAAABAABAAABAAAAAAAABgAAAgAAAAAAAAAAAAAAAwAAAAAAAAEAAAABAAEAAAAFAgEAAAIGAAAAAAA=' \
	//   -H 'x-api-version: 3.0.76' \
	//   -H 'x-requested-with: fetch' \
	//   -H 'x-zse-93: 101_3_3.0' \
	//   -H 'x-zse-96: 2.0_jeqO4/L+S3iXKsc/q1IRS9ZuxATuv/XRT4nQzMzQS2roqegINL5SGBls+UHdRjT4' \
	//   -H 'x-zst-81: 3_2.0aR_sn77yn6O92wOB8hPZnQr0EMYxc4f18wNBUgpTQ6nxERFZfRY0-4Lm-h3_tufIwJS8gcxTgJS_AuPZNcXCTwxI78YxEM20s4PGDwN8gGcYAupMWufIeQuK7AFpS6O1vukyQ_R0rRnsyukMGvxBEqeCiRnxEL2ZZrxmDucmqhPXnXFMTAoTF6RhRuLPFAL9KcH1Q7gLcDSByqxKZUeB20V_eJOOSG2GW9Y_ngLLcipLoBofXCwfXccTveesfDXqnJLBThe9OhwOe9wYPhS1fJc8bwwMhJLYJUwLybULWJHYUbe1l9e0bcUmyCx9hUYftwpfJu3mNqoMrcLGYBpCUrX9c939YvSqswYZkwe_LGoxCqfz5UxGoJg9hDwpybSV8XwGrXx1trXM2BH_K_pBxvV_Q_gOZuw0pLFO2u3MxqSf60SfK7CLbrU13BtY2MHLBvLfYwOL6cnY8wXL5gFCpUXMz92xQ0XOwBgC' \
	//   --compressed

	req, err := http.NewRequest("GET", "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true", nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authority", "www.zhihu.com")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7,en-GB;q=0.6")
	req.Header.Set("Cookie", "_zap=18a488c7-6ef9-4c36-8606-35a82e333071; d_c0=ASCXqr_64BWPTos3pCjevcfdMdEziBvoPak=|1668655370; YD00517437729195%3AWM_NI=BVt1BK6gexxyRoWLorwFSQR%2FB0EZ1bdtcU%2FO8%2F4zYu6TFUKAK4kqm041ObiBUtC8xnDhwYFMEo3rhorrX4GUEwOGXuMUwn2L%2BkAiu%2FbAEErCfowPmUMp%2BryhyHXIAGfYelY%3D; YD00517437729195%3AWM_NIKE=9ca17ae2e6ffcda170e2e6ee86cc4490bdfbdab54e888e8ab7d45e879f8fb0d846a19aaad2b64af2e7ba8cd92af0fea7c3b92a828abcd2c141f6b9ab99fb7392b199adee7ab59a9db2d12582f08bbbb46aa8acadb6ae45fcf1fc9bdc7ff68c8ab6ca4b89eca592f03cf8ef8183ec4a97ac8dd9fb6688b8b9bac54787b1fd93eb65b2ae98d5c65f87a981a8db7f8f8bb98fb148e9b58f8af825a29f8ba6f149a6b1b9b1f05f8fb0e593fc64bb9898aac44da687ab8ddc37e2a3; YD00517437729195%3AWM_TID=yBqLrK%2Byp0xBRBABBUPUYdQbluVipr1c; q_c1=f6e4bf4f79ba4f61b3708a2556bc8936|1668660242000|1668660242000; z_c0=2|1:0|10:1671295693|4:z_c0|80:MS4xZ0E0OE1nQUFBQUFtQUFBQVlBSlZUVEgtaW1TYWJ5NndTTWhDUWlobGYwWDR3WnRSYmtoWVl3PT0=|6bbf0791b227349d9d288cb7df17734836e29f8620db636f72a1bd80c8757e74; Hm_lvt_98beee57fd2ef70ccdd5ca52b9740c49=1671621340,1671627042,1671631335,1671675157; _xsrf=6aef426f-1a2b-47ed-8490-bbb20f428390; tst=h; SESSIONID=fMfJhbNusLPPObffaDXGk6TRoluYJfCTw2gnk9mtXRf; JOID=UlgUB0P_0unlQ1pqCvuavi-LpIIemZ6KoyM4JXqEnriwADAMbdWpToJLW24DIJA1f61lahXev7VRHRzXKXoOSEo=; osd=U1wTCkr-1u7oSltuDfaTvyuMqYsfnZmHqiI8IneNn7y3DTkNadKkR4NPXGMKIZQycqRkbhLTtrRVGhHeKH4JRUM=; KLBRSID=fb3eda1aa35a9ed9f88f346a7a3ebe83|1672387517|1672386594")
	req.Header.Set("Referer", "https://www.zhihu.com/hot")
	req.Header.Set("Sec-Ch-Ua", "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Microsoft Edge\";v=\"108\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	req.Header.Set("X-Ab-Pb", "CpIBCAAbAD8ARwC0AGkBagF0ATsCzALXAtgCoAOhA6IDtwOmBNYEEQVRBYsFjAWeBTAGMQbrBicHdAh2CHkI2gg/CWAJwwnECcUJxgnHCcgJyQnKCcsJzAnRCfQJBApJCmUKawqpCr4K1ArdCu0K/go7CzwLQwtGC3ELhwuNC9cL4AvlC+YLOAxxDI8MrAzDDMkM+AwSSQIHAAAAAAAAAQAAAAAAAAAABAABAAABAAAAAAAABgAAAgAAAAAAAAAAAAAAAwAAAAAAAAEAAAABAAEAAAAFAgEAAAIGAAAAAAA=")
	req.Header.Set("X-Api-Version", "3.0.76")
	req.Header.Set("X-Requested-With", "fetch")
	req.Header.Set("X-Zse-93", "101_3_3.0")
	req.Header.Set("X-Zse-96", "2.0_jeqO4/L+S3iXKsc/q1IRS9ZuxATuv/XRT4nQzMzQS2roqegINL5SGBls+UHdRjT4")
	req.Header.Set("X-Zst-81", "3_2.0aR_sn77yn6O92wOB8hPZnQr0EMYxc4f18wNBUgpTQ6nxERFZfRY0-4Lm-h3_tufIwJS8gcxTgJS_AuPZNcXCTwxI78YxEM20s4PGDwN8gGcYAupMWufIeQuK7AFpS6O1vukyQ_R0rRnsyukMGvxBEqeCiRnxEL2ZZrxmDucmqhPXnXFMTAoTF6RhRuLPFAL9KcH1Q7gLcDSByqxKZUeB20V_eJOOSG2GW9Y_ngLLcipLoBofXCwfXccTveesfDXqnJLBThe9OhwOe9wYPhS1fJc8bwwMhJLYJUwLybULWJHYUbe1l9e0bcUmyCx9hUYftwpfJu3mNqoMrcLGYBpCUrX9c939YvSqswYZkwe_LGoxCqfz5UxGoJg9hDwpybSV8XwGrXx1trXM2BH_K_pBxvV_Q_gOZuw0pLFO2u3MxqSf60SfK7CLbrU13BtY2MHLBvLfYwOL6cnY8wXL5gFCpUXMz92xQ0XOwBgC")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	var autoGenerated AutoGenerated
	bytes, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bytes, &autoGenerated)
	fmt.Println(string(bytes))
}

type AutoGenerated struct {
	Data []struct {
		Type         string `json:"type"`
		StyleType    string `json:"style_type"`
		ID           string `json:"id"`
		CardID       string `json:"card_id"`
		FeedSpecific struct {
			AnswerCount int `json:"answer_count"`
		} `json:"feed_specific"`
		Target struct {
			TitleArea struct {
				Text string `json:"text"`
			} `json:"title_area"`
			ExcerptArea struct {
				Text string `json:"text"`
			} `json:"excerpt_area"`
			ImageArea struct {
				URL string `json:"url"`
			} `json:"image_area"`
			MetricsArea struct {
				Text       string `json:"text"`
				FontColor  string `json:"font_color"`
				Background string `json:"background"`
				Weight     string `json:"weight"`
			} `json:"metrics_area"`
			LabelArea struct {
				Type        string `json:"type"`
				Trend       int    `json:"trend"`
				NightColor  string `json:"night_color"`
				NormalColor string `json:"normal_color"`
			} `json:"label_area"`
			Link struct {
				URL string `json:"url"`
			} `json:"link"`
			TextTagArea struct {
				Text       string `json:"text"`
				FontColor  string `json:"font_color"`
				Background string `json:"background"`
			} `json:"text_tag_area"`
		} `json:"target,omitempty"`
	} `json:"data"`
	Paging struct {
		IsEnd    bool   `json:"is_end"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
	} `json:"paging"`
	FreshText      string `json:"fresh_text"`
	DisplayNum     int    `json:"display_num"`
	FbBillMainRise int    `json:"fb_bill_main_rise"`
}
