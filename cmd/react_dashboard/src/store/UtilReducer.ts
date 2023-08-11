const utilReducer = (state: UtilState, action: UtilAction): UtilState => {
  switch (action.type) {
    case "ON":
      return {
        isActive: true,
        type: action.payload.type,
        message: action.payload.message,
      };
    case "OFF":
      return {
        ...state,
        isActive: false,
      };
    default:
      return state;
  }
};

export default utilReducer;
