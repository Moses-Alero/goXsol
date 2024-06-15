import {createRoot} from "react-dom/client"

import { BalanceDisplay } from "./components/balance";
import { WalletButton } from "./components/button";

const balance = document.getElementById("balance")
if (!balance) {
    throw new Error('Could not find element with id balance');
}


const balanceReactRoot = createRoot(balance)

balanceReactRoot.render(WalletButton())
