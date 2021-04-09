package main

//RequestData model
type RequestData struct {
	Branch     string `json:"branch"`
	Region     string `json:"region"`
	Mainbr     string `json:"mainbr"`
	Page       string `json:"page"`
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
	Ds         string `json:"ds"`
	// RGDESC     string  `json:"rgdesc"`
	// MBDESC     string  `json:"mbdesc"`
	// BRDESC     string  `json:"brdesc"`
}

//topfreq
type topfreqsummary struct {
	TOP_FREQ string `json:"top_freq"`
}

// fasilitas segmentasi
type Fasilitassegmentasi struct {
	GENERATOR     string `json:"generator"`
	LEAKER        string `json:"leaker"`
	SAVER_PASSIVE string `json:"saver_passive"`
	SAVER_ACTIVE  string `json:"saver_active"`
	NOT_SEGMENTED string `json:"not_segmented"`
}

//ResponsChurn model

type ChurnResponse struct {
	REGION        string `json:"region" gorm:"column:REGION"`
	RGDESC        string `json:"rgdesc" gorm:"column:RGDESC"`
	MAINBR        string `json:"mainbr" gorm:"column:MAINBR"`
	MBDESC        string `json:"mbdesc" gorm:"column:MBDESC"`
	BRANCH        string `json:"branch" gorm:"column:BRANCH"`
	BRDESC        string `json:"brdesc" gorm:"column:BRDESC"`
	CHURN_LEVEL   string `json:"churn_level" gorm:"column:CHURN_LEVEL"`
	TOTAL_CHURN   string `json:"total_churn" gorm:"column:TOTAL_CHURN"`
	DS            string `json:"ds" gorm:"column:DS"`
	POSISI_DATE   float64 `json:"posisi_date" gorm:"column:POSISI_DATE"`
	MODIFIED_DATE string `json:"modified_date" gorm:"column:MODIFIED_DATE"`
}

type statusSegmentasi struct {
	PASSER        Fasilitassegmentasi `json:"PASSER"`
	GENERATOR     Fasilitassegmentasi `json:"GENERATOR"`
	LEAKER        Fasilitassegmentasi `json:"LEAKER"`
	SAVER_PASSIVE Fasilitassegmentasi `json:"SAVER_PASSIVE"`
	SAVER_ACTIVE  Fasilitassegmentasi `json:"SAVER_ACTIVE"`
	NOT_SEGMENTED Fasilitassegmentasi `json:"NOT_SEGMENTED"`
}

// aktif dan tidak aktif
type aktif_tidakaktif struct {
	AKTIF       float64 `json:"aktif"`
	TIDAK_AKTIF float64 `json:"tidak_aktif"`
}

type statusFasilitas struct {
	REG_SMS            aktif_tidakaktif `json:"reg_sms"`
	SMS_FIN            aktif_tidakaktif `json:"sms_fin"`
	REG_INTERNET       aktif_tidakaktif `json:"reg_internet"`
	INTERNET_FIN       aktif_tidakaktif `json:"internet_fin"`
	REG_PHONE          aktif_tidakaktif `json:"reg_phone"`
	BRIVA_AUTOPAYMENT  aktif_tidakaktif `json:"briva_autopayment"`
	DPLK_AUTOPAYMENT   aktif_tidakaktif `json:"dplk_autopayment"`
	FIF_AUTOPAYMENT    aktif_tidakaktif `json:"fif_autopayment"`
	PLN_AUTOPAYMENT    aktif_tidakaktif `json:"pln_autopayment"`
	TELKOM_AUTOPAYMENT aktif_tidakaktif `json:"telkom_autopayment"`
	NOTIF_SMS          aktif_tidakaktif `json:"notif_sms"`
	NOTIF_EMAIL        aktif_tidakaktif `json:"notif_email"`
}

