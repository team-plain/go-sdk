package customercards

import gosdk "github.com/team-plain/go-sdk"

type Request struct {
	CardKeys []string `json:"cardKeys"`
	Customer struct {
		Email      string `json:"email"`
		ExternalID string `json:"externalId"`
	} `json:"customer"`
}

type Response struct {
	Cards []*gosdk.CustomerCard `json:"cards"`
}
