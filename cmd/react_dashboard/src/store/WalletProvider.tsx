import React, { createContext, useState } from "react";

export const WalletContext = createContext<WalletStore>({
  minerWallet: {
    amount: "",
    balance: "0.00",
    blockchainAddress: "",
    privateKey: "",
    publicKey: "",
    recipientAddress: "",
  },
  userWallet: {
    amount: "",
    balance: "0.00",
    blockchainAddress: "",
    privateKey: "",
    publicKey: "",
    recipientAddress: "",
  },
  setMinerWallet: () => {},
  setUserWallet: () => {},
});

// For some reason, this is not working
interface WalletProviderProps {
  children: React.ReactNode;
}

export const WalletProvider: React.FC<WalletProviderProps> = ({ children }) => {
  const [minerWallet, setMinerWallet] = useState<WalletState>({
    amount: "",
    balance: "0.00",
    blockchainAddress: "",
    privateKey: "",
    publicKey: "",
    recipientAddress: "",
  });

  const [userWallet, setUserWallet] = useState<WalletState>({
    amount: "",
    balance: "0.00",
    blockchainAddress: "",
    privateKey: "",
    publicKey: "",
    recipientAddress: "",
  });

  return (
    <WalletContext.Provider
      value={{ minerWallet, setMinerWallet, userWallet, setUserWallet }}
    >
      {children}
    </WalletContext.Provider>
  );
};
