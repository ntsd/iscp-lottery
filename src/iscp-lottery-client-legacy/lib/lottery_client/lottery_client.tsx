import { Configs } from "../../config";
import {
  BasicClient,
  Buffer,
  Colors,
  IKeyPair,
  IOffLedger,
  IOnLedger,
  OffLedger,
  WalletService,
} from "../wasp_client";
import { createNanoEvents, Emitter, Unsubscribe } from "nanoevents";
import { HName } from "../wasp_client/crypto/hname";
import { useState } from "react";
import useWebSocket from "react-use-websocket";

type MessageHandlers = { [key: string]: (index: number) => void };
type ParameterResult = { [key: string]: Buffer };

export const LotteryClientProvider: React.FC = ({ children }) => {
  const scName: string = Configs.contractName;
  const scHName: string = HName.HashAsString(scName);
  const webSocketUrl = Configs.waspWebSocketUrl.replace(
    "%chainId",
    Configs.chainId
  );

  // Public websocket api that will echo messages sent to it back to the client
  const [socketUrl, setSocketUrl] = useState(webSocketUrl);

  // Initialize websocket
  const { sendMessage, lastMessage, readyState } = useWebSocket(socketUrl);

  return <>{children}</>;
};
