package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pool struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Address     string             `json:"address,omitempty" validate:"required"`
	Symbol      string             `json:"symbol,omitempty" validate:"required"`
	Description string             `json:"description,omitempty" validate:"required"`
	Type        string             `json:"type,omitempty" validate:"required"`
	Token       string             `json:"token,omitempty" validate:"required"`
	Apr         string             `json:"apr,omitempty" validate:"required"`
	Label       string             `json:"label,omitempty" validate:"required"`
	Color       string             `json:"color,omitempty" validate:"required"`
	Info        string             `json:"info,omitempty" validate:"required"`
}
