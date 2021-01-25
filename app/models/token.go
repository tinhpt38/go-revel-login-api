package models

type AccessToken struct {
	ID string `json:"_id,omitempty" bson:"_id,omitempty"`
    userID string `json:"user_id" bson:"user_id"`
    token string `json:"token" bson:"token"`
    tokenType int64 `json:"type" bson:"type"`
    expiresIn int64 `json:"expires_in" bson:"expires_in"`
    createdAt string `json:"created_at" bson:"created_at"`
    updatedAt string `json:"updated_at" bson:"updated_at"`
}

