package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/repository"
)

type AuthorizationServer struct {
	server          *http.Server
	TokenRepository repository.TokenRepository
	redirectURL     string
}

func NewAuthorizationServer(tokenRepository repository.TokenRepository, redirectURL string) *AuthorizationServer {
	return &AuthorizationServer{TokenRepository: tokenRepository, redirectURL: redirectURL}
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

	chatID, err := strconv.ParseInt(chatIDParam, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestToken, err := s.TokenRepository.Get(chatID, repository.RequestTokens)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("chat_id: %d\nrequest_token: %s\n", chatID, requestToken)

	// err = s.TokenRepository.Save(chatID, repository.AccesTokens, requestToken)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	w.Header().Add("Location", s.redirectURL)
	w.WriteHeader(http.StatusMovedPermanently)
}
