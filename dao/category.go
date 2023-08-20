package dao

import (
	"go-blog/models"
	"log"
)

func GetCategoryNameById(category_id int) string {
	row := DB.QueryRow("select name from blog_category where cid = ?", category_id)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var category string
	row.Scan(&category)

	return category
}

func GetALLCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("GetAllCategory Failed:", err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory Scan Failed:", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}
