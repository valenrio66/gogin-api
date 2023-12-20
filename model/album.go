package model

type AlbumDb struct {
	IDAlbum string  `gorm:"primaryKey;column:id_album" json:"-"`
	Judul   string  `gorm:"column:judul" json:"judul"`
	Artis   string  `gorm:"column:artis" json:"artis"`
	Harga   float64 `gorm:"column:harga" json:"harga"`
}

func (a *AlbumDb) TableName() string {
	return "album"
}
