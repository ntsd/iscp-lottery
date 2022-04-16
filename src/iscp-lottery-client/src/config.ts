const requiredENV = (key: string): string => {
  const value = process.env[key]
  if (!value) {
    throw Error(`required env ${key} not found`)
  }
  return value
}

export const Configs = {
  seed: requiredENV("NEXT_PUBLIC_WALLET_SEED"),
  waspWebSocketUrl: requiredENV("NEXT_PUBLIC_WASP_WEB_SOCKET_URL"),
  waspApiUrl: requiredENV("NEXT_PUBLIC_WASP_API_URL"),
  goshimmerApiUrl: requiredENV("NEXT_PUBLIC_GOSHIMMER_API_URL"),
  chainId: requiredENV("NEXT_PUBLIC_CHAIN_ID"),
  contractName: requiredENV("NEXT_PUBLIC_CONTRACT_NAME"),
};
