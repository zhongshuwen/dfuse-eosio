import { Config } from "../models/config"

export const LOGO_PLACEHOLDER =
  "https://raw.githubusercontent.com/eoscafe/eos-airdrops/master/logos/placeholder.png"

export const LOGO_LG_PLACEHOLDER =
  "https://raw.githubusercontent.com/eoscafe/eos-airdrops/master/logos/placeholder-lg.png"

export interface TokenInfo {
  name: string
  logo?: string
  logo_lg?: string
  symbol: string
  account: string
  chain: "bos" | "eos" | "jungle" | "worbli" | "wax" | "snax" | "telos"
}

export function getTokenInfosForNetwork(network: string): TokenInfo[] {
  const chain = networkToName[network]
  if (!chain) {
    return []
  }

  return AIRDROPS.filter((element) => element.chain === chain)
}

export type TokenInfoKey = string

export function getTokenInfoKey(info: TokenInfo): TokenInfoKey {
  return info.account + info.symbol
}

export function getTokenInfosByKeyMap(): Record<TokenInfoKey, TokenInfo> {
  const mappings: ReturnType<typeof getTokenInfosByKeyMap> = {}
  getTokenInfosForNetwork(Config.network_id).forEach((info) => {
    mappings[getTokenInfoKey(info)] = info
  })

  return mappings
}

const networkToName: Record<string, TokenInfo["chain"]> = {
  "eos-mainnet": "eos",
  "eos-jungle": "jungle",
  "eos-worbli": "worbli",
  "wax-mainnet": "wax",
}

const eosCafeList: TokenInfo[] = [
]

// List of tokens that do not actual work correctly anymore
const removedTokens = [
  { chain: "eos", account: "nutscontract" },
  { chain: "eos", account: "uxfyretoken1" },
  { chain: "eos", account: "triviatokens" },
]

function isRemovedToken(tokenInfo: TokenInfo): boolean {
  return removedTokens.some(
    (removedToken) =>
      tokenInfo.chain === removedToken.chain && tokenInfo.account === removedToken.account
  )
}

export const AIRDROPS: TokenInfo[] = [...eosCafeList]
  .filter((element) => !isRemovedToken(element))
  .map((element) => {
    if (element.logo === LOGO_PLACEHOLDER) {
      element.logo = undefined
    }

    if (element.logo_lg === LOGO_LG_PLACEHOLDER) {
      element.logo_lg = undefined
    }

    return element
  })
