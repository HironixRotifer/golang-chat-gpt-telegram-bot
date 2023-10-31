package token

import (
	"net/http"
)

type AuthorizationServer struct {
	server *http.Server
	// TokenRepository TokenRepository
	redirectURL string
}

func NewAuthorizationServer(redirectURL string) *AuthorizationServer {
	return &AuthorizationServer{redirectURL: redirectURL}
}

func (s *AuthorizationServer) Start() error {
	s.server = &http.Server{
		Addr:    ":8080",
		Handler: s,
	}

	return s.server.ListenAndServe()
}

func (s *AuthorizationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	chatIDParam := r.URL.Query().Get("chat_id")
	if chatIDParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// chatID, err := strconv.ParseInt(chatIDParam, 10, 64)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// requestToken, err := s.TokenRepository.Get(chatID, RequestTokens)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// log.Printf("chat_id: %d\nrequest_token: %s\n", chatID)

	// err = s.TokenRepository.Save(chatID, requestToken, AccesTokens)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	w.Header().Add("Location", s.redirectURL)
	w.WriteHeader(http.StatusMovedPermanently)
}
