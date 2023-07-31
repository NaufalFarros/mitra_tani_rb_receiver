package models

import "time"

type TSpta struct {
	ID                    int       `gorm:"column:id;primary_key;AUTO_INCREMENT" bson:"id"`        // id spta
	NoSpat                string    `gorm:"column:no_spat;NOT NULL" bson:"no_spat"`                // no spat
	KodePlant             string    `gorm:"column:kode_plant" bson:"kode_plant"`                   // kode pabrik gula pembuat spta
	KodeBlok              string    `gorm:"column:kode_blok;NOT NULL" bson:"kode_blok"`            // kode sap
	TglSpta               time.Time `gorm:"column:tgl_spta" bson:"tgl_spta"`                       // tgl spta di buat
	TglExpired            time.Time `gorm:"column:tgl_expired" bson:"tgl_expired"`                 // tgl akhir masa berlaku spta
	TglCetak              time.Time `gorm:"column:tgl_cetak" bson:"tgl_cetak" `                    // tgl_cetak spta
	PersnoPta             string    `gorm:"column:persno_pta" bson:"persno_pta"`                   // id_sap
	IDPetaniSap           string    `gorm:"column:id_petani_sap" bson:"id_petani_sap"`             // id petani dari sap
	TebangPg              int       `gorm:"column:tebang_pg;default:0" bson:"tebang_pg"`           // 0 = tidak tebang pg, 1= tebang pg
	AngkutPg              int       `gorm:"column:angkut_pg;default:0" bson:"angkut_pg"`           // 0 = tidak angkut pg, 1 = angkut pg
	JenisSpta             string    `gorm:"column:jenis_spta" bson:"jenis_spata"`                  // jenis spta, 1=tebang, 2=angkut, 3=tebang angkut
	JarakID               int       `gorm:"column:jarak_id" bson:"jarak_id"`                       // jarak id
	VendorAngkut          int       `gorm:"column:vendor_angkut" bson:"vendor_angkut"`             // vendor angkut
	KodeAffd              string    `gorm:"column:kode_affd" bson:"kode_affd"`                     // refrens sap_m_affdeling kode
	KodeKatLahan          string    `gorm:"column:kode_kat_lahan" bson:"kode_kat_lahan"`           // refrens sap_m_kat_lahan kode
	KodePlantTrasnfer     string    `gorm:"column:kode_plant_trasnfer" bson:"kode_plant_transfer"` // kode pg pengirim tebu
	KodePlantKe           string    `gorm:"column:kode_plant_ke" bson:"kode_plant_ke"`
	MetodeTma             int       `gorm:"column:metode_tma" bson:"metode_tma"`                         // 1=manual,2=semi mekanisasi,3=mekanisasi
	Ket                   string    `gorm:"column:ket" bson:"ket"`                                       // ket spta
	IDJenisAngkutan       int       `gorm:"column:id_jenis_angkutan" bson:"id_jenis_angkutan"`           // mengacu pada tabel m_jenis_angkutan
	BuatSptaStatus        int       `gorm:"column:buat_spta_status;default:0" bson:"buat_spta_status"`   // 0 di retur 1. di buat
	BuatSptaTgl           time.Time `gorm:"column:buat_spta_tgl" bson:"buat_spta_tgl"`                   // tgl saat buat spta
	CetakSptaStatus       int       `gorm:"column:cetak_spta_status;default:0" bson:"cetak_spta_status"` // 0,1
	CetakSptaTgl          time.Time `gorm:"column:cetak_spta_tgl" bson:"cetak_spta_tgl"`
	SelektorStatus        int       `gorm:"column:selektor_status;default:0" bson:"selektor_status"`       // 0,1,2=ditolak
	SelektorTgl           time.Time `gorm:"column:selektor_tgl" bson:"selektor_tgl"`                       // tgl saat selektor
	PintuMasukStatus      int       `gorm:"column:pintu_masuk_status;default:0" bson:"pintu_masuk_status"` // 0,1
	PintuMasukTgl         time.Time `gorm:"column:pintu_masuk_tgl" bson:"pintu_masuk_tgl"`
	TimbBrutoStatus       int       `gorm:"column:timb_bruto_status;default:0" bson:"timb_bruto_status"` // 0,1
	TimbBrutoTgl          time.Time `gorm:"column:timb_bruto_tgl" bson:"timb_bruto_tgl"`
	TimbNettoStatus       int       `gorm:"column:timb_netto_status;default:0" bson:"Timb_netto_status"` // 0,1
	TimbNettoTgl          time.Time `gorm:"column:timb_netto_tgl" bson:"timb_netto_tgl"`                 // tgl saat timb netto
	MejaTebuStatus        int       `gorm:"column:meja_tebu_status;default:0" bson:"meja_tebu_status"`   // 0,1
	MejaTebuTgl           time.Time `gorm:"column:meja_tebu_tgl" bson:"meja_tebu_tgl"`                   // tgl saat meja tebu
	AriStatus             int       `gorm:"column:ari_status;default:0" bson:"ari_status"`               // 0,1,2=ditolak
	AriTgl                time.Time `gorm:"column:ari_tgl" bson:"ari_tgl"`
	TglTimbang            time.Time `gorm:"column:tgl_timbang" bson:"tgl_timbang"`           // tgl saat timbang
	HariGiling            int       `gorm:"column:hari_giling;default:0" bson:"hari_giling"` // hari giling pabrik
	TglGiling             time.Time `gorm:"column:tgl_giling" bson:"tgl_giling"`
	NoUrutAnalisaRendemen int       `gorm:"column:no_urut_analisa_rendemen;default:0" bson:"no_urut_analisa_rendemen"` // no urut analisa rendemen
	ReturAlasan           string    `gorm:"column:retur_alasan" bson:"retur_alasan"`
	ReturStatus           int       `gorm:"column:retur_status;default:0" bson:"retur_status"` // 0. belum 1. sudah 2. validasi 3. kirim sap
	ReturTgl              time.Time `gorm:"column:retur_tgl" bson:"retur_tgl"`
	UpahTebangStatus      int       `gorm:"column:upah_tebang_status;default:0" bson:"upah_tebang_status"` // 0. belum 1. sudah 2. validasi 3. kirim sap
	UpahTebangTgl         time.Time `gorm:"column:upah_tebang_tgl" bson:"upah_tebang_tgl"`
	UpahAngkutStatus      int       `gorm:"column:upah_angkut_status;default:0" bson:"upah_angkut_status"` // 0. belum 1. sudah 2. validasi 3. sap
	UpahAngkutTgl         time.Time `gorm:"column:upah_angkut_tgl" bson:"upah_angkut_tgl"`
	SbhStatus             int       `gorm:"column:sbh_status;default:0" bson:"sbh_status"` // 0. belum 1. sudah 2. pengolahan 3. tanaman 4. AKU
	SbhTgl                time.Time `gorm:"column:sbh_tgl" bson:"sbh_tgl"`
	SbhSap                int       `gorm:"column:sbh_sap;default:0" bson:"sbh_sap"`
	SbhSapTgl             time.Time `gorm:"column:sbh_sap_tgl" bson:"sbh_sap_tgl"`
	RfidSticker           string    `gorm:"column:rfid_sticker" bson:"rfid_sticker"`
	RfidStickerTagging    time.Time `gorm:"column:rfid_sticker_tagging" bson:"rfid_sticker_tagging"`
	RfidStickerStatus     int       `gorm:"column:rfid_sticker_status;default:0" bson:"rfid_sticker_status"` // 0. kosong 1. sudah tag 2. close tag
	RfidStickerClosetag   time.Time `gorm:"column:rfid_sticker_closetag" bson:"rfid_sticker_closetag"`
	IDTruck               string    `gorm:"column:id_truck" bson:"id_truck"`
	PersnoMandor          string    `gorm:"column:persno_mandor" bson:"persno_mandor"`
	StatusDistribusi      int       `gorm:"column:status_distribusi;default:0" bson:"status_distribusi"` // 0. belum distribusi 1. sudah distribusi 2. non distribusi
	TglDistribusi         time.Time `gorm:"column:tgl_distribusi" bson:"tgl_distribusi"`
	SptStatus             int       `gorm:"column:spt_status;default:0" bson:"spt_status"`
	NaturaStatus          int       `gorm:"column:natura_status;default:0" bson:"natura_status"`
	TgljamTebang          time.Time `gorm:"column:tgljam_tebang" bson:"tgljam_tebang"`
	HaTertebang           float64   `gorm:"column:ha_tertebang" bson:"ha_tertebang"`
	JenisPembayaran       string    `gorm:"column:jenis_pembayaran" bson:"jenis_pembayaran"` // jenis pembayaran
	SumberDana            int       `gorm:"column:sumber_dana;default:0" bson:"sumber_dana"` // 0. PTPN 1. PEN
	SptSpKategori         string    `gorm:"column:spt_sp_kategori" bson:"spt_sp_kategori"`
	NoAnalisa             string    `gorm:"column:no_analisa;default:0" bson:"no_analisa"` // digenerate saat masuk selektor / masuk coresampler jika coresampler
}