type FasilitasRekeningResponse struct {
	STATUS             string  `json:"status" gorm:"column:STATUS"`
	REG_SMS            float64 `json:"reg_sms" gorm:"column:REG_SMS"`
	SMS_FIN            float64 `json:"sms_fin" gorm:"column:SMS_FIN"`
	REG_INTERNET       float64 `json:"reg_internet" gorm:"column:REG_INTERNET"`
	INTERNET_FIN       float64 `json:"internet_fin" gorm:"column:INTERNET_FIN"`
	REG_PHONE          float64 `json:"reg_phone" gorm:"column:REG_PHONE"`
	BRIVA_AUTOPAYMENT  float64 `json:"briva_autopayment" gorm:"column:BRIVA_AUTOPAYMENT"`
	DPLK_AUTOPAYMENT   float64 `json:"dplk_autopayment" gorm:"column:DPLK_AUTOPAYMENT"`
	FIF_AUTOPAYMENT    float64 `json:"fif_autopayment" gorm:"column:FIF_AUTOPAYMENT"`
	PLN_AUTOPAYMENT    float64 `json:"pln_autopayment" gorm:"column:PLN_AUTOPAYMENT"`
	TELKOM_AUTOPAYMENT float64 `json:"telkom_autopayment" gorm:"column:TELKOM_AUTOPAYMENT"`
	NOTIF_SMS          float64 `json:"notif_sms" gorm:"column:NOTIF_SMS"`
	NOTIF_EMAIL        float64 `json:"notif_email" gorm:"column:NOTIF_EMAIL"`
	POSISI_DATE        float64 `json:"posisi_date" gorm:"column:POSISI_DATE"`
	MODIFIED_DATE      string  `json:"modified_date" gorm:"column:MODIFIED_DATE"`
}

type SegmentasiResponse struct {
	PASSER        float64 `json:"passer" gorm:"column:PASSER"`
	GENERATOR     float64 `json:"generator" gorm:"column:GENERATOR"`
	LEAKER        float64 `json:"leaker" gorm:"column:LEAKER"`
	SAVER_PASSIVE float64 `json:"saver_passive" gorm:"column:SAVER_PASSIVE"`
	SAVER_ACTIVE  float64 `json:"saver_active" gorm:"column:SAVER_ACTIVE"`
	NOT_SEGMENTED float64 `json:"not_segmented" gorm:"column:NOT_SEGMENTED"`
	MODIFIED_DATE string  `json:"modified_date" gorm:"column:MODIFIED_DATE"`
}

type RatassaldoResponse struct {
	NAIK  float64 `json:"naik" gorm:"column:NAIK"`
	TURUN float64 `json:"turun" gorm:"column:TURUN"`
}

type TopFrekuensiMerchantResponse struct {
	MERCHANT_GROUP string `json:"merchant_group" gorm:"column:MERCHANT_GROUP"`
	TOP_FREQ       string `json:"top_freq" gorm:"column:TOP_FREQ"`
}

type TopNominalFrekuensiMerchantResponse struct {
	MERCHANT_GROUP string `json:"merchant_group" gorm:"column:MERCHANT_GROUP"`
	TOP_AMT        string `json:"top_amt" gorm:"column:TOP_AMT"`
}

type PosisiSaldoResponse struct {
	POSISI_SALDO  float64 `json:"posisi_saldo" gorm:"column:POSISI_SALDO"`
	RATAS_SALDO   float64 `json:"ratas_saldo" gorm:"column:RATAS_SALDO"`
	DS            string  `json:"ds" gorm:"column:DS"`
	MODIFIED_DATE string  `json:"modified_date" gorm:"column:MODIFIED_DATE"`
}

type statusSaldo struct {
	SALDO_RATA_RATA map[string]float64 `json:"saldo_rata_rata"`
	SALDO_SAAT_INI  map[string]float64 `json:"saldo_saat_ini"`
}

type TopProductResponse struct {
	PRODUCT_GROUP string `json:"product_group" gorm:"column:PRODUCT_GROUP"`
	TOP_JUMLAH    string `json:"top_jumlah" gorm:"column:TOP_JUMLAH"`
}

// type PosisiSaldo struct {

// }
