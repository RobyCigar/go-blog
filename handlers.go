package main

import (
	"strconv"

	"github.com/gin-gonic/gin"

	entity "github.com/robycigar/goblog/entity"
)

func Index(c *gin.Context) {
	// Get query parameters for pagination, sorting, and searching
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	sortBy := c.Query("sortBy")
	searchTerm := c.Query("search")
	db := dbConn()

	// Set default values if not provided
	if page == "" {
		page = "1"
	}
	if pageSize == "" {
		pageSize = "10"
	}

	// Convert string parameters to integers
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		c.JSON(400, gin.H{"error": "Invalid page number"})
		return
	}

	pageSizeNum, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeNum < 1 {
		c.JSON(400, gin.H{"error": "Invalid page size"})
		return
	}

	// Calculate offset for pagination
	offset := (pageNum - 1) * pageSizeNum

	// Query authors from the database with pagination, sorting, and searching
	var authors []entity.Author
	query := db.Offset(offset).Limit(pageSizeNum)

	if sortBy != "" {
		// If sortBy parameter is provided, add sorting
		query = query.Order(sortBy)
	}

	if searchTerm != "" {
		// If search parameter is provided, add searching
		query = query.Where("name LIKE ?", "%"+searchTerm+"%")
	}

	result := query.Find(&authors)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	// Return paginated, sorted, and searched authors
	c.JSON(200, authors)
}

func CreateAuthor(c *gin.Context) {
	var input entity.Author
	db := dbConn()

	// Bind JSON request to Author struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create author in the database
	db.Create(&input)

	// Return success response
	c.JSON(201, input)
}

// Create a new post with a relation to an author
func CreatePost(c *gin.Context) {
	var request struct {
		Title    string `json:"title" binding:"required"`
		Content  string `json:"content" binding:"required"`
		AuthorID uint   `json:"author_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db := dbConn()

	// Check if the author exists
	var author entity.Author
	if db.First(&author, request.AuthorID).Error != nil {
		c.JSON(404, gin.H{"error": "Author not found"})
		return
	}

	// Create a new post
	post := entity.Post{
		Title:    request.Title,
		Content:  request.Content,
		AuthorID: request.AuthorID,
	}

	// Save the post to the database
	db.Create(&post)

	c.JSON(201, post)
}
