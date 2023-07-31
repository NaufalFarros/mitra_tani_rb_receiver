package models

import "time"

type TAri struct {
	NoSpat              string    `gorm:"column:no_spat;NOT NULL" bson:"no_spta"` // no spat
	IDAri               int       `gorm:"column:id_ari;primary_key;AUTO_INCREMENT" bson:"id_ari"`
	IDSpta              int       `gorm:"column:id_spta;NOT NULL" bson:"id_spta"` // refrens t_spta id_spta
	PolTebu             float64   `gorm:"column:pol_tebu;default:0.00" bson:"pol_tebu"`
	PersenBrixAri       float64   `gorm:"column:persen_brix_ari;default:0.00" bson:"persen_brix_ari"`     // ari
	PersenPolAri        float64   `gorm:"column:persen_pol_ari;default:0.00" bson:"persen_pol_ari"`       // ari
	PhAri               float64   `gorm:"column:ph_ari;default:0.00" bson:"ph_ari"`                       // ari
	Hk                  float64   `gorm:"column:hk;default:0.00" bson:"hk"`                               // ari
	FaktorPerah         float64   `gorm:"column:faktor_perah;default:0.0000" bson:"faktor_perah"`         // faktor perah dari liquid
	FaktorRegresi       float64   `gorm:"column:faktor_regresi;default:0.0000" bson:"faktor_regresi"`     // faktor regresi kedawung / faktor Kristal Rendemen
	NilaiNira           float64   `gorm:"column:nilai_nira;default:0.00" bson:"nilai_nira"`               // ari
	FaktorRendemen      float64   `gorm:"column:faktor_rendemen;default:0.00" bson:"faktor_rendemen"`     // jatiroto metod menjadi faktor perah
	RendemenAri         float64   `gorm:"column:rendemen_ari;default:0.00" bson:"rendemen_ari"`           // jatmed menjadi rendemen contoh
	FaktorKonversi      float64   `gorm:"column:faktor_konversi;default:0.00" bson:"faktor_konversi"`     // faktor konversi
	RendemenIndividu    float64   `gorm:"column:rendemen_individu;default:0.00" bson:"rendemen_individu"` // jatmed
	HablurAri           float64   `gorm:"column:hablur_ari;default:0.00" bson:"hablur_ari"`               // ari
	GulaTotal           float64   `gorm:"column:gula_total;default:0.00" bson:"gula_total"`               // ari
	TetesTotal          float64   `gorm:"column:tetes_total;default:0.00" bson:"tetes_total"`             // ari
	RendemenPtr         float64   `gorm:"column:rendemen_ptr;default:0.0000" bson:"rendemen_ptr"`         // ari
	GulaPtr             float64   `gorm:"column:gula_ptr;default:0.00" bson:"gula_ptr"`                   // ari
	TetesPtr            float64   `gorm:"column:tetes_ptr;default:0.00" bson:"tetes_ptr"`                 // ari
	GulaPg              float64   `gorm:"column:gula_pg;default:0.00" bson:"gula_pg"`                     // ari
	TetesPg             float64   `gorm:"column:tetes_pg;default:0.00" bson:"tetes_pg"`                   // ari
	DitolakAri          int       `gorm:"column:ditolak_ari;default:0" bson:"ditolak_ari"`                // 0 , 1 dotolak
	DitolakAlasan       string    `gorm:"column:ditolak_alasan" bson:"ditolak_alasan"`                    // ari
	TglAri              time.Time `gorm:"column:tgl_ari" bson:"tgl_ari"`                                  // ari
	PtgsAri             string    `gorm:"column:ptgs_ari" bson:"ptgs_ari"`                                // ari
	SbhAriStatus        int       `gorm:"column:sbh_ari_status;default:0" bson:"sbh_ari_status"`          // 0. belum 1. sudah 2. pengolahan 3. tanaman 4. AKU
	SbhAriTgl           time.Time `gorm:"column:sbh_ari_tgl" bson:"sbh_ari_tgl"`                          // ari
	SbhAriUser          string    `gorm:"column:sbh_ari_user" bson:"sbh_ari_user"`                        // ari
	PengolahanStatus    int       `gorm:"column:pengolahan_status;default:0" bson:"pengolahan_status"`    // 0. belum 1. sudah 2. pengolahan 3. tanaman 4. AKU
	PengolahanTgl       time.Time `gorm:"column:pengolahan_tgl" bson:"pengolahan_tgl"`                    // ari
	PengolahanUser      string    `gorm:"column:pengolahan_user" bson:"pengolahan_user"`                  // ari
	TanamanStatus       int       `gorm:"column:tanaman_status;default:0" bson:"tanaman_status"`          // 0. belum 1. sudah 2. pengolahan 3. tanaman 4. AKU
	TanamanTgl          time.Time `gorm:"column:tanaman_tgl" bson:"tanaman_tgl"`                          // ari
	TanamanUser         string    `gorm:"column:tanaman_user" bson:"tanaman_user"`                        // ari
	AkuStatus           int       `gorm:"column:aku_status;default:0" bson:"aku_status"`                  // 0. belum 1. sudah 2. pengolahan 3. tanaman 4. AKU
	AkuTgl              time.Time `gorm:"column:aku_tgl" bson:"aku_tgl"`                                  // ari
	AkuUser             string    `gorm:"column:aku_user" bson:"aku_user"`                                // ari
	SembilanpuluhPersen float64   `gorm:"column:sembilanpuluh_persen;default:0.00" bson:"sembilanpuluh_persen"`
	SepuluhPersen       float64   `gorm:"column:sepuluh_persen;default:0.00" bson:"sepuluh_persen"`
	RSpg                float64   `gorm:"column:r_spg;default:0.0000" bson:"r_spg"`
	KopensasiGula       float64   `gorm:"column:kopensasi_gula;default:0.0000" bson:"kopensasi_gula"`
	Coresampler         int       `gorm:"column:coresampler;default:0" bson:"coresampler"`
}
