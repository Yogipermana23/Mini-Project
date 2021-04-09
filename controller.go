package main

import (
	"fmt"
	"math"
	"strconv"

	_ "github.com/apache/calcite-avatica-go"
)

//Churn
func (param RequestData) gethandlerChurn() ([]ChurnResponse, int64, int64, error) {
	var chr []ChurnResponse = []ChurnResponse{}
	// var crh []ChurnResponse = []ChurnResponse{}
	//Create Connection with Phoenix
	var db, err = connect()

	var total_page int64
	var total_row int64

	if err != nil {
		fmt.Println(err.Error())
		return chr, total_page, total_row, err
	}

	page, err_page := strconv.Atoi(param.Page)
	offset := (page - 1) * 10

	if param.Branch == "" && (err_page != nil || page < 1) {
		return chr, total_page, total_row, err
	}

	//Get Data From Phoenix Table
	if param.Region == "" {

		err = db.
			Table("CUSTOMER360_SUMMARY_CHURN").
			Select("churn_level, ds, sum(total_churn) as total_churn").
			Limit(10).
			Offset(offset).
			Group("churn_level, ds").
			Scan(&chr).Error

		if err != nil {
			fmt.Println(err.Error())
			return chr, total_page, total_row, err
		}

		err = db.
			Table("CUSTOMER360_SUMMARY_CHURN").
			Select("churn_level, ds, sum(total_churn) as total_churn").
			Group("churn_level, ds").
			Count(&total_row).Error

		total_page = int64(math.Ceil(float64(total_row) / 10))

		if err != nil {
			fmt.Println(err.Error())
			return chr, total_page, total_row, err
		}

		fmt.Println("count page : ", total_page)
		fmt.Println("count row :", total_row)

		// return chr, total_page, total_row, nil
	} else {
		err = db.
			Table("CUSTOMER360_SUMMARY_CHURN").
			Select("churn_level, ds, region, mainbr, branch, sum(total_churn) as total_churn").
			Where(generateWhereClause(param)).
			Limit(10).
			Offset(offset).
			Group("churn_level, ds, region, mainbr, branch").
			Scan(&chr).Error

		if err != nil {
			fmt.Println(err.Error())
			return chr, total_page, total_row, err
		}

		err = db.
			Table("CUSTOMER360_SUMMARY_CHURN").
			Where(startdateWhereClause(param)).
			Count(&total_row).Error

		total_page = int64(math.Ceil(float64(total_row) / 10))

		if err != nil {
			fmt.Println(err.Error())
			return chr, total_page, total_row, err
		}

	}
	return chr, total_page, total_row, nil

}

func generateWhereClause(param RequestData) string {
	var whereClause string

	if param.Region != "" {
		whereClause = "REGION ='" + param.Region + "'"
		if param.Mainbr != "" {
			whereClause += " AND MAINBR ='" + param.Mainbr + "'"
			if param.Branch != "" {
				whereClause += "AND BRANCH = '" + param.Mainbr + "'"
			} else {
				whereClause = "REGION ='" + param.Region + "'"

			}
		}

	}
	return whereClause
}

//Fasilitas Rekening

