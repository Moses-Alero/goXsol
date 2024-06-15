package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/associated_token_account"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/program/token"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

func CreateMint() string {

	c := client.NewClient(rpc.DevnetRPCEndpoint)
	//create the new int account
	mint := types.NewAccount()

	//get the rent exception amount
	rentExemptBal, err := c.GetMinimumBalanceForRentExemption(context.Background(), token.MintAccountSize)
	if err != nil {
		log.Fatalf(err.Error())
	}
	//get the latest blockhash
	resp, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}
	//specify new tx for creating the mint
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: resp.Blockhash,
			Instructions: []types.Instruction{
				system.CreateAccount(system.CreateAccountParam{
					From:     feePayer.PublicKey,
					New:      mint.PublicKey,
					Owner:    common.TokenProgramID,
					Lamports: rentExemptBal,
					Space:    token.MintAccountSize,
				}),
				token.InitializeMint(token.InitializeMintParam{
					Mint:       mint.PublicKey,
					MintAuth:   feePayer.PublicKey,
					Decimals:   8,
					FreezeAuth: nil,
				}),
			},
		}),
		Signers: []types.Account{feePayer, mint},
	})

	//send tx
	txHash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(txHash)
	return mint.PublicKey.String()
}

func CreateTokenAccount(mintPubKey, userPubKey common.PublicKey) string {
	// create associated token account
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	ata, _, err := common.FindAssociatedTokenAddress(userPubKey, mintPubKey)
	if err != nil {
		log.Fatal("Error finding associated token address: ", err)
	}

	fmt.Println("ata: ", ata.ToBase58())
	//send tx
	resp, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("Error getting latest blockhash: %v", err)
	}

	//create associated token account transaction instruction
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: resp.Blockhash,
			Instructions: []types.Instruction{
				associated_token_account.Create(associated_token_account.CreateParam{
					Funder:                 feePayer.PublicKey,
					Owner:                  userPubKey,
					Mint:                   mintPubKey,
					AssociatedTokenAccount: ata,
				}),
			},
		}),
		Signers: []types.Account{feePayer},
	})

	if err != nil {
		log.Fatalf("Error creating transaction: %v", err)
	}

	//send tx
	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Printf("Error sending transaction: %v", err.Error())
	}

	fmt.Println("Transaction Hash: ", txhash)
	return ata.String()
}
func TransferSol(receiverAddr string) {

	sender := common.PublicKeyFromString("27jwuwuUaVbgsSgjTKznbZ3zQDQtFmphWrhEtedk5SWs")

	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// to fetch recent blockhash
	recentBlockhashResponse, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get recent blockhash, err: %v", err)
	}

	// create a transfer tx
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        sender,
			RecentBlockhash: recentBlockhashResponse.Blockhash,
			Instructions: []types.Instruction{
				system.Transfer(system.TransferParam{
					From:   sender,
					To:     common.PublicKeyFromString(receiverAddr),
					Amount: 1e9, // 0.1 SOL
				}),
			},
		}),
	})
	if err != nil {
		log.Fatalf("failed to create a new transaction, err: %v", err)
	}

	// send tx
	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("failed to send tx, err: %v", err)
	}

	log.Println("txhash:", txhash)
}
