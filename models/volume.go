package models

type Volume struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Status     string   `json:"status"`
	AttachedTo []string `json:"attachedTo"`
	Size       int      `json:"size"`
	Region     string   `json:"region"`
}
