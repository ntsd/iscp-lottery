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
import { BasicClient, Colors, WalletService } from "./wasp_client";
import { Seed } from "./wasp_client/crypto/seed";
import { LotteryClientProvider } from "./lottery_client/lottery_client";

export const WaspProvider: React.FC = (props) => {
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

  const seed = Seed.generate();
  setSeedState(seed);
  const addressIndex = 0;
  setAddressIndexState(addressIndex);

  const client = new BasicClient({
    GoShimmerAPIUrl: Configs.goshimmerApiUrl,
    WaspAPIUrl: Configs.waspApiUrl,
    SeedUnsafe: seed,
  });
  setClientState(client);
  const walletService = new WalletService(client);
  setWalletServiceState(walletService);

  // set the address
  setAddressState(Seed.generateAddress(seed, addressIndex));
  setKeyPairState(Seed.generateKeyPair(seed, addressIndex));

  // update fund
  let _balance = 0n;
  try {
    walletService
      .getFunds(addressState, Colors.IOTA_COLOR_STRING)
      .then((newBalance) => {
        _balance = newBalance;
      });
  } finally {
    setBalanceState(_balance);
  }

  // const lotteryClient = new LotteryClient(client, Configs.chainId);
  // lotteryClient.getCurrentRound().then((round) => {
  //   console.log(round);
  // });

  return <LotteryClientProvider>{props.children}</LotteryClientProvider>;
};
