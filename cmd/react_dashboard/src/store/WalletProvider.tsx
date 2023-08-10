import React, { createContext, useState, useEffect } from "react";
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
    util: {
      isActive: false,
      type: "info",
      message: "",
    },
  },
  setMinerWallet: () => {},
  userWallet: {
    amount: "",
    balance: "0.00",
    blockchainAddress: "",
    privateKey: "",
    publicKey: "",
    recipientAddress: "",
    util: {
      isActive: false,
      type: "info",
      message: "",
    },
  },
  setUserWallet: () => {},
});

// For some reason, this is not working
interface WalletProviderProps {
  children: React.ReactNode;
}

export const WalletProvider: React.FC<WalletProviderProps> = ({ children }) => {
  const [minerWallet, setMinerWallet] = useState<StoreWallet>({
    amount: "",
    balance: "0.00",
    blockchainAddress: "",
    privateKey: "",
    publicKey: "",
    recipientAddress: "",
    util: {
      isActive: false,
      type: "info",
      message: "",
    },
  });

  const [userWallet, setUserWallet] = useState<StoreWallet>({
    amount: "",
    balance: "0.00",
    blockchainAddress: "",
    privateKey: "",
    publicKey: "",
    recipientAddress: "",
    util: {
      isActive: false,
      type: "info",
      message: "",
    },
  });

  // Fetch wallet details
  useEffect(() => {
    // Fetch miner wallet details
    fetchMinerWalletDetails("1")
      .then((minerDetails) => {
        setMinerWallet((prevDetails) => ({
          ...prevDetails,
          ...minerDetails,
        }));
      })
      .catch((error) => {});

    // Fetch user wallet details
    fetchUserWalletDetails()
      .then((userDetails) => {
        setUserWallet((prevDetails) => ({
          ...prevDetails,
          ...userDetails,
        }));
      })
      .catch((error) => {});
  }, []);

  // Fetch wallet balance
  useEffect(() => {
    // Fetch miner wallet balance
    fetchWalletBalance(minerWallet.blockchainAddress)
      .then((minerBalance) => {
        setMinerWallet((prevDetails) => ({
          ...prevDetails,
          balance: minerBalance,
        }));
      })
      .catch((error) => {});

    // Fetch user wallet balance
    fetchWalletBalance(userWallet.blockchainAddress)
      .then((userBalance) => {
        setUserWallet((prevDetails) => ({
          ...prevDetails,
          balance: userBalance,
        }));
      })
      .catch((error) => {});
  }, [minerWallet.blockchainAddress, userWallet.blockchainAddress]);

  return (
    <WalletContext.Provider
      value={{
        minerWallet,
        userWallet,
        setUserWallet,
        setMinerWallet,
      }}
    >
      {children}
    </WalletContext.Provider>
  );
};
