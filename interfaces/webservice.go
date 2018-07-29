package interfaces

import (
	"github.com/emeve89/glean/usecases"
	"net/http"
	"strconv"
	"io"
	"fmt"
)

type OrderInteractor interface {
	Items(userId, orderId int) ([]usecases.Item, error)
	Add(userId, orderId, itemId int) error
}

type WebServiceHandler struct {
	OrderInteractor OrderInteractor
}

func (handler *WebServiceHandler) ShowOrder(res http.ResponseWriter, req *http.Request) {
	userId, _ := strconv.Atoi(req.FormValue("userId"))
	orderId, _ := strconv.Atoi(req.FormValue("orderId"))
	items, _ := handler.OrderInteractor.Items(userId, orderId)
	for _, item := range items {
		io.WriteString(res, fmt.Sprintf("item id: %d\n", item.Id))
		io.WriteString(res, fmt.Sprintf("item name: %s\n", item.Name))
		io.WriteString(res, fmt.Sprintf("item value: %f\n", item.Value))
	}
}


