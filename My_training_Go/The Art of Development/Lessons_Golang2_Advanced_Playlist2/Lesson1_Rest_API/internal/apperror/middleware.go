package apperror

import (
	"errors"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *AppError
		err := h(w, r)
		if err != nil {
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrNotFound.Marshal())
					return
				}
				//} else if errors.Is(err, NoAuthErr) {
				//	w.WriteHeader(http.StatusUnauthorized)
				//	w.Write(ErrNotFound.Marshal())
				//	return
				//}
				err = err.(*AppError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(ErrNotFound.Marshal())
			}

			w.WriteHeader(http.StatusTeapot) //не наша ошибка, отдаем как есть. А надо обернуть в систем. ошибку
			w.Write([]byte(err.Error()))
		}
	}
}
