import {
  addressStateAtom,
  keyPairStateAtom,
  seedStateAtom,
  addressIndexStateAtom,
  balanceStateAtom,
  walletServiceStateAtom,
  clientStateAtom,
} from "../state/recoil";
import { useRecoilState } from "recoil";
import { Configs } from "../config";
import { Base58, BasicClient, Colors, WalletService } from "./wasp_client";
import { Seed } from "./wasp_client/crypto/seed";
import { LotteryClientProvider } from "./lottery_client/lottery_client";
import { useEffect } from "react";

interface WaspProviderProps {
  children: React.ReactNode
}

export const WaspProvider: React.FC<WaspProviderProps> = (props) => {
  const [clientState, setClientState] = useRecoilState(clientStateAtom);
  const [walletServiceState, setWalletServiceState] = useRecoilState(
    walletServiceStateAtom
  );
  const [seedState, setSeedState] = useRecoilState(seedStateAtom);
  const [addressIndexState, setAddressIndexState] = useRecoilState(
    addressIndexStateAtom
  );
  const [addressState, setAddressState] = useRecoilState(addressStateAtom);
  const [keyPairState, setKeyPairState] = useRecoilState(keyPairStateAtom);
  const [balanceState, setBalanceState] = useRecoilState(balanceStateAtom);

  useEffect(() => {
    // set seed
    let seed;
    if (Configs.seed) {
      // use seed from env
      seed = Base58.decode(Configs.seed);
    } else {
      // random seed
      seed = Seed.generate();
    }
    setSeedState(seed);

    // create wasp client
    const client = new BasicClient({
      GoShimmerAPIUrl: Configs.goshimmerApiUrl,
      WaspAPIUrl: Configs.waspApiUrl,
      SeedUnsafe: seed,
    });
    setClientState(client);

    // create wallet service
    const walletService = new WalletService(client);
    setWalletServiceState(walletService);

    // set the address and key pair
    const addressIndex = 0;
    setAddressIndexState(addressIndex);

    const address = Seed.generateAddress(seed, addressIndex);
    const keyPair = Seed.generateKeyPair(seed, addressIndex);
    setAddressState(address);
    setKeyPairState(keyPair);

    // update fund
    walletService
      .getFunds(address, Colors.IOTA_COLOR_STRING)
      .then((newBalance) => {
        setBalanceState(newBalance);
      });

    // const lotteryClient = new LotteryClient(client, Configs.chainId);
    // lotteryClient.getCurrentRound().then((round) => {
    //   console.log(round);
    // });
  }, [setAddressIndexState, setAddressState, setBalanceState, setClientState, setKeyPairState, setSeedState, setWalletServiceState])

  return <LotteryClientProvider>{props.children}</LotteryClientProvider>;
};
