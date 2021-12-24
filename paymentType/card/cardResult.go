package card

type ResultPayload struct {
	TypePgt 	string 	`json:"typePgt" bson:"typePgt"`
	Value   	string	`json:"value" bson:"value"`
	Installment string	`json:"installment" bson:"installment"`
	Tid     	string  `json:"tid" bson:"tid"`
	CustomerID	string  `json:"customerID" bson:"customerID"`
	Status 		string	`json:"status" bson:"status"`
}
