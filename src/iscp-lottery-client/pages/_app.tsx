import "../styles/globals.css";
import type { AppProps } from "next/app";
import { WaspProvider } from "../lib/wasp_provider";
import { RecoilRoot } from "recoil";

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <RecoilRoot>
      <WaspProvider>
        <Component {...pageProps} />
      </WaspProvider>
    </RecoilRoot>
  );
}

export default MyApp;
