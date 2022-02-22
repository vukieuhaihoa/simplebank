package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vukieuhaihoa/simplebank/db/sqlc"
)

type createTransferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req createTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	if !server.validAccount(ctx, req.FromAccountID, req.Currency) {
		return 	
	}


	if !server.validAccount(ctx, req.ToAccountID, req.Currency) {
		return 	
	}

	result, err := server.store.TransferTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	acccount, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}
	if acccount.Currency != currency {
		msg := fmt.Errorf("acconnt [%d] currency %s mismath %s", accountID, acccount.Currency, currency)
		ctx.JSON(http.StatusInternalServerError, errorResponse(msg))
		return false
	}
	return true
}
