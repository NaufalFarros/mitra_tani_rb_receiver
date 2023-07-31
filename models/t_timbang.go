package models

import "time"

type TTimbang struct {
	NoSpat              string    `gorm:"column:no_spat;NOT NULL" bson:"no_spat"` // no spat
	IDTimbangan         int       `gorm:"column:id_timbangan;primary_key;AUTO_INCREMENT" bson:"id_timbangan"`
	IDSpat              int       `gorm:"column:id_spat;default:0;NOT NULL" bson:"id_spat"`
	LokasiTimbang1      string    `gorm:"column:lokasi_timbang_1" bson:"lokasi_timbang_1"` // id_timbangan
	LokasiTimbang2      string    `gorm:"column:lokasi_timbang_2" bson:"lokasi_timbang_2"` // id_timbangan
	Bruto               int       `gorm:"column:bruto;default:0" bson:"bruto"`
	Tara                int       `gorm:"column:tara;default:0" bson:"tara"`
	Netto               int       `gorm:"column:netto;default:0" bson:"netto"`
	NettoFinal          int       `gorm:"column:netto_final;default:0" bson:"netto_final"`                 // nilai netto akhir setelah rafaksi
	NettoRafaksi        int       `gorm:"column:netto_rafaksi" bson:"netto_rafaksi"`                       // nilai neto rafaksi
	RafaksiProsentis    float64   `gorm:"column:rafaksi_prosentis;default:0" bson:"rafaksi_prosentis"`     // prosentis rafaksi
	TglBruto            time.Time `gorm:"column:tgl_bruto" bson:"tgl_bruto"`                               // waktu pengambilan bruto
	TglTara             time.Time `gorm:"column:tgl_tara" bson:"tgl_tara"`                                 // waktu pengambilan tara
	TglNetto            time.Time `gorm:"column:tgl_netto" bson:"tgl_netto"`                               // waktu pengambilan netto
	TglRafaksi          time.Time `gorm:"column:tgl_rafaksi" bson:"tgl_rafaksi"`                           // waktu entry data rafaksi
	TransloadingStatus  int       `gorm:"column:transloading_status;default:0" bson:"transloading_status"` // 0=tidak transloading,1=transloading
	NoTransloading      string    `gorm:"column:no_transloading" bson:"no_transloading"`                   // no lori, no kontainer, no emplasement
	PtgsTransloading    string    `gorm:"column:ptgs_transloading" bson:"ptgs_transloading"`               // nama petugas transloading
	PtgsTimbang1        string    `gorm:"column:ptgs_timbang_1" bson:"ptgs_timbang_1"`                     // user penimbang
	PtgsTimbang2        string    `gorm:"column:ptgs_timbang_2" bson:"ptgs_timbang_2"`                     // user penimbang
	TglTransloading     time.Time `gorm:"column:tgl_transloading" bson:"tgl_transloading"`                 // waktu tranloading
	MultiSling          string    `gorm:"column:multi_sling" bson:"multi_sling"`                           // text penyimpanan multisling
	NettoSebelumKoreksi int       `gorm:"column:netto_sebelum_koreksi" bson:"netto_sebelum_koreksi"`       // netto sebelum terjadi koreksi timbangan
	KetKoreksiTimbangan string    `gorm:"column:ket_koreksi_timbangan" bson:"ket_koreksi_timbangan"`       // keterangan penyebab koreksi
	TrainStat           string    `gorm:"column:train_stat" bson:"train_stat"`                             // no trainstat
	NoLori              string    `gorm:"column:no_lori" bson:"no_lori"`                                   // no lori
	NoLoko              string    `gorm:"column:no_loko" bson:"no_loko"`                                   // no_loko
}
