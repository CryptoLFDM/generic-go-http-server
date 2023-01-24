package engine

import (
	"fmt"
	"generic-http-server/config"
	"generic-http-server/utils"
	"net/http"
)

func HarvestCoinPrice() {
	url := fmt.Sprintf("%s:%d/coins/bitcoin", config.Cfg.APIAdress, config.Cfg.APIPort)
	resp, err := http.Get(url)
	utils.HandleHttpError(err)
	defer resp.Body.Close()

	fmt.Println(resp)
}
