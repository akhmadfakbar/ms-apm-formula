package structs

import "time"

type ApmFormula1 struct {
	Id           	uint      	`json:"id"`
	MetricId        string      `json:"metric_id"`
	MetricName      string      `json:"metric_name"`
	Value_1_BB      float64   	`json:"value_1_bb"`
	Value_1_BA      float64   	`json:"value_1_ba"`
	Value_2_BB      float64    	`json:"value_2_bb"`
	Value_2_BA      float64    	`json:"value_2_ba"`
	Value_3_BB      float64    	`json:"value_3_bb"`
	Value_3_BA      float64    	`json:"value_3_ba"`
	Value_4_BB      float64    	`json:"value_4_bb"`
	Value_4_BA      float64    	`json:"value_4_ba"`
	Value_5_BB      float64    	`json:"value_5_bb"`
	Value_5_BA      float64    	`json:"value_5_ba"`
	Is_range		int			`json:"is_range"`
	CreatedAt       time.Time 	`json:"created_at"`
	CreatedBy       string    	`json:"created_by"`
	ModifiedAt      time.Time 	`json:"modified_at"`
	ModifiedBy      string    	`json:"modified_by"`
	RecordStatus    int       	`json:"rec_status"`
}

type ApmFormula2 struct {
	Id           	uint      	`json:"id"`
	MetricId        string      `json:"metric_id"`
	MetricName      string      `json:"metric_name"`
	Value_1         string      `json:"value_1"`
	Value_2         string      `json:"value_2"`
	Value_3         string      `json:"value_3"`
	Value_4         string      `json:"value_4"`
	Value_5         string      `json:"value_5"`
	Is_range		int			`json:"is_range"`
	CreatedAt       time.Time 	`json:"created_at"`
	CreatedBy       string    	`json:"created_by"`
	ModifiedAt      time.Time 	`json:"modified_at"`
	ModifiedBy      string    	`json:"modified_by"`
	RecordStatus    int       	`json:"rec_status"`
}

type Result struct{
	ApmFormula1  []ApmFormula1 `json:"APM Formula 1"`
	ApmFormula2  []ApmFormula2 `json:"APM Formula 2"`
}

type RequestBody struct{
	MetricId   string  `json:"metric_id" validate:"required"`
	// Is_range   string  `json:"is_range" validate:"required"`
}

type CreateResponse struct{
	Message   string  `json:"message"`
}

type RequestBodyData struct{
	MetricId        string      `json:"metric_id" validate:"required"`
	MetricName      string      `json:"metric_name"`
	Value_1_BB      float64   	`json:"value_1_bb"`
	Value_1_BA      float64   	`json:"value_1_ba"`
	Value_2_BB      float64    	`json:"value_2_bb"`
	Value_2_BA      float64    	`json:"value_2_ba"`
	Value_3_BB      float64    	`json:"value_3_bb"`
	Value_3_BA      float64    	`json:"value_3_ba"`
	Value_4_BB      float64    	`json:"value_4_bb"`
	Value_4_BA      float64    	`json:"value_4_ba"`
	Value_5_BB      float64    	`json:"value_5_bb"`
	Value_5_BA      float64    	`json:"value_5_ba"`
	Value_1         string      `json:"value_1"`
	Value_2         string      `json:"value_2"`
	Value_3         string      `json:"value_3"`
	Value_4         string      `json:"value_4"`
	Value_5         string      `json:"value_5"`
	// Is_range		int			`json:"is_range" validate:"required"`
}