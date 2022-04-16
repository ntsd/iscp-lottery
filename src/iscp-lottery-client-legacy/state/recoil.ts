import { IKeyPair } from "./../lib/wasp_client/models/IKeyPair";
import { atom } from "recoil";
import { BasicClient, Buffer, WalletService } from "../lib/wasp_client";

export const seedStateAtom = atom<Buffer>({
  key: "seedState",
  default: undefined,
});

export const addressIndexStateAtom = atom<number>({
  key: "addressIndexState",
  default: 0,
});

export const addressStateAtom = atom<string>({
  key: "addressState",
  default: "",
});

export const keyPairStateAtom = atom<IKeyPair>({
  key: "keyPairstate",
  default: undefined,
});

export const balanceStateAtom = atom<bigint>({
  key: "balanceState",
  default: 0n,
});

export const clientStateAtom = atom<BasicClient>({
  key: "clientState",
  default: undefined,
});

export const walletServiceStateAtom = atom<WalletService>({
  key: "walletServiceState",
  default: undefined,
});
