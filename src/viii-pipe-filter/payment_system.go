package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

// Data Models
type Customer struct {
	RefNo     string  `json:"ref_no" xml:"ref_no"`
	Name      string  `json:"name" xml:"name"`
	Address   string  `json:"address" xml:"address"`
	Amount    float64 `json:"amount" xml:"amount"`
	PaymentID string  `json:"payment_id" xml:"payment_id"`
}

type PaymentResult struct {
	Status      string    `json:"status" xml:"status"`
	PaymentID   string    `json:"payment_id" xml:"payment_id"`
	Amount      float64   `json:"amount" xml:"amount"`
	PaymentDate time.Time `json:"payment_date" xml:"payment_date"`
}

// Filter Interface
type Filter interface {
	Process(data interface{}) (interface{}, error)
}

// Pipe Structure
type Pipe struct {
	filters []Filter
}

func NewPipe() *Pipe {
	return &Pipe{filters: make([]Filter, 0)}
}

func (p *Pipe) AddFilter(filter Filter) {
	p.filters = append(p.filters, filter)
}

func (p *Pipe) Execute(data interface{}) (interface{}, error) {
	var err error
	result := data

	for _, filter := range p.filters {
		result, err = filter.Process(result)
		if err != nil {
			return nil, fmt.Errorf("pipe execution error: %v", err)
		}
	}
	return result, nil
}

// Electricity Payment Filters
type ElectricityCustomerLookup struct{}

func (f *ElectricityCustomerLookup) Process(data interface{}) (interface{}, error) {
	refNo := data.(string)
	// Simulate API call to fetch customer details
	customer := Customer{
		RefNo:   refNo,
		Name:    "John Doe",
		Address: "123 Main St",
		Amount:  150.50,
	}
	return customer, nil
}

type ElectricityPaymentProcessor struct{}

func (f *ElectricityPaymentProcessor) Process(data interface{}) (interface{}, error) {
	customer := data.(Customer)
	// Simulate payment API call
	paymentResult := PaymentResult{
		Status:      "SUCCESS",
		PaymentID:   fmt.Sprintf("ELEC-%d", time.Now().Unix()),
		Amount:      customer.Amount,
		PaymentDate: time.Now(),
	}
	return paymentResult, nil
}

// Water Payment Filters
type WaterPaymentXMLTransformer struct{}

func (f *WaterPaymentXMLTransformer) Process(data interface{}) (interface{}, error) {
	customer := data.(Customer)
	
	type WaterPaymentXML struct {
		XMLName     xml.Name  `xml:"WaterPayment"`
		Customer    Customer  `xml:"Customer"`
		PaymentDate time.Time `xml:"PaymentDate"`
	}

	payment := WaterPaymentXML{
		Customer:    customer,
		PaymentDate: time.Now(),
	}

	xmlData, err := xml.MarshalIndent(payment, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("XML transformation error: %v", err)
	}

	return xmlData, nil
}

type WaterPaymentFTPUploader struct{}

func (f *WaterPaymentFTPUploader) Process(data interface{}) (interface{}, error) {
	xmlData := data.([]byte)
	// Simulate FTP upload
	fmt.Printf("Uploading to FTP:\n%s\n", string(xmlData))
	return PaymentResult{
		Status:      "SUCCESS",
		PaymentID:   fmt.Sprintf("WATER-%d", time.Now().Unix()),
		PaymentDate: time.Now(),
	}, nil
}

func main() {
	// Process Electricity Payment
	electricityPipe := NewPipe()
	electricityPipe.AddFilter(&ElectricityCustomerLookup{})
	electricityPipe.AddFilter(&ElectricityPaymentProcessor{})

	// Process Water Payment
	waterPipe := NewPipe()
	waterPipe.AddFilter(&WaterPaymentXMLTransformer{})
	waterPipe.AddFilter(&WaterPaymentFTPUploader{})

	// Example usage for electricity payment
	fmt.Println("Processing Electricity Payment:")
	electricityRefNo := "ELEC123456"
	electricityResult, err := electricityPipe.Execute(electricityRefNo)
	if err != nil {
		fmt.Printf("Electricity payment error: %v\n", err)
		return
	}
	printResult(electricityResult)

	// Example usage for water payment
	fmt.Println("\nProcessing Water Payment:")
	waterCustomer := Customer{
		RefNo:     "WATER789012",
		Name:      "Jane Smith",
		Address:   "456 Water St",
		Amount:    75.25,
		PaymentID: fmt.Sprintf("WATER-%d", time.Now().Unix()),
	}
	waterResult, err := waterPipe.Execute(waterCustomer)
	if err != nil {
		fmt.Printf("Water payment error: %v\n", err)
		return
	}
	printResult(waterResult)
}

func printResult(result interface{}) {
	paymentResult := result.(PaymentResult)
	resultJSON, _ := json.MarshalIndent(paymentResult, "", "  ")
	fmt.Printf("Payment Result:\n%s\n", string(resultJSON))
} 