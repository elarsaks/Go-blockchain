import React, { createContext, useState } from "react";
import { fetchUserWalletDetails, fetchWalletBalance } from "api/wallet";
import { fetchMinerWalletDetails } from "api/miner";

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
  util: {
    isActive: false,
    type: "info",
    message: "",
  },
  setUtil: () => {},
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

  const [util, setUtil] = useState<UtilState>({
    isActive: false,
    type: "info",
    message: "",
  });

  return (
    <WalletContext.Provider
      value={{
        minerWallet,
        userWallet,
        util,
        setUtil,
        setUserWallet,
        setMinerWallet,
      }}
    >
      {children}
    </WalletContext.Provider>
  );
};
