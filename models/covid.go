package models

type CovidData struct {
	ConfirmDate 		string `json:"confirmDate"`
	No 							bool `json:"no"`
	Age							int `json:"Age"`	
	Gender					string `json:"Gender"` 	
	GenderEn 				string  `json:"GenderEn"`
	Nation					bool `json:"Nation"`		
	NationEn				string `json:"NationEn"`
	Province				string `json:"Province"`	 
	ProvinceId			int `json:"ProvinceId"` 		
	District 				bool `json:"District"`		
	ProvinceEn 			string `json:"ProvinceEn"`	
	StatQuarantine	int `json:"StatQuarantine"`	
}

type CovidResponseData struct {
	Data []CovidData `json:"Data"`
}