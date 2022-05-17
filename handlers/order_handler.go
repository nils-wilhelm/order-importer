package handlers

import (
	"net/http"
	. "order-importer/api_connector"
)

func NewOrderHandler(
	connector ApiConnector,
) http.Handler {
	return &orderHandler{
		connector: connector,
	}
}

type orderHandler struct {
	connector ApiConnector
}

func (o *orderHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		o.getOrders(writer, request)
	}
	//body, err := io.ReadAll(request.Body)
	//if err != nil {
	//	writer.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//var order model.Order
	//err = json.Unmarshal(body, &order)
	//if err != nil {
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
}

func (o *orderHandler) getOrders(writer http.ResponseWriter, request *http.Request) {
	body, err := o.connector.Get("/orders", nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	writer.Write(body)

}
