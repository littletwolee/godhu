package models

type Router struct {
	Post   map[string]Data `json:"post"`
	Get    map[string]Data `json:"get"`
	Put    map[string]Data `json:"put"`
	Delete map[string]Data `json:"delete"`
}

// type Path struct {
// 	Path string `json:"path"`
// 	Data Data   `json:"data"`
// }

type Data struct {
	Data    string `json:"data"`
	Cookie  []KV   `json:"cookie"`
	Session []KV   `json:"session"`
}

type KV struct {
	K string `json:"k"`
	V string `json:"v"`
}
