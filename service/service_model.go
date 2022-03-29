package service

type StandardRequest struct {
	CorrelatedId string `json:"correlatedId"`
}

type StandardResponse struct {
	UUID      string `json:"uuid"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

type RequestPrepaidCatalogue struct {
	CorrelatedId   string                 `json:"correlatedId"`
	CallerUuid     string                 `json:"calleruuid"`
	Data           map[string]interface{} `json:"data"`
	KafkaTimestamp int64                  `json:kafkaTimestamp`
}

type ResponsePrepaidCatalogue struct {
	UUID      string `json:"uuid"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

type RequestUpsertPrepaidCatalogue RequestPrepaidCatalogue

type ResponseUpsertPrepaidCatalogue ResponsePrepaidCatalogue

type RequestUpdatePrepaidCatalogue RequestPrepaidCatalogue

type ResponseUpdatePrepaidCatalogue ResponsePrepaidCatalogue

type RequestRemovePrepaidCatalogue RequestPrepaidCatalogue

type ResponseRemovePrepaidCatalogue ResponsePrepaidCatalogue

type PrepaidCatalogueRequest struct {
	ID                *string             `json:"id" validate:"required"`
	Code              *string             `json:"code" validate:"required"`
	Name              *string             `json:"name" validate:"required"`
	Description       *string             `json:"description"`
	SaleEffDateStr    *string             `json:"saleEffDate" `
	SaleExpDateStr    *string             `json:"saleExpDate" `
	Type              *string             `json:"type"`
	TypeDesc          *string             `json:"typeDesc"`
	ProductType       *string             `json:"productType"`
	SaleContext       *string             `json:"saleContext"`
	UrNo              *string             `json:"urNo"`
	VersionEffDateStr *string             `json:"versionEffDate"`
	VersionExpDateStr *string             `json:"versionExpDate"`
	IsActive          *bool               `json:"isActive"`
	Properties        map[string]string   `json:"properties"`
	OfferItem         []PrepaidOfferItem  `json:"offerItem"`
	ChildOffer        []PrepaidChildOffer `json:"childOffer"`
}

type PrepaidOfferItem struct {
	ID                  *string           `json:"id"`
	Name                *string           `json:"name"`
	Code                *string           `json:"code"`
	Description         *string           `json:"description"`
	VersionEffDateStr   *string           `json:"versionEffDate"`
	VersionExpDateStr   *string           `json:"versionExpDate"`
	OfferItemProperties map[string]string `json:"offerItemProperties"`
	OfferItemParam      map[string]string `json:"offerItemParam"`
}

type PrepaidChildOffer struct {
	ChildOfferId      *string            `json:"childOfferId"`
	RelationType      *string            `json:"relationType"`
	SelectedByDefault *bool              `json:"selectedByDefault"`
	Name              *string            `json:"name"`
	Description       *string            `json:"description"`
	SaleEffDateStr    *string            `json:"saleEffDate"`
	SaleExpDateStr    *string            `json:"saleExpDate"`
	Type              *string            `json:"type"`
	TypeDesc          *string            `json:"typeDesc"`
	ProductType       *string            `json:"productType"`
	SaleContext       *string            `json:"saleContext"`
	UrNo              *string            `json:"urNo"`
	VersionEffDateStr *string            `json:"versionEffDate"`
	VersionExpDateStr *string            `json:"versionExpDate"`
	IsActive          *bool              `json:"isActive"`
	Properties        map[string]string  `json:"properties"`
	OfferItem         []PrepaidOfferItem `json:"offerItem"`
}
