package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/obasootom/langtranslator/db/sqlc"
)

type Order struct {
	SourceLanguage           string `json:"source_language" form:"source_language" binding:"required"`
	TargetLanguage           string `json:"target_language" form:"target_language" binding:"required"`
	Translator               string `json:"translator" form:"translator" binding:"required"`
	ProofReader              string `json:"proof_reader" form:"proof_reader" binding:"required"`
	TranslationDelivaryDate  string `json:"translation_delivary_date" form:"translation_delivary_date" binding:"required"`
	ProofReadingDelivaryDate string `json:"proof_reading_delivary_date" form:"proof_reading_delivary_date" binding:"required"`
	ProjectEndDate           string `json:"project_end_date" form:"project_end_date" binding:"required"`
	ServiceLevel             string `json:"service_level" form:"service_level" binding:"required"`
	Profession               string `json:"profession" form:"profession" binding:"required"`
	TranslatorCategory       string `form:"translator_category" json:"translator_category" binding:"required"`
	DelivarySpeed            string `form:"delivary_speed" json:"delivary_speed" binding:"required"`
	TranslatorRequest        string `form:"translator_request" json:"translator_request" binding:"required"`
	DelivaryAddress          string `form:"delivary_address" json:"delivary_address" binding:"required"`
}

func (server *Server) createOrder(ctx *gin.Context) {
	var req Order
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	arg := db.CreateOrdersParams{
		SourceLanguage:           req.SourceLanguage,
		TargetLanguage:           req.TargetLanguage,
		Translator:               req.Translator,
		ProofReader:              req.ProofReader,
		TranslationDelivaryDate:  req.TranslationDelivaryDate,
		ProofReadingDelivaryDate: req.ProofReadingDelivaryDate,
		ProjectEndDate:           req.ProjectEndDate,
		ServiceLevel:             req.ServiceLevel,
		Profession:               req.Profession,
		TranslatorCategory:       req.TranslatorCategory,
		DelivarySpeed:            req.DelivarySpeed,
		TranslatorRequest:        req.TranslatorRequest,
		DelivaryAddress:          req.DelivaryAddress,
	}
	order, err := server.store.CreateOrders(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, order)
}
