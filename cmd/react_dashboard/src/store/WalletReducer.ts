type WalletAction =
  | { type: "SET_WALLET"; payload: Partial<StoreWallet> }
  | { type: "SET_UTIL"; payload: UtilState };

export default (state: StoreWallet, action: WalletAction): StoreWallet => {
  switch (action.type) {
    case "SET_WALLET":
      return {
        ...state,
        ...action.payload,
      };
    case "SET_UTIL":
      return {
        ...state,
        util: action.payload,
      };
    default:
      return state;
  }
};
