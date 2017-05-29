package models

type TagStorer interface {
	FindTag(*Tag) error
	FindTags() ([]Tag, error)
	FindTagOrInit(string) (Tag, error)
}

type Tag struct {
	ID            uint
	Name          string `gorm:"unique_index:index_tags_on_name"`
	TaggingsCount uint
	Articles      []Article `gorm:"many2many:taggings;"`
}

func (db *DB) FindTag(tag *Tag) error {
	return db.Where(&tag).Find(tag).Error
}

func (db *DB) FindTags() ([]Tag, error) {
	var tags []Tag
	err := db.Find(&tags).Error
	return tags, err
}

func (db *DB) FindTagOrInit(tagName string) (tag Tag, err error) {
	err = db.DB.FirstOrInit(&tag, Tag{Name: tagName}).Error
	return
}
