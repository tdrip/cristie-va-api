package auth

type OAuthRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type OAuthResponse struct {
	AccessToken string `json:"access_token,omitempty"`
}

type SessionResponse struct {
	SessionId     string `json:"session_id,omitempty"`
	UserSessionId string `json:"user_session_id,omitempty"`
}
