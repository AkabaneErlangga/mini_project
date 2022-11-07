package testing

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"main/config"
	"main/controller"
	"main/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func InitDBTest() *echo.Echo {
	config.InitDB()
	e := echo.New()
	return e
}

func TestGetItems(t *testing.T) {
	var testCases = []struct {
		name       string
		data	model.Item
		expectCode int
	}{
		{
			name:       "get user normal",
			data: 		model.Item{ID: 2, Brand: "Colorful", Type: "GTX 1650", LaunchYear: "2019", CategoryId: 1, Stock: 5, Price: 3500000},
			expectCode: http.StatusOK,
		},
	}
	e := InitDBTest()

	req := httptest.NewRequest(http.MethodGet, "/item/2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	//c.SetParamValues("2")
	//c.SetPath("/item")
	for _, testCase := range testCases {
		if assert.NoError(t, controller.GetItem(c)) {
			body := rec.Body.String()

			//var item model.Item
			//err := json.Unmarshal([]byte(body), &item)

			//if err != nil {
			//	assert.Error(t, err, "error")
			//}
			assert.Equal(t, testCase.expectCode, rec.Code)
			assert.Equal(t, testCase.data, body)
		}
	}
}