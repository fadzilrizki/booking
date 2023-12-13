package model

type InsertBundleBook struct {
	Genre string `json:"name"`
	Books []Book `json:"works"`
}

type Book struct {
	Key     string    `json:"key"`
	Title   string    `json:"title"`
	Edition int64     `json:"edition_count"`
	Author  []*Author `json:"authors"`
}

// Book Author Schema
type Author struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}
