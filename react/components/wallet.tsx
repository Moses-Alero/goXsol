import { FC, ReactNode, useEffect, useMemo, useState } from "react";
import {
  ConnectionProvider,
  WalletProvider,
  useWallet,
  useConnection
} from "@solana/wallet-adapter-react";
import { WalletModalProvider } from "@solana/wallet-adapter-react-ui";
import * as web3 from "@solana/web3.js";
import * as walletAdapterWallets from "@solana/wallet-adapter-wallets";
//  require("@solana/wallet-adapter-react-ui/styles.css");

const WalletStuff = ()=>{
  const [balance, setBalance] = useState(0);
  const {publicKey} = useWallet() 
  const {connection} = useConnection()
  useEffect(() => {
    if (!connection || !publicKey) { return }

    // Ensure the balance updates after the transaction completes
    connection.onAccountChange(
        publicKey, 
        (updatedAccountInfo) => {
            setBalance(updatedAccountInfo.lamports / web3.LAMPORTS_PER_SOL)
        }, 
        'confirmed'
    )
   
    connection.getAccountInfo(publicKey).then(info => {
        setBalance(info.lamports);

        console.log("rererer")
    })
    
}, [connection, publicKey])
  return (
    updateEverything(balance, publicKey),
    <>
    </>
    
  )
}

export const WalletContextProvider: FC<{ children: ReactNode }> = ({ children }) => {
  const endpoint = web3.clusterApiUrl("devnet");
//   const wallets = [new walletAdapterWallets.PhantomWalletAdapter()]
  const wallets = useMemo(() => [], []);
  return (
   <div>
     <ConnectionProvider endpoint={endpoint}>
      <WalletProvider wallets={wallets}>
        <WalletStuff/>
        <WalletModalProvider>{children}</WalletModalProvider>
      </WalletProvider>
    </ConnectionProvider>
   </div>
  );
};

const updateEverything = (balance, publicKey) => { 
  if (publicKey){
    document.getElementById("balance-wallet").innerHTML = (balance/web3.LAMPORTS_PER_SOL).toString()
    document.getElementById("token-acct")["value"] = publicKey.toString()
  }else{
    document.getElementById("balance-wallet").innerHTML = "0"
    document.getElementById("token-acct")["value"] = ""
  }
}