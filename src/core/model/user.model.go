package model

// type UserEnum_Type int32
// type UserStatus int
// type Idp int

// const (
// 	UserEnum_User        UserEnum_Type = 0
// 	UserEnum_Application UserEnum_Type = 1
// )
// const (
// 	Active UserStatus = iota
// 	Inactive
// 	Deleted
// )
// const (
// 	Auth0 Idp = iota
// 	AzureAd
// )

// func (r UserEnum_Type) String() string {
// 	return [...]string{"User", "Application"}[r]
// }

// func (r UserStatus) String() string {
// 	return [...]string{"Active", "Inactive", "Deleted"}[r]
// }

// func (r Idp) String() string {
// 	return [...]string{"Auth0", "AzureAd"}[r]
// }

type UserModel struct {
	Id             string `json:"id,omitempty" bson:"_id,omitempty"`
	Name           string `json:"name,omitempty" bson:"name,omitempty"`
	Email          string `json:"email,omitempty" bson:"email,omitempty"`
	State          string `json:"state,omitempty" bson:"state,omitempty"`
	Pin            string `json:"pin,omitempty" bson:"pin,omitempty"`
	Country        string `json:"country,omitempty" bson:"country,omitempty"`
	Avatar         string `json:"avatar,omitempty" bson:"avatar,omitempty"`
	UserType       string `json:"userType,omitempty" bson:"userType,omitempty"`
	OrgId          string `json:"orgId,omitempty" bson:"orgId,omitempty"`
	Idp            string `json:"idp,omitempty" bson:"idp,omitempty"`
	IdpUserId      string `json:"idpUserId,omitempty" bson:"idpUserId,omitempty"`
	InternalUserId string `json:"internalUserId,omitempty" bson:"internalUserId,omitempty"`
	Status         string `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt      string `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt      string `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	DeletedAt      string `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}
