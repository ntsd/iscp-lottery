const requiredENV = (key: string): string => {
  const value = process.env[key]
  if (!value) {
    throw Error(`required env ${key} not found`)
  }
  return value
}

export const Configs = {
  seed: process.env["REACT_APP_WALLET_SEED"],
  waspWebSocketUrl: requiredENV("REACT_APP_WASP_WEB_SOCKET_URL"),
  waspApiUrl: requiredENV("REACT_APP_WASP_API_URL"),
  goshimmerApiUrl: requiredENV("REACT_APP_GOSHIMMER_API_URL"),
  chainId: requiredENV("REACT_APP_CHAIN_ID"),
  contractName: requiredENV("REACT_APP_CONTRACT_NAME"),
};
