package models

import "time"

type TMejaTebu struct {
	NoSpat        string    `gorm:"column:no_spat;NOT NULL" bson:"no_spat"` // no spat
	IDMejatebu    int       `gorm:"column:id_mejatebu;primary_key;AUTO_INCREMENT" bson:"id_mejatebu"`
	IDSpta        int       `gorm:"column:id_spta;NOT NULL" bson:"id_spta"`      // refrens t_spta id
	Daduk         int       `gorm:"column:daduk" bson:"daduk"`                   // meja tebu
	Sogolan       int       `gorm:"column:sogolan" bson:"sogolan"`               // meja tebu
	Pucuk         int       `gorm:"column:pucuk" bson:"pucuk"`                   // meja tebu
	AkarTanah     int       `gorm:"column:akar_tanah" bson:"akar_tanah"`         // meja tebu
	NonTebu       int       `gorm:"column:non_tebu" bson:"non_tebu"`             // meja tebu
	Terbakar      int       `gorm:"column:terbakar" bson:"terbakar"`             // meja tebu
	Cacahan       int       `gorm:"column:cacahan" bson:"cacahan"`               // meja tebu
	Brondolan     int       `gorm:"column:brondolan" bson:"brondolan"`           // meja tebu
	KondisiTebu   string    `gorm:"column:kondisi_tebu" bson:"kondisi_tebu"`     // meja tebu
	PtgsMejaTebu  string    `gorm:"column:ptgs_meja_tebu" bson:"ptgs_meja_tebu"` // meja tebu
	Gilingan      int       `gorm:"column:gilingan" bson:"gilingan"`             // meja tebu
	KodeMejaTebu  string    `gorm:"column:kode_meja_tebu" bson:"kode_meja_tebu"` // meja tebu
	WarnaMejaTebu string    `gorm:"column:warna_meja_tebu" bson:"warna_meja_tebu"`
	TglMejaTebu   time.Time `gorm:"column:tgl_meja_tebu" bson:"tgl_meja_tebu"`             // meja tebu
	RafraksiAktif int       `gorm:"column:rafraksi_aktif;default:0" bson:"rafraksi_aktif"` // 0=rafaksi tidak aktif, 1=rafaksi aktif

}
