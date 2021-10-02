package handlers

import (
	"github.com/akhmadfakbar/ms-apm-formula/connection"
	"github.com/akhmadfakbar/ms-apm-formula/structs"
	"encoding/json"
	"log"
	"net/http"
	// "time"
	// "fmt"
	"strings"
	"github.com/gorilla/mux"
)

func GetAllApmFormula(w http.ResponseWriter, r *http.Request) {
    apm_formula_1 := []structs.ApmFormula1{}
	apm_formula_2 := []structs.ApmFormula2{}
    result := structs.Result{}

	err := connection.DB.
		Table("apm_formula1").
		Select("*").
	    Scan(&apm_formula_1).Error

	if err != nil {
		log.Println(err)
	}

	err = connection.DB.
		Table("apm_formula2").
		Select("*").
		Scan(&apm_formula_2).Error

    if err != nil {
	   log.Println(err)
    }
	result.ApmFormula1 = apm_formula_1
	result.ApmFormula2 = apm_formula_2
	json.NewEncoder(w).Encode(result)
}

func GetApmFormulaById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	param := mux.Vars(r)
	metric_id := strings.ToLower(param["metric_id"])
	apm_formula_1 := structs.ApmFormula1{}
	apm_formula_2 := structs.ApmFormula2{}

	err := connection.DB.Model(&structs.ApmFormula1{}).Where("metric_id = ?", metric_id).Find(&apm_formula_1).Error
	if err != nil {
		log.Println("Error occured: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} 

	if strings.ToLower(apm_formula_1.MetricId) == metric_id{
		json.NewEncoder(w).Encode(apm_formula_1)
		return
	}

	err = connection.DB.Model(&structs.ApmFormula2{}).Where("metric_id = ?", strings.ToLower(metric_id)).Find(&apm_formula_2).Error
	if err != nil {
		log.Println("Error occured: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if strings.ToLower(apm_formula_2.MetricId) == metric_id{
		json.NewEncoder(w).Encode(apm_formula_2)
		return
	}else {
		json.NewEncoder(w).Encode("Data Tidak Ditemukan")
	}

}

// func StoreApmFormula(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
	
// 	var recordCount int64
// 	connection.DB.Model(&structs.MsKategoriAplikasiOjk{}).Where("record_status = ?", 1).Count(&recordCount)
// 	newKode := recordCount + 1

// 	decoder := json.NewDecoder(r.Body)
// 	payload := structs.MsKategoriAplikasiOjk{
// 		KodeOjk: int(newKode),
// 		CreatedAt: time.Now(),
// 		CreatedBy: "SYSTEM",
// 		ModifiedAt: time.Now(),
// 		ModifiedBy: "SYSTEM",
// 		RecordStatus: 1,
// 	}

// 	if err := decoder.Decode(&payload); err != nil {
// 		log.Println("Error occured: ", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	if err := connection.DB.Create(&payload).Error; err != nil {
// 		log.Println("Error occured: ", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(CreateSuccessResponse{Message: "Kategori OJK has been created successfully!"})
// }

func UpdateApmFormula(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, POST")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	apm_formula_1 := structs.ApmFormula1{}
	apm_formula_2 := structs.ApmFormula2{}

	decoder := json.NewDecoder(r.Body)
	request := structs.RequestBodyData{}

    err := decoder.Decode(&request)
    if err != nil {
        panic(err)
    }

	err = connection.DB.Model(&structs.ApmFormula1{}).Where("metric_id = ?", request.MetricId).Find(&apm_formula_1).Error
	if err != nil {
		log.Println("Error occured: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} 

	if strings.ToLower(apm_formula_1.MetricId) == strings.ToLower(request.MetricId){
		temp := structs.ApmFormula1{}
		temp.MetricId = request.MetricId
		temp.MetricName = request.MetricName
		temp.Value_1_BA = request.Value_1_BA
		temp.Value_1_BB = request.Value_1_BB
		temp.Value_2_BA = request.Value_2_BA
		temp.Value_2_BB = request.Value_2_BB
		temp.Value_3_BA = request.Value_3_BA
		temp.Value_3_BB = request.Value_3_BB
		temp.Value_4_BA = request.Value_4_BA
		temp.Value_4_BB = request.Value_4_BB
		temp.Value_5_BA = request.Value_5_BA
		temp.Value_5_BB = request.Value_5_BB

		if err := connection.DB.Model(&structs.ApmFormula1{}).Where("metric_id = ?", request.MetricId).Updates(temp).Error; err != nil {
			log.Println("Error occured: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(structs.CreateResponse{Message: "Metric Id " + request.MetricId + " has been updated successfully!"})
		return
	}

	err = connection.DB.Model(&structs.ApmFormula2{}).Where("metric_id = ?", request.MetricId).Find(&apm_formula_2).Error
	if err != nil {
		log.Println("Error occured: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if strings.ToLower(apm_formula_2.MetricId) == strings.ToLower(request.MetricId){
		temp := structs.ApmFormula2{}
		temp.MetricId = request.MetricId
		temp.MetricName = request.MetricName
		temp.Value_1 = request.Value_1
		temp.Value_2 = request.Value_2
		temp.Value_3 = request.Value_3
		temp.Value_4 = request.Value_4
		temp.Value_5 = request.Value_5		
		
		if err := connection.DB.Model(&structs.ApmFormula2{}).Where("metric_id = ?", request.MetricId).Updates(temp).Error; err != nil {
			log.Println("Error occured: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(structs.CreateResponse{Message: "Metric Id " + request.MetricId + " has been Updated successfully!"})
		return
	}else {
		json.NewEncoder(w).Encode(structs.CreateResponse{Message: "Metric Id " + request.MetricId + " not found!"})
	}
}

func DestroyApmFormula(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	decoder := json.NewDecoder(r.Body)
	request := structs.RequestBody{}
	apm_formula_1 := structs.ApmFormula1{}
	apm_formula_2 := structs.ApmFormula2{}
    err := decoder.Decode(&request)
    if err != nil {
        panic(err)
    }
	metric_id := strings.ToLower(request.MetricId)

	err = connection.DB.Model(&structs.ApmFormula1{}).Where("metric_id = ?", metric_id).Find(&apm_formula_1).Error
	if err != nil {
		log.Println("Error occured: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} 

	if strings.ToLower(apm_formula_1.MetricId) == metric_id{
		if err := connection.DB.Model(&structs.ApmFormula1{}).Where("metric_id = ?", metric_id).Update("record_status", 0).Error; err != nil {
			log.Println("Error occured: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(structs.CreateResponse{Message: "Metric Id " + metric_id + " has been deleted successfully!"})
		return
	}

	err = connection.DB.Model(&structs.ApmFormula2{}).Where("metric_id = ?", strings.ToLower(metric_id)).Find(&apm_formula_2).Error
	if err != nil {
		log.Println("Error occured: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if strings.ToLower(apm_formula_2.MetricId) == metric_id{
		if err := connection.DB.Model(&structs.ApmFormula2{}).Where("metric_id = ?", metric_id).Update("record_status", 0).Error; err != nil {
			log.Println("Error occured: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(structs.CreateResponse{Message: "Metric Id " + metric_id + " has been deleted successfully!"})
		return
	}else {
		json.NewEncoder(w).Encode(structs.CreateResponse{Message: "Metric Id " + metric_id + " not found!"})
	}
}