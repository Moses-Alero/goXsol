import { useConnection, useWallet } from '@solana/wallet-adapter-react';
import { LAMPORTS_PER_SOL } from '@solana/web3.js';
import {FC, useEffect, useState } from 'react'
import { WalletContextProvider } from './wallet';
import { WalletMultiButton } from '@solana/wallet-adapter-react-ui';

const Balance: FC = () => {
    const [balance, setBalance] = useState(0);
    const { connection } = useConnection();
    const { publicKey } = useWallet();
    useEffect(() => {
        if (!connection || !publicKey) { return }

        // Ensure the balance updates after the transaction completes
        connection.onAccountChange(
            publicKey, 
            (updatedAccountInfo) => {
                setBalance(updatedAccountInfo.lamports / LAMPORTS_PER_SOL)
            }, 
            'confirmed'
        )
       
        connection.getAccountInfo(publicKey).then(info => {
            setBalance(info.lamports);
            document.getElementById("balance-wallet").innerHTML = (balance/LAMPORTS_PER_SOL).toString()
            console.log("rererer")
        })
    }, [connection, publicKey])

    return (
        <div>
            <h3>{publicKey ? 
               <>
                {`Balance: `}<span>{balance/LAMPORTS_PER_SOL}</span>
               </> :
             'Connect wallet'}</h3>
        </div>
    )
}

export const BalanceDisplay = () => {
    return (
        <div className="balance">
            <WalletContextProvider>
                <WalletMultiButton> 
                    <Balance/>
                </WalletMultiButton>
            </WalletContextProvider>
        </div>
    )
}