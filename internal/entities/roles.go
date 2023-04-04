package entities

type Roles struct {
	Id          int    `json:"id"`
	RoleName    string `json:"rolename"`
	Description string `json:"description"`
}

type Role struct {
	Name        string `json:"name,omitempty" `
	Description string `json:"description,omitempty" `
}
