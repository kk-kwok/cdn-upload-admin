package models

// DEFINES TOKEN SECRET KEY
const (
	TokenSecretKey = "ZWFmMGI0YTJhOTlhMzJmYzYxNTc4Zjg1ODA2MTE0MzkgIC0K"
)

// Response : DEFINES RETURN JSON STRUCT
type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}

// UserCredentials : DEFINES USERCREDENTIALS JSON STRUCT
type UserCredentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// UserAuthRes : DEFINES USER AUTH RESPONSE
type UserAuthRes struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"username"`
	IsAdmin  int    `json:"is_admin"`
	Token    string `json:"token"`
}

// User : DEFINES USER JSON STRUCT
type User struct {
	ID          int    `json:"id"`
	IsDelete    int    `json:"is_delete"`
	UserName    string `json:"username"`
	OldPassword string `json:"old_password"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	IsAdmin     int    `json:"is_admin"`
	Status      string `json:"status"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

// Project : DEFINES PROJECT JSON STRUCT
type Project struct {
	ID           int    `json:"id"`
	IsDelete     int    `json:"is_delete"`
	Name         string `json:"name"`
	PlatformID   int    `json:"platform_id"`
	PlatformName string `json:"platform_name"`
	Path         string `json:"path"`
	Domain       string `json:"domain"`
	CdnID        int    `json:"cdn_id"`
	Status       string `json:"status"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
}

// Platform : DEFINES PLATFORM JSON STRUCT
type Platform struct {
	ID         int    `json:"id"`
	IsDelete   int    `json:"is_delete"`
	Name       string `json:"name"`
	SecretID   string `json:"secret_id"`
	SecretKey  string `json:"secret_key"`
	APIUrl     string `json:"api_url"`
	Action     string `json:"action"`
	Status     string `json:"status"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// FileSuffix : DEFINES FILE SUFFIX JSON STRUCT
type FileSuffix struct {
	ID         int    `json:"id"`
	IsDelete   int    `json:"is_delete"`
	FileSuffix string `json:"file_suffix"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// CDNFile : DEFINES CDN FILE JSON STRUCT
type CDNFile struct {
	ID          int    `json:"id"`
	IsDelete    int    `json:"is_delete"`
	ProjectID   int    `json:"project_id"`
	ProjectName string `json:"project_name"`
	UserID      int    `json:"user_id"`
	UserName    string `json:"username"`
	Path        string `json:"path"`
	Domain      string `json:"domain"`
	FileName    string `json:"file_name"`
	FileMD5     string `json:"file_md5"`
	FileSize    string `json:"file_size"`
	Comment     string `json:"comment"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

// CDNPush : DEFINES CDN PUSH JSON STRUCT
type CDNPush struct {
	CDNName   string `json:"cdn_name"`
	CDNID     int    `json:"cdn_id"`
	Domain    string `json:"domain"`
	Path      string `json:"path"`
	FileName  string `json:"file_name"`
	FileSize  string `json:"file_size"`
	SecretID  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
	APIURL    string `json:"api_url"`
	Action    string `json:"action"`
}
