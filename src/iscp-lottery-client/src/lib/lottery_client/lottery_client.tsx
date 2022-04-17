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

interface LotteryClientProviderProps {
  children: React.ReactNode
}


export const LotteryClientProvider: React.FC<LotteryClientProviderProps> = ({ children }) => {
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

  const handleVmMessage = (message: string[]): void => {
    const messageHandlers: MessageHandlers = {
    };

    const topicIndex = 3;
    const topic = message[topicIndex];

    if (typeof messageHandlers[topic] != "undefined") {
      messageHandlers[topic](topicIndex);
    }
  }

  const handleMessage = (message: MessageEvent<string>): void => {
    const msg = message.data.toString().split("|");

    if (msg.length === 0) {
      return;
    }

    if (msg[0] !== "vmmsg") {
      return;
    }

    handleVmMessage(msg);
  }

  return <>{children}</>;
};
