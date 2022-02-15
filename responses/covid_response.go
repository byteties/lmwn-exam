package responses

type CovidResponse struct {
    AgeGroup    map[string] int  `json:"AgeGroup"`
    Province    map[string] int `json:"Province"`
}