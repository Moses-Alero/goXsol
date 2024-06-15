import {WalletContextProvider} from "./wallet";
import { WalletMultiButton } from "@solana/wallet-adapter-react-ui";
import { PingButton } from "./ping";
import { BalanceDisplay } from "./balance";

export const WalletButton = () => {
  return (
    <WalletContextProvider>
      <WalletMultiButton />
    </WalletContextProvider>
  );
}