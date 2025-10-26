package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// 全局配置
const (
	TokenExpirationTime = time.Hour * 24 // Token有效期为24小时
)

// 用户注册请求结构体
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// 用户登录请求结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 创建文章请求结构体
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// 更新文章请求结构体
type UpdatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// 创建评论请求结构体
type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

func main() {
	// 连接数据库
	ConnectDatabase()

	// 初始化 Gin 引擎
	router := gin.Default()

	// 公共路由组
	public := router.Group("/")
	{
		// 用户注册
		public.POST("/register", register)
		// 用户登录
		public.POST("/login", login)
		// 获取所有文章列表
		public.GET("/posts", getPosts)
		// 获取单篇文章详情
		public.GET("/posts/:id", getPost)
		// 获取文章的评论列表
		public.GET("/posts/:id/comments", getCommentsByPostID)
	}

	// 需要认证的路由组
	private := router.Group("/")
	private.Use(JWTAuthMiddleware())
	{
		// 创建文章
		private.POST("/posts", createPost)
		// 更新文章
		private.PUT("/posts/:id", updatePost)
		// 删除文章
		private.DELETE("/posts/:id", deletePost)
		// 创建评论
		private.POST("/posts/:id/comments", createComment)
	}

	// 启动服务器
	router.Run(":8080")
}

// 用户注册函数
func register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查用户名是否已存在
	var existingUser User
	if err := DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "用户名已经存在"})
		return
	}

	// 检查邮箱是否已存在
	if err := DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "邮件地址已经存在"})
		return
	}

	// 密码加密
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码设置失败"})
		return
	}

	// 创建新用户
	user := User{
		Username: req.Username,
		Password: string(passwordHash),
		Email:    req.Email,
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "用户请求失败"})
}

// 用户登录函数
func login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查找用户
	var user User
	if err := DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码无效"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码无效"})
		return
	}

	// 生成JWT Token
	expirationTime := time.Now().Add(TokenExpirationTime)
	claims := &jwt.MapClaims{
		"user_id": user.ID,
		"exp":     expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token验证失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":      tokenString,
		"expires_at": expirationTime.Format(time.RFC3339),
	})
}

// 创建文章函数
func createPost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未经过身份验证"})
		return
	}

	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID.(uint),
	}

	if err := DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建文章失败"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// 获取所有文章列表函数
func getPosts(c *gin.Context) {
	var posts []Post
	if err := DB.Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取列表失败"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// 获取单篇文章详情函数
func getPost(c *gin.Context) {
	id := c.Param("id")
	var post Post
	if err := DB.Preload("User").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "获取文章详情失败"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// 更新文章函数
func updatePost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未经过身份验证"})
		return
	}

	id := c.Param("id")
	var post Post
	if err := DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章列表未发现"})
		return
	}

	// 检查用户是否有权限更新文章
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "没有权限修改内容"})
		return
	}

	var req UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新文章
	post.Title = req.Title
	post.Content = req.Content
	post.UpdatedAt = time.Now()

	if err := DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "修改失败"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// 删除文章函数
func deletePost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未经过身份验证"})
		return
	}

	id := c.Param("id")
	var post Post
	if err := DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	// 检查用户是否有权限删除文章
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "没有权限删除文章"})
		return
	}

	if err := DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// 创建评论函数
func createComment(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未经过身份验证"})
		return
	}

	postID := c.Param("id")

	// 检查文章是否存在
	var post Post
	if err := DB.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := Comment{
		Content: req.Content,
		UserID:  userID.(uint),
		PostID:  post.ID,
	}

	if err := DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败"})
		return
	}

	// 查询完整的评论信息（包含用户信息）
	var fullComment Comment
	DB.Preload("User").First(&fullComment, comment.ID)

	c.JSON(http.StatusCreated, fullComment)
}

// 获取文章的评论列表函数
func getCommentsByPostID(c *gin.Context) {
	postID := c.Param("id")

	// 检查文章是否存在
	var post Post
	if err := DB.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	var comments []Comment
	if err := DB.Preload("User").Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "失败"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
