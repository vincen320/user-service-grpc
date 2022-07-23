package helper

import (
	"errors"
	"log"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/vincen320/user-service-grpc/exception"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ReturnProtoError(e error) error {
	var (
		badRequest    *exception.BadRequestError
		notFound      *exception.NotFoundError
		conflict      *exception.ConflictError
		validationErr validator.ValidationErrors
	)

	log.Println(reflect.TypeOf(e))

	if errors.As(e, &badRequest) {
		return status.Errorf(codes.InvalidArgument, "bad request "+badRequest.Error())
	}

	if errors.As(e, &notFound) {
		return status.Errorf(codes.NotFound, notFound.Error())
	}

	if errors.As(e, &conflict) {
		return status.Errorf(codes.AlreadyExists, conflict.Error())
	}

	// KALAU ERROR TIDAK BISA ASSERTION  ATAU ERROR.AS SEBAGAI validator.ValidationErrors , coba lihat package yang diimport
	// APAKAH ITU v10 atau bukan, contoh
	// "github.com/go-playground/validator/v10" >> benar
	// "github.com/go-playground/validator" >>salah
	if errors.As(e, &validationErr) {
		//return status.Errorf(codes.InvalidArgument, validationErrorMessage(&validationErr))
		return status.Errorf(codes.InvalidArgument, validationErrorMessageSemua(&validationErr))
	}
	return status.Errorf(codes.Unknown, "internal server error")
}

func validationErrorMessage(errVal *validator.ValidationErrors) string {
	//Custom Error
	//Ambil Salah satu saja
	switch (*errVal)[0].Tag() {
	//Cek errornya
	case "required":
		//Disini bisa isi pesannya
		//penjelasan (*slice) = https://stackoverflow.com/questions/38468258/why-is-indexing-on-the-slice-pointer-not-allowed-in-golang
		return (*errVal)[0].Field() + " harus diisi"
	case "min":
		return (*errVal)[0].Field() + " minimal " + (*errVal)[0].Param() + " karakter"
	case "max":
		return (*errVal)[0].Field() + " maksimal " + (*errVal)[0].Param() + " karakter"
	case "email":
		return (*errVal)[0].Value().(string) + "bukan email yang valid"
	default:
		return "validasi error"
	}
}

//coba ambil semua validasinya, sama aja seperti diatas, cuma kalau ini kyk ambil semua
func validationErrorMessageSemua(errVal *validator.ValidationErrors) string {
	var message string
	for _, err := range *errVal {
		//untuk kebutuhan :ENTER TIAP BARIS SAAT SUDAH ADA SATU KALIMAT DAN MASIH ADA ERROR(DALAM LOOP), KALAU BELUM ADA KALIMAT JANGAN DIENTER DULU
		if message != "" {
			message += "\n"
		}

		switch err.Tag() {
		//Cek errornya
		case "required":
			message += err.Field() + " harus diisi"
		case "min":
			message += err.Field() + " minimal " + err.Param() + " karakter"
		case "max":
			message += err.Field() + " maksimal " + err.Param() + " karakter"
		case "email":
			message += err.Value().(string) + "bukan email yang valid"
		}
	}
	return message
}
