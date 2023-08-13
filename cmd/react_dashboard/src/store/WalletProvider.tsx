import { fetchUserWalletDetails, fetchWalletBalance } from "api/wallet";
import { fetchMinerWalletDetails } from "api/miner";
import WalletReducer from "store/WalletReducer";
import React, {
  createContext,
  useCallback,
  useEffect,
  useReducer,
} from "react";

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
  setUserWalletUtil: (util: UtilState) => {},
  setMinerWalletUtil: (util: UtilState) => {},
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
        message:
          "Registering the user wallet on the blockchain. This process can take up to 28 seconds.",
      },
    });

    fetchUserWalletDetails()
      .then((userDetails) => {
        dispatchUserWallet({ type: "SET_WALLET", payload: userDetails });
        dispatchMinerWallet({
          type: "SET_WALLET",
          payload: {
            recipientAddress: userDetails.blockchainAddress,
          },
        });

        // dispatchUserWallet({
        //   type: "SET_WALLET_UTIL",
        //   payload: {
        //     isActive: false,
        //     type: "info",
        //     message: "",
        //   },
        // });
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
        dispatchUserWallet({
          type: "SET_WALLET",
          payload: {
            recipientAddress: minerDetails.blockchainAddress,
          },
        });

        // dispatchMinerWallet({
        //   type: "SET_WALLET_UTIL",
        //   payload: {
        //     isActive: false,
        //     type: "info",
        //     message: "",
        //   },
        // });
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

  const getUserWalletWalletBalance = useCallback(() => {
    fetchWalletBalance(userWallet.blockchainAddress)
      .then((userBalance) => {
        dispatchUserWallet({
          type: "SET_WALLET",
          payload: { balance: userBalance },
        });
        // dispatchUserWallet({
        //   type: "SET_WALLET_UTIL",
        //   payload: {
        //     isActive: false,
        //     type: "info",
        //     message: "",
        //   },
        // });
      })
      .catch((error) =>
        dispatchUserWallet({
          type: "SET_WALLET_UTIL",
          payload: {
            isActive: true,
            type: "error",
            message: "Failed to fetch user wallet details",
          },
        })
      );
  }, [userWallet.blockchainAddress]);

  const getMinerWalletWalletBalance = useCallback(() => {
    fetchWalletBalance(minerWallet.blockchainAddress)
      .then((minerBalance) => {
        dispatchMinerWallet({
          type: "SET_WALLET",
          payload: { balance: minerBalance },
        });
        // dispatchMinerWallet({
        //   type: "SET_WALLET_UTIL",
        //   payload: {
        //     isActive: false,
        //     type: "info",
        //     message: "",
        //   },
        // });
      })
      .catch((error) =>
        dispatchMinerWallet({
          type: "SET_WALLET_UTIL",
          payload: {
            isActive: true,
            type: "error",
            message: "Failed to fetch miner wallet details", // Fixed the message to refer to miner instead of user
          },
        })
      );
  }, [minerWallet.blockchainAddress]);

  // Fetch wallet details
  useEffect(() => {
    getMinerWallet();
    getUserWallet();
  }, []);

  // Fetch wallet balance
  useEffect(() => {
    if (minerWallet.blockchainAddress) getMinerWalletWalletBalance();
    if (userWallet.blockchainAddress) getUserWalletWalletBalance();
  }, [
    minerWallet.blockchainAddress,
    userWallet.blockchainAddress,
    getMinerWalletWalletBalance,
    getUserWalletWalletBalance,
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
        setUserWalletUtil: (util: UtilState) =>
          dispatchUserWallet({ type: "SET_WALLET_UTIL", payload: util }),
        setMinerWalletUtil: (util: UtilState) =>
          dispatchMinerWallet({ type: "SET_WALLET_UTIL", payload: util }),
      }}
    >
      {children}
    </WalletContext.Provider>
  );
};

export default WalletProvider;
