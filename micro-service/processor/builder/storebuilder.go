package builder

import "processor/model"

type StoreProcess interface {
	WithTitle() model.Store
	WithAddress() model.Store
	WithPinCode() model.Store
	WithPhoneNo() model.Store
	WithTinNo() model.Store
	WithCategory() model.Store

	Build() Store
}

type Store struct {
	ID       string         `json:"id"`
	Title    string         `json:"title"`
	Address  string         `json:"address"`
	PinCode  string         `json:"pincode"`
	PhoneNo  string         `json:"phone_no"`
	Category model.Category `json:"category"`
}
