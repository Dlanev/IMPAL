package models

type PlaylistSong struct {
	ID uint `gorm:"primaryKey"`

	PlaylistID uint `json:"playlist_id"`
	SongID uint `json:"song_id"`

	Song Song `gorm:"foreignKey:SongID" json:"song"`
}