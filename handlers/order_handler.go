package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	. "order-importer/api_connector"
	"order-importer/model/intern"
	"order-importer/order_converter"
)

func NewOrderHandler(
	connector ApiConnector,
	orderConverter order_converter.OrderConverter,
) http.Handler {
	return &orderHandler{
		connector:      connector,
		orderConverter: orderConverter,
	}
}

type orderHandler struct {
	connector      ApiConnector
	orderConverter order_converter.OrderConverter
}

func (o *orderHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		o.getOrders(writer, request)
		break
	case http.MethodPost:
		o.postOrder(writer, request)
		break
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
	writer.Write(body.Payload)

}

func (o *orderHandler) postOrder(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	body, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	var order intern.Order
	err = json.Unmarshal(body, &order)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	externalOrder := o.orderConverter.Convert(order)

	payload, err := json.Marshal(externalOrder)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	//writer.Write(payload)
	response, err := o.connector.Post("/orders", payload, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(response.StatusCode)
	writer.Write(response.Payload)
}
