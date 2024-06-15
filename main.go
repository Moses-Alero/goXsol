package main

import (
	"fmt"
	"goxsol/components"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/a-h/templ"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/joho/godotenv"
)

var feePayerKey = os.Getenv("FEE_PAYER")
var feePayer, _ = types.AccountFromBase58(os.Getenv(feePayerKey))

func main() {
	godotenv.Load()
	component := components.Base()

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	mux.Handle("/create-mint", http.HandlerFunc(createMintHandler))
	mux.Handle("/create-token-account", http.HandlerFunc(createTokenAccountHandler))
	mux.Handle("/", templ.Handler(component))

	fmt.Println("server running on http://localhost:3000")
	http.ListenAndServe(":3000", mux)

	//component.Render(context.Background(), os.Stdout)
}

func createMintHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	mintAddr := CreateMint()
	fmt.Println(mintAddr)

	components.Mint(mintAddr).Render(r.Context(), w)

	fmt.Println(values)
}

func createTokenAccountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tokenMint := values["token-mint"][0]
	tokenAccount := values["token-acct-owner"][0]
	fmt.Println(tokenMint)
	fmt.Println(tokenAccount)
	// ata := CreateTokenAccount(common.PublicKeyFromString(tokenMint), common.PublicKeyFromString(tokenAccount))

	components.CreateToken(tokenMint, "Acct-example").Render(r.Context(), w)
}
