package structs

// Users is a representation of a post
type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Risk_profiles struct {
	ID      int     `json:"id"`
	Id_user int     `json:"id_user"`
	MM      float32 `json:"MM"`
	Bond    float32 `json:"Bond"`
	Stock   float32 `json:"Stock"`
}

// Result is an array of post
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
