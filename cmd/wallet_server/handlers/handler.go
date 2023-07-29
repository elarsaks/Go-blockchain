package handlers

type WalletServer interface {
	Port() uint16
	Gateway() string
}

type WalletServerHandler struct {
	server WalletServer
}

func NewWalletServerHandler(s WalletServer) *WalletServerHandler {
	return &WalletServerHandler{server: s}
}
