package models

import "encoding/json"

// Stages stages
type Stages struct {

	// Pointer to next page, base64 encoded
	NextPageKey string `json:"nextPageKey,omitempty"`

	// Size of returned page
	PageSize float64 `json:"pageSize,omitempty"`

	// stages
	Stages []*Stage `json:"stages"`

	// Total number of stages
	TotalCount float64 `json:"totalCount,omitempty"`
}

func (s *Stages) ToJSON() ([]byte, error) {
	if s == nil {
		return nil, nil
	}
	return json.Marshal(s)
}

func (s *Stages) FromJSON(b []byte) error {
	var res Stages
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*s = res
	return nil
}