func (param RequestData) getFasilitasRekening() (statusFasilitas, int64, string, error) {
	var fasilitas_rekening []FasilitasRekeningResponse = []FasilitasRekeningResponse{}
	var fasrek statusFasilitas = statusFasilitas{}
	var status aktif_tidakaktif = aktif_tidakaktif{}
	// var temp statusFasilitas = statusFasilitas{}
	//Create Connection with Phoenix
	var db, err = connect()

	var period_date int64
	var modified_date string

	if err != nil {
		fmt.Println(err.Error())
		return fasrek, period_date, modified_date, err
	}

	err = db.
		Table("BRI360_SUMMARY_FASREK").
		Select("SUM(REG_SMS) as REG_SMS, SUM(SMS_FIN) as SMS_FIN, SUM(REG_INTERNET) as REG_INTERNET, SUM(INTERNET_FIN) as INTERNET_FIN, SUM(REG_PHONE) as REG_PHONE, SUM(BRIVA_AUTOPAYMENT) as BRIVA_AUTOPAYMENT, SUM(DPLK_AUTOPAYMENT) as DPLK_AUTOPAYMENT, SUM(FIF_AUTOPAYMENT) as FIF_AUTOPAYMENT, SUM(PLN_AUTOPAYMENT) as PLN_AUTOPAYMENT, SUM(TELKOM_AUTOPAYMENT) as TELKOM_AUTOPAYMENT, SUM(NOTIF_SMS) as NOTIF_SMS, SUM(NOTIF_EMAIL) as NOTIF_EMAIL, STATUS").
		Where(startdateWhereClause(param)).
		Group("STATUS").
		Scan(&fasilitas_rekening).Error

	for _, element := range fasilitas_rekening {
		if element.STATUS == "AKTIF" {
			status.AKTIF = element.BRIVA_AUTOPAYMENT
			status.TIDAK_AKTIF = fasilitas_rekening[1].BRIVA_AUTOPAYMENT
			fasrek.BRIVA_AUTOPAYMENT = status

			status.AKTIF = element.DPLK_AUTOPAYMENT
			status.TIDAK_AKTIF = fasilitas_rekening[1].DPLK_AUTOPAYMENT
			fasrek.DPLK_AUTOPAYMENT = status

			status.AKTIF = element.FIF_AUTOPAYMENT
			status.TIDAK_AKTIF = fasilitas_rekening[1].FIF_AUTOPAYMENT
			fasrek.FIF_AUTOPAYMENT = status

			status.AKTIF = element.INTERNET_FIN
			status.TIDAK_AKTIF = fasilitas_rekening[1].INTERNET_FIN
			fasrek.INTERNET_FIN = status

			status.AKTIF = element.NOTIF_EMAIL
			status.TIDAK_AKTIF = fasilitas_rekening[1].NOTIF_EMAIL
			fasrek.NOTIF_EMAIL = status

			status.AKTIF = element.NOTIF_SMS
			status.TIDAK_AKTIF = fasilitas_rekening[1].NOTIF_SMS
			fasrek.NOTIF_SMS = status

			status.AKTIF = element.PLN_AUTOPAYMENT
			status.TIDAK_AKTIF = fasilitas_rekening[1].PLN_AUTOPAYMENT
			fasrek.PLN_AUTOPAYMENT = status

			status.AKTIF = element.REG_INTERNET
			status.TIDAK_AKTIF = fasilitas_rekening[1].REG_INTERNET
			fasrek.REG_INTERNET = status

			status.AKTIF = element.REG_PHONE
			status.TIDAK_AKTIF = fasilitas_rekening[1].REG_PHONE
			fasrek.REG_PHONE = status

			status.AKTIF = element.REG_SMS
			status.TIDAK_AKTIF = fasilitas_rekening[1].REG_SMS
			fasrek.REG_SMS = status

			status.AKTIF = element.SMS_FIN
			status.TIDAK_AKTIF = fasilitas_rekening[1].SMS_FIN
			fasrek.SMS_FIN = status

			status.AKTIF = element.TELKOM_AUTOPAYMENT
			status.TIDAK_AKTIF = fasilitas_rekening[1].TELKOM_AUTOPAYMENT
			fasrek.TELKOM_AUTOPAYMENT = status
		} else {
			status.TIDAK_AKTIF = element.BRIVA_AUTOPAYMENT
			status.AKTIF = fasilitas_rekening[1].BRIVA_AUTOPAYMENT
			fasrek.BRIVA_AUTOPAYMENT = status

			status.TIDAK_AKTIF = element.DPLK_AUTOPAYMENT
			status.AKTIF = fasilitas_rekening[1].DPLK_AUTOPAYMENT
			fasrek.DPLK_AUTOPAYMENT = status

			status.TIDAK_AKTIF = element.FIF_AUTOPAYMENT
			status.AKTIF = fasilitas_rekening[1].FIF_AUTOPAYMENT
			fasrek.FIF_AUTOPAYMENT = status

			status.TIDAK_AKTIF = element.INTERNET_FIN
			status.AKTIF = fasilitas_rekening[1].INTERNET_FIN
			fasrek.INTERNET_FIN = status

			status.TIDAK_AKTIF = element.NOTIF_EMAIL
			status.AKTIF = fasilitas_rekening[1].NOTIF_EMAIL
			fasrek.NOTIF_EMAIL = status

			status.TIDAK_AKTIF = element.NOTIF_SMS
			status.AKTIF = fasilitas_rekening[1].NOTIF_SMS
			fasrek.NOTIF_SMS = status

			status.TIDAK_AKTIF = element.PLN_AUTOPAYMENT
			status.AKTIF = fasilitas_rekening[1].PLN_AUTOPAYMENT
			fasrek.PLN_AUTOPAYMENT = status

			status.TIDAK_AKTIF = element.REG_INTERNET
			status.AKTIF = fasilitas_rekening[1].REG_INTERNET
			fasrek.REG_INTERNET = status

			status.TIDAK_AKTIF = element.REG_PHONE
			status.AKTIF = fasilitas_rekening[1].REG_PHONE
			fasrek.REG_PHONE = status

			status.TIDAK_AKTIF = element.REG_SMS
			status.AKTIF = fasilitas_rekening[1].REG_SMS
			fasrek.REG_SMS = status

			status.TIDAK_AKTIF = element.SMS_FIN
			status.AKTIF = fasilitas_rekening[1].SMS_FIN
			fasrek.SMS_FIN = status

			status.TIDAK_AKTIF = element.TELKOM_AUTOPAYMENT
			status.AKTIF = fasilitas_rekening[1].TELKOM_AUTOPAYMENT
			fasrek.TELKOM_AUTOPAYMENT = status

		}
		break
	}

	fmt.Println(fasilitas_rekening)
	if err != nil {
		fmt.Println(err.Error())
		return fasrek, period_date, modified_date, err
	}

	if len(fasilitas_rekening) > 0 {
		row := db.
			Table("BRI360_SUMMARY_FASREK").
			Select("MAX(MODIFIED_DATE) as modified_date").
			Where(startdateWhereClause(param)).
			Row()

		row.Scan(&modified_date)

		fmt.Println(modified_date)
		fmt.Println(period_date)

		if err != nil {
			fmt.Println(err.Error())
			return fasrek, period_date, modified_date, err
		}
	}

	return fasrek, period_date, modified_date, nil
}

