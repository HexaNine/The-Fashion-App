package entity

type CategoryModel struct {
	ID        string `bson:"_id,omitempty" json:"category_id"`
	Name      string `bson:"name" json:"category_name"`
	ProductID string `bson:"product_id" json:"product_id"`
}