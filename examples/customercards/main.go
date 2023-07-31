package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	gosdk "github.com/team-plain/go-sdk"
	"github.com/team-plain/go-sdk/customercards"
)

func main() {
	http.HandleFunc("/customer-card-endpoint", customerCardsHandler)

	port := 8080
	fmt.Printf("Server listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func customerCardsHandler(w http.ResponseWriter, r *http.Request) {
	var req customercards.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json request", http.StatusBadRequest)
		return
	}
	if len(req.CardKeys) == 0 {
		http.Error(w, "at least one card key is required", http.StatusBadRequest)
	}

	ttl := 86400
	resp := customercards.Response{
		Cards: []*gosdk.CustomerCard{
			getRandomCustomerCard(req.CardKeys[0], ttl),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
		return
	}
}

func getRandomCustomerCard(key string, ttl int) *gosdk.CustomerCard {
	mutedTextColor := gosdk.ComponentTextColorMuted
	return &gosdk.CustomerCard{
		Key:               key,
		TimeToLiveSeconds: &ttl,
		Components: []gosdk.Component{
			{
				ComponentSpacer: &gosdk.ComponentSpacer{
					SpacerSize: gosdk.ComponentSpacerSizeS,
				},
			},
			{
				ComponentRow: &gosdk.ComponentRow{
					RowMainContent: []gosdk.ComponentRowContentUnionInput{
						{
							ComponentText: &gosdk.ComponentText{
								Text:      "Registered at",
								TextColor: &mutedTextColor,
							},
						},
					},
					RowAsideContent: []gosdk.ComponentRowContentUnionInput{
						{
							ComponentText: &gosdk.ComponentText{
								Text: getRandomDateTime(),
							},
						},
					},
				},
			},
			{
				ComponentSpacer: &gosdk.ComponentSpacer{
					SpacerSize: gosdk.ComponentSpacerSizeM,
				},
			},
			{
				ComponentRow: &gosdk.ComponentRow{
					RowMainContent: []gosdk.ComponentRowContentUnionInput{
						{
							ComponentText: &gosdk.ComponentText{
								Text:      "Last signed in",
								TextColor: &mutedTextColor,
							},
						},
					},
					RowAsideContent: []gosdk.ComponentRowContentUnionInput{
						{
							ComponentText: &gosdk.ComponentText{
								Text: getRandomDateTime(),
							},
						},
					},
				},
			},
			{
				ComponentSpacer: &gosdk.ComponentSpacer{
					SpacerSize: gosdk.ComponentSpacerSizeM,
				},
			},
			{
				ComponentRow: &gosdk.ComponentRow{
					RowMainContent: []gosdk.ComponentRowContentUnionInput{
						{
							ComponentText: &gosdk.ComponentText{
								Text:      "Last device used",
								TextColor: &mutedTextColor,
							},
						},
					},
					RowAsideContent: []gosdk.ComponentRowContentUnionInput{
						{
							ComponentText: &gosdk.ComponentText{
								Text: getRandomDeviceText(),
							},
						},
					},
				},
			},
			{
				ComponentSpacer: &gosdk.ComponentSpacer{
					SpacerSize: gosdk.ComponentSpacerSizeM,
				},
			},
			{
				ComponentRow: &gosdk.ComponentRow{
					RowMainContent: []gosdk.ComponentRowContentUnionInput{
						{
							ComponentText: &gosdk.ComponentText{
								Text:      "Marketing preferences",
								TextColor: &mutedTextColor,
							},
						},
					},
					RowAsideContent: []gosdk.ComponentRowContentUnionInput{
						{
							ComponentText: &gosdk.ComponentText{
								Text: getRandomMarketingPreferenceText(),
							},
						},
					},
				},
			},
		},
	}
}

func getRandomDateTime() string {
	now := time.Now()
	randomDuration := time.Duration(rand.Intn(1000)) * time.Minute
	dateTime := now.Add(-randomDuration)
	return dateTime.Format(time.RFC3339)
}

func getRandomDeviceText() string {
	if rand.Float64() > 0.5 {
		return "iPhone 13 🍎"
	}
	return "Galaxy S22  🤖"
}

func getRandomMarketingPreferenceText() string {
	if rand.Float64() > 0.5 {
		return "Opted in 📨"
	}
	return "Opted out 🙅"
}
