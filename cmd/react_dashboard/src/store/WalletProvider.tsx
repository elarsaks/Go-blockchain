import React, { createContext, useReducer, useEffect } from "react";
import { fetchUserWalletDetails, fetchWalletBalance } from "api/wallet";
import { fetchMinerWalletDetails } from "api/miner";
import WalletReducer from "store/WalletReducer";

const initialState: StoreWallet = {
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
};

export const WalletContext = createContext({
  minerWallet: initialState,
  userWallet: initialState,
  setUserWallet: (wallet: Partial<StoreWallet>) => {},
  setMinerWallet: (wallet: Partial<StoreWallet>) => {},
});

interface WalletProviderProps {
  children: React.ReactNode;
  previousHash?: string;
}

export const WalletProvider: React.FC<WalletProviderProps> = ({
  children,
  previousHash,
}) => {
  const [minerWallet, dispatchMinerWallet] = useReducer(
    WalletReducer,
    initialState
  );
  const [userWallet, dispatchUserWallet] = useReducer(
    WalletReducer,
    initialState
  );

  function getUserWallet() {
    dispatchUserWallet({
      type: "SET_WALLET_UTIL",
      payload: {
        isActive: true,
        type: "info",
        message: "Fetching user wallet details",
      },
    });

    fetchUserWalletDetails()
      .then((userDetails) => {
        dispatchUserWallet({ type: "SET_WALLET", payload: userDetails });
        dispatchUserWallet({
          type: "SET_WALLET_UTIL",
          payload: {
            isActive: false,
            type: "info",
            message: "",
          },
        });
      })
      .catch((error) => {
        dispatchUserWallet({
          type: "SET_WALLET_UTIL",
          payload: {
            isActive: true,
            type: "error",
            message: "Failed to fetch user wallet details",
          },
        });
      });
  }

  function getMinerWallet() {
    dispatchMinerWallet({
      type: "SET_WALLET_UTIL",
      payload: {
        isActive: true,
        type: "info",
        message: "Fetching miner wallet details",
      },
    });
    fetchMinerWalletDetails("1")
      .then((minerDetails) => {
        dispatchMinerWallet({ type: "SET_WALLET", payload: minerDetails });
        dispatchMinerWallet({
          type: "SET_WALLET_UTIL",
          payload: {
            isActive: false,
            type: "info",
            message: "",
          },
        });
      })
      .catch((error) => {
        dispatchMinerWallet({
          type: "SET_WALLET_UTIL",
          payload: {
            isActive: true,
            type: "error",
            message: "Failed to fetch miner wallet details",
          },
        });
      });
  }

  // Fetch wallet details
  useEffect(() => {
    getMinerWallet();
    getUserWallet();
  }, []);

  // Fetch wallet balance
  useEffect(() => {
    fetchWalletBalance(minerWallet.blockchainAddress)
      .then((minerBalance) => {
        dispatchMinerWallet({
          type: "SET_WALLET",
          payload: { balance: minerBalance },
        });
      })
      .catch((error) => {
        // Handle error
      });

    fetchWalletBalance(userWallet.blockchainAddress)
      .then((userBalance) => {
        dispatchUserWallet({
          type: "SET_WALLET",
          payload: { balance: userBalance },
        });
      })
      .catch((error) => {
        // Handle error
      });
  }, [
    minerWallet.blockchainAddress,
    userWallet.blockchainAddress,
    previousHash,
  ]);

  return (
    <WalletContext.Provider
      value={{
        minerWallet,
        userWallet,
        setUserWallet: (wallet: Partial<StoreWallet>) =>
          dispatchUserWallet({ type: "SET_WALLET", payload: wallet }),
        setMinerWallet: (wallet: Partial<StoreWallet>) =>
          dispatchMinerWallet({ type: "SET_WALLET", payload: wallet }),
      }}
    >
      {children}
    </WalletContext.Provider>
  );
};

export default WalletProvider;
