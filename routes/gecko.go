package routes

import (
	"encoding/json"
	"generic-http-server/thirdapp"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// PingExample godoc
// @Summary harvest coin price from coingecko
// @Schemes
// @Description harvest coingecko data
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} GetCoinsPrice
// @Router /DeFi/coins/:token [get]
func GetCoinsPrice(c *gin.Context) {
	tokenName := c.Param("token")

	tokenList := []string{tokenName}
	CoinsPrice, err := thirdapp.GetCurrencyValue(tokenList...)
	if err != nil {
		c.String(500, "Unable to get token price")
	}
	CoinsPriceJson, _ := json.Marshal(CoinsPrice)
	c.String(200, string(CoinsPriceJson))
}