//start&end date wheere clause
func startdateWhereClause(param RequestData) string {
	var whereClause string
	if param.Start_date != "" {
		if param.End_date != "" {
			whereClause = "DS BETWEEN '" + param.Start_date + "' AND '" + param.End_date + "'"
		} else {
			whereClause = "DS='" + param.Start_date + "'"
		}
		if param.Region != "" {
			whereClause += " AND REGION ='" + param.Region + "'"
			if param.Mainbr != "" {
				whereClause += " AND MAINBR =" + param.Mainbr + ""
				if param.Branch != "" {
					whereClause += "AND BRANCH = " + param.Branch + ""
				}
			}
		}
	} else {
		whereClause = "DS ='" + param.Start_date + "'"

	}

	return whereClause
}

//Segmentasi
func (param RequestData) getSegmentasi() ([]SegmentasiResponse, [2]string, string, error) {
	var segmentasi []SegmentasiResponse = []SegmentasiResponse{}

	//Create Connection with Phoenix
	var db, err = connect()

	var period_date [2]string
	var min_date string
	var max_date string
	var modified_date string

	if err != nil {
		fmt.Println(err.Error())
		return segmentasi, period_date, modified_date, err
	}

	err = db.
		Table("BRI360_SUMMARY_SEGMENT").
		Where(startdateWhereClause(param)).
		Scan(&segmentasi).Error

	if err != nil {
		fmt.Println(err.Error())
		return segmentasi, period_date, modified_date, err
	}

	row := db.
		Table("BRI360_SUMMARY_SEGMENT").
		Select("MAX(MODIFIED_DATE) as modified_date").
		Where(startdateWhereClause(param)).
		Row()

	row.Scan(&modified_date)

	fmt.Println(modified_date)

	if err != nil {
		fmt.Println(err.Error())
		return segmentasi, period_date, modified_date, err
	}

	row = db.
		Table("BRI360_SUMMARY_SEGMENT").
		Select("MIN(DS) AS min_date, MAX(DS) as max_date").
		Where(startdateWhereClause(param)).
		Row()

	row.Scan(&min_date, &max_date)

	period_date[0] = min_date
	period_date[1] = max_date

	fmt.Println(modified_date)

	if err != nil {
		fmt.Println(err.Error())
		return segmentasi, period_date, modified_date, err
	}

	return segmentasi, period_date, modified_date, nil
}

