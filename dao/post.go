package dao

import (
	"go-blog/models"
	"log"
)

func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update blog_post set title=?,content=?,markdown=?,category_id=?,type=?,slug=?,update_at=? where pid=?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.Pid,
	)
	if err != nil {
		log.Println(err)
	}
}

func SavePost(post *models.Post) {
	ret, err := DB.Exec("insert into blog_post "+
		"(title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) "+
		"values(?,?,?,?,?,?,?,?,?,?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt,
	)
	if err != nil {
		log.Println(err)
	}
	pid, _ := ret.LastInsertId()
	post.Pid = int(pid)
}

func CountGetAllPost() int {
	row := DB.QueryRow("select count(*) from blog_post")
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var count int
	row.Scan(&count)
	return count
}
func CountGetAllPostByCategoryId(cid int) int {
	row := DB.QueryRow("select count(*) from blog_post where category_id = ?", cid)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var count int
	row.Scan(&count)
	return count
}

func GetPostPage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post limit ?,?", page, pageSize)
	if err != nil {
		log.Println("GetAllPost Failed:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("GetAllPost Scan Failed:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetSearchPost(condition string) ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post where title like ? ", "%"+condition+"%")
	if err != nil {
		log.Println("GetAllPost Failed:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("GetAllPost Scan Failed:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetALLPost() ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post")
	if err != nil {
		log.Println("GetAllPost Failed:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("GetAllPost Scan Failed:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetCategoryPage(page, pageSize, cid int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where category_id = ? limit ?,?", cid, page, pageSize)
	if err != nil {
		log.Println("GetAllPost Failed:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("GetAllPost Scan Failed:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func DeletePost(pid int) (*models.Post, error) {
	_, err := DB.Exec("delete from blog_post where pid=?", pid)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	post := &models.Post{
		Pid: pid,
	}
	return post, nil
}

func GetPostById(pid int) (*models.Post, error) {
	row := DB.QueryRow("select * from blog_post where pid =?", pid)
	if row.Err() != nil {
		log.Println("GetPost Falied:", row.Err())
		return nil, row.Err()
	}
	var post models.Post
	err := row.Scan(
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt,
	)
	if err != nil {
		log.Println("GetPost Scan Failed:", err)
		return nil, err
	}
	return &post, nil
}
