/*
 * ProjectMart
 *
 * an API for simple E-Commerce application
 *
 * API version: v1.0.1
 * Written by: Badar Wildani
 */

package web

type UserRequest struct {
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Gender string `json:"gender" form:"gender"`
	Address string `json:"address" form:"address"`
	Avatar string `json:"avatar" form:"avatar"`
}