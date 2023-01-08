package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/obasootom/langtranslator/db/sqlc"
)

type Profile struct {
	Name           string `json:"name" form:"name" binding:"required"`
	Gender         string `json:"gender" form:"gender" binding:"required,oneof=female male"`
	Email          string `json:"email" form:"email" binding:"required"`
	PhoneNumber    string `json:"phonenumber" form:"phonenumber" binding:"required"`
	Country        string `json:"country" form:"country" binding:"required"`
	NativeLanguage string `json:"nativelanguage" form:"nativelanguage" binding:"required"`
	AddressLine    string `jso:"addressline" form:"addressline" binding:"required"`
}

func (server *Server) createProfile(ctx *gin.Context) {
	var req Profile
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateProfileParams{
		Name:           req.Name,
		Gender:         req.Gender,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		AddressLine:    req.AddressLine,
		Country:        req.Country,
		NativeLanguage: req.NativeLanguage,
	}

	profile, err := server.store.CreateProfile(ctx, arg)
	if err != nil {
		if pkErr, ok := err.(*pq.Error); ok {
			switch pkErr.Code.Name() {
			case "UNIQUE_VIOLATION":
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	ctx.JSON(http.StatusOK, profile)
}