//Ratas saldo
func (param RequestData) getRatassaldo() (RatassaldoResponse, [2]string, string, error) {
	var ratassaldo RatassaldoResponse = RatassaldoResponse{}

	//Create Connection with Phoenix
	var db, err = connect()

	var period_date [2]string
	var min_date string
	var max_date string
	var modified_date string

	if err != nil {
		fmt.Println(err.Error())
		return ratassaldo, period_date, modified_date, err
	}

	err = db.
		Table("BRI360_SUMMARY_RATAS").
		Where(startdateWhereClause(param)).
		Scan(&ratassaldo).Error

	if err != nil {
		fmt.Println(err.Error())
		return ratassaldo, period_date, modified_date, err
	}

	row := db.
		Table("BRI360_SUMMARY_RATAS").
		Select("MAX(MODIFIED_DATE) as modified_date").
		Where(startdateWhereClause(param)).
		Row()

	row.Scan(&modified_date)

	fmt.Println(modified_date)

	if err != nil {
		fmt.Println(err.Error())
		return ratassaldo, period_date, modified_date, err
	}

	row = db.
		Table("BRI360_SUMMARY_RATAS").
		Select("MIN(DS) AS min_date, MAX(DS) as max_date").
		Where(startdateWhereClause(param)).
		Row()

	row.Scan(&min_date, &max_date)

	period_date[0] = min_date
	period_date[1] = max_date

	fmt.Println(modified_date)

	if err != nil {
		fmt.Println(err.Error())
		return ratassaldo, period_date, modified_date, err
	}

	return ratassaldo, period_date, modified_date, nil
}

//Top Frekuensi Merchant
func (param RequestData) getTopFrekuensiMerchant() (interface{}, [2]string, string, error) {
	var top_frekuensi_merchant []TopFrekuensiMerchantResponse = []TopFrekuensiMerchantResponse{}
	var data map[string]int = map[string]int{}
	// var top_frekuensi_summary []topfreqsummary

	//Create Connection with Phoenix
	var db, err = connect()
	var period_date [2]string
	var min_date string
	var max_date string
	var modified_date string

	if err != nil {
		fmt.Println(err.Error())
		return data, period_date, modified_date, err
	}

	//Get Data From Phoenix Table
	err = db.
		Table("BRI360_SUMMARY_SPENDING_DEBIT").
		Select("DS, MERCHANT_GROUP, sum(FREQ) AS TOP_FREQ").
		Where(startdateWhereClause(param)).
		Group("DS, MERCHANT_GROUP").
		Order("TOP_FREQ DESC").
		Limit(10).
		Scan(&top_frekuensi_merchant).Error

	if err != nil {
		fmt.Println(err.Error())
		return data, period_date, modified_date, err
	}

	for _, element := range top_frekuensi_merchant {
		top_freq, _ := strconv.Atoi(element.TOP_FREQ)
		data[element.MERCHANT_GROUP] = top_freq
	}

	row := db.
		Table("BRI360_SUMMARY_SPENDING_DEBIT").
		Select("MAX(MODIFIED_DATE) as modified_date").
		Where(startdateWhereClause(param)).
		Row()

	row.Scan(&modified_date)
	fmt.Println(modified_date)

	if err != nil {
		fmt.Println(err.Error())
		return data, period_date, modified_date, err
	}
	row = db.
		Table("BRI360_SUMMARY_SPENDING_DEBIT").
		Select("MIN(DS) AS min_date, MAX(DS) as max_date").
		Where(startdateWhereClause(param)).
		Row()

	row.Scan(&min_date, &max_date)

	period_date[0] = min_date
	period_date[1] = max_date

	fmt.Printf(modified_date)

	if err != nil {
		fmt.Println(err.Error())
		return data, period_date, modified_date, err
	}

	// return top_frekuensi_merchant, period_date, modified_date, nil
	return data, period_date, modified_date, nil
}

//Top Nominal Merchant
func (param RequestData) getTopNominalFrekuensiMerchant() (interface{}, [2]string, string, error) {
	var top_nominal_frekuensi_merchant []TopNominalFrekuensiMerchantResponse = []TopNominalFrekuensiMerchantResponse{}
	var data map[string]int = map[string]int{}
	//Create Connection with Phoenix
	var db, err = connect()
	var period_date [2]string
	var min_date string
	var max_date string
	var modified_date string

	if err != nil {
		fmt.Println(err.Error())
		return data, period_date, modified_date, err
	}

	//Get Data From Phoenix Table
	err = db.
		Table("BRI360_SUMMARY_SPENDING_DEBIT").
		Select("DS, MERCHANT_GROUP,  sum(AMT) AS TOP_AMT").
		Where(startdateWhereClause(param)).
		Group("DS, MERCHANT_GROUP").
		Order("TOP_AMT DESC").
		Limit(10).
		Scan(&top_nominal_frekuensi_merchant).Error

	if err != nil {
		fmt.Println(err.Error())
		return data, period_date, modified_date, err
	}

	for _, element := range top_nominal_frekuensi_merchant {
		top_amt, _ := strconv.Atoi(element.TOP_AMT)
		data[element.MERCHANT_GROUP] = top_amt
	}

	row := db.
		Table("BRI360_SUMMARY_SPENDING_DEBIT").
		Select("MAX(MODIFIED_DATE) as modified_date").
		Where(startdateWhereClause(param)).
		Row()

	row.Scan(&modified_date)
	fmt.Println(modified_date)

	if err != nil {
		fmt.Println(err.Error())
		return data, period_date, modified_date, err
	}

	row = db.
		Table("BRI360_SUMMARY_SPENDING_DEBIT").
		Select("MIN(DS) AS min_date, MAX(DS) as max_date").
		Where(startdateWhereClause(param)).
		Row()

	row.Scan(&min_date, &max_date)

	period_date[0] = min_date
	period_date[1] = max_date
	// period_date[2] = total_date

	// fmt.Println(modified_date)
	fmt.Println(modified_date)

	if err != nil {
		fmt.Println(err.Error())
		return data, period_date, modified_date, err
	}

	return data, period_date, modified_date, nil
}

func (param RequestData) getPosisiSaldo() (statusSaldo, int64, string, error) {
	var posisi_saldo []PosisiSaldoResponse = []PosisiSaldoResponse{}
	var saldo statusSaldo = statusSaldo{}
	var saldo_rata2 map[string]float64 = map[string]float64{}
	var saldo_saat_ini map[string]float64 = map[string]float64{}

	//Create Connection With Phoenix
	var db, err = connect()

	var period_date int64
	// var min_date string
	// var max_date string
	var modified_date string

	if err != nil {
		fmt.Println(err.Error())
		return saldo, period_date, modified_date, err
	}

	err = db.
		Table("BRI360_SUMMARY_POSISI_SALDO").
		Select("SUM(POSISI_SALDO) as POSISI_SALDO, SUM(RATAS_SALDO) as RATAS_SALDO, DS").
		Where(startdateWhereClause(param)).
		Group("DS").
		Scan(&posisi_saldo).Error

	if err != nil {
		fmt.Println(err.Error())
		return saldo, period_date, modified_date, err
	}

	for _, element := range posisi_saldo {
		saldo_rata2[element.DS] = element.RATAS_SALDO
		saldo_saat_ini[element.DS] = element.POSISI_SALDO
	}

	saldo.SALDO_RATA_RATA = saldo_rata2
	saldo.SALDO_SAAT_INI = saldo_saat_ini

	if len(posisi_saldo) > 0 {
		row := db.
			Table("BRI360_SUMMARY_POSISI_SALDO").
			Select("MAX(MODIFIED_DATE) as modified_date").
			Where(startdateWhereClause(param)).
			Row()

		row.Scan(&modified_date)

		fmt.Println(modified_date)
		fmt.Println(period_date)
		fmt.Println("Hello Cok")

		if err != nil {
			fmt.Println(err.Error())
			return saldo, period_date, modified_date, err
		}
	}

	return saldo, period_date, modified_date, nil
}

func (param RequestData) getTopProduct() (interface{}, [2]string, string, error) {
	var top_product []TopProductResponse = []TopProductResponse{}
	var data map[string]int = map[string]int{}
	//Create Connection with Phoenix
	var db, err = connect()
	var period_date [2]string
	var min_date string
	var max_date string
	var modified_date string

	if err != nil {
		fmt.Println(err.Error())
		return data, period_date, modified_date, err
	}

	//Get Data From Phoenix Table
	err = db.
		Table("BRI360_SUMMARY_PRODUCT").
		Select("DS, PRODUCT_GROUP, sum(JUMLAH) AS TOP_JUMLAH").
		Where(startdateWhereClause(param)).
		Group("DS, PRODUCT_GROUP").
		Order("TOP_JUMLAH DESC").
		Limit(10).
		Scan(&top_product).Error

	if err != nil {
		fmt.Println(err.Error())
		return data, period_date, modified_date, err
	}

	for _, element := range top_product {
		top_jumlah, _ := strconv.Atoi(element.TOP_JUMLAH)
		data[element.PRODUCT_GROUP] = top_jumlah
	}

	row := db.
		Table("BRI360_SUMMARY_PRODUCT").
		Select("MAX(MODIFIED_DATE) as modified_date").
		Where(startdateWhereClause(param)).
		Row()

	row.Scan(&modified_date)
	fmt.Println(modified_date)

	if err != nil {
		fmt.Println(err.Error())
		return data, period_date, modified_date, err
	}

	row = db.
		Table("BRI360_SUMMARY_PRODUCT").
		Select("MIN(DS) AS min_date, MAX(DS) as max_date").
		Where(startdateWhereClause(param)).
		Row()

	row.Scan(&min_date, &max_date)

	period_date[0] = min_date
	period_date[1] = max_date
	// period_date[2] = total_date

	// fmt.Println(modified_date)
	fmt.Println(modified_date)

	if err != nil {
		fmt.Println(err.Error())
		return data, period_date, modified_date, err
	}

	return data, period_date, modified_date, nil
}

//
// func (param RequestData)  getTopProduct() (interface{}, [2]string, string, error){
// 	var top_product []TopProductResponse = []TopProductResponse
// 	var data map[string]int=map[string] int
// }
