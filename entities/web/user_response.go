/*
 * ProjectMart
 *
 * an API for simple E-Commerce application
 *
 * API version: v1.0.1
 * Written by: Badar Wildani
 */

package web

import "time"

type UserResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Username string `json:"username"`
	Gender string `json:"gender"`
	Address string `json:"address"`
	Avatar string `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}