package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	dto "sanberhub-test/dto/result"
	transactionsdto "sanberhub-test/dto/transactions"
	userdto "sanberhub-test/dto/user"
	"sanberhub-test/models"
	"sanberhub-test/repository"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type HandlerUser struct {
	UserRepository repository.UserRepository
}

type database struct {
	db *gorm.DB
}

func UserHandler(userRepository repository.UserRepository) *HandlerUser {
	return &HandlerUser{userRepository}
}

func (h *HandlerUser) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: users,
	})
}

func (h *HandlerUser) Saldo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("no_rekening"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: ("No Rekening Tidak Ada"),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertSaldoResponse(user),
	})
}

func (h *HandlerUser) Mutasi(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("no_rekening"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: ("No Rekening Tidak Ada"),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: user,
	})
}

func (h *HandlerUser) Daftar(c echo.Context) error {

	request := new(userdto.CreateUser)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	rand.Seed(time.Now().UnixNano())
	first := "12"
	second := fmt.Sprintf("%04d", rand.Intn(10000))
	third := fmt.Sprintf("%04d", rand.Intn(10000))

	noRekening := fmt.Sprintf("%s%s%s", first, second, third)
	nominal := 0

	newRekening, err := strconv.Atoi(noRekening)

	user := models.User{
		ID:      newRekening,
		Nama:    request.Nama,
		Nik:     request.Nik,
		NoHp:    request.NoHp,
		Nominal: nominal,
	}

	data, err := h.UserRepository.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertResponse(data),
	})

}

func (h *HandlerUser) Tabung(c echo.Context) error {
	request := new(transactionsdto.CreateTransaction)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	waktu := time.Now()

	parsingWaktu := waktu.Format("2006-01-02 15:04:05")

	request.KodeTransaksi = "C"

	transaction := models.Transaction{
		UserID:        request.UserID,
		Nominal:       request.Nominal,
		Waktu:         parsingWaktu,
		KodeTransaksi: request.KodeTransaksi,
	}

	data, err := h.UserRepository.Transaction(transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: ("No Rekening tidak ditemukan")})
	}

	user, err := h.UserRepository.GetUser(data.UserID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	user.Nominal = data.Nominal + user.Nominal

	newSaldo, err := h.UserRepository.UpdateUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertSaldoResponse(newSaldo)})
}

func (h *HandlerUser) Tarik(c echo.Context) error {
	request := new(transactionsdto.CreateTransaction)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	waktu := time.Now()
	parsingWaktu := waktu.Format("2006-01-02 15:04:05")
	request.KodeTransaksi = "D"

	transaction := models.Transaction{
		UserID:        request.UserID,
		Nominal:       request.Nominal,
		Waktu:         parsingWaktu,
		KodeTransaksi: request.KodeTransaksi,
	}

	data, err := h.UserRepository.Transaction(transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: ("No Rekening tidak ditemukan")})
	}

	user, err := h.UserRepository.GetUser(data.UserID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	if data.Nominal > user.Nominal {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: ("Saldo Tidak Cukup")})
	}

	user.Nominal = user.Nominal - data.Nominal

	newSaldo, err := h.UserRepository.UpdateUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertSaldoResponse(newSaldo)})
}

func convertResponse(u models.User) userdto.UserResponse {
	return userdto.UserResponse{
		ID:   u.ID,
		Nama: u.Nama,
		Nik:  u.Nik,
		NoHp: u.NoHp,
	}
}

func convertSaldoResponse(u models.User) userdto.SaldoResponse {
	return userdto.SaldoResponse{
		Nominal: u.Nominal,
	}
}

func convertTransactionResponse(u models.Transaction) transactionsdto.TransactionResponse {
	return transactionsdto.TransactionResponse{
		UserID:        u.UserID,
		Nominal:       u.Nominal,
		Waktu:         u.Waktu,
		KodeTransaksi: u.KodeTransaksi,
	}
}
