# Morpheus Lumerin Node
The purpose of this software is to enable interaction with distributed, decentralized LLMs on the Morpheus network through a desktop chat experience.

The steps listed below are for both the Consumer and Provider to get started with the Morpheus Lumerin Node. As the software is developed, both onboarding & configuration of the provider and consumer roles will be simplified, automated and more transparent to the end user.

# **NOTE: ARBITRUM SEPOLIA TESTNET ONLY at this time - DEVELOPER PREVIEW ONLY**

**Components that are included in this repository are:**
* Lumerin Proxy-Router is a background process that monitors sepcific blockchain contract events, manages secure sessions between consumers and providers and routes prompts and responses between them
* Lumerin UI-Desktop is the front end UI to interact with LLMs and the Morpheus network via the proxy-router
* Local Llama.cpp and tinyllama model to run locally 

## Tokens and Contract Information 
* Morpheus saMOR Token: `0xc1664f994fd3991f98ae944bc16b9aed673ef5fd` 
* Lumerin Morpheus Smart Contract : `0x70768f0fF919e194E11aBFC3a2eDF43213359dc1`
* Blockchain Explorer: `https://sepolia.arbiscan.io/`

## Prerequisites
* **WALLET:** For testing, you will need both `saMOR` and `saETH` tokens in your wallet. You should be able to get either of these from the usual Sepolia Arbitrum testnet faucets.
    * `saMOR` is the token used to pay for the model provider staking and consumer usage
    * `saETH` is the token used to pay for the gas fees on the network  
    * At this time, we recommend setting up a new wallet in either MetaMask or other existing crypto wallet 
        * You will need both the wallet's `private key` and `mnemonic` from your wallet to startup and interact with both the proxy-router and the UI-Desktop 
        * AS ALWAYS:  **NEVER** share either of these two items with anyone else, as they are the keys to your wallet and can be used to access your funds
        * In the future, as the UI-Desktop functionality is developed, all of this will be handled in the UI and you will not need to manually enter these items

* **ETH NODE:** You will need access to either a public or private Sepolia Arbitrum ETH node that the proxy-router and ui-desktop will use to monitor events on the blockchain
    * We have provided an existing public node example `https://arbitrum-sepolia.blockpi.network/v1/rpc/public` in the .env-example file for you to use, however, 
    * We stronly recommend either a wss: or https: private endpoint from a trusted service like Alchemy or Infura (both offer free accounts for testing/private usage)


## Getting Started (using the provided releases - run from source instructions coming soon)
### Common Steps for both Consumer and Provider: 
1. PreRequisites: 
    1. Have Mnemonic and Private key for a Consumer Wallet (Suggest setting one up in MetaMask so you can pull the mnemonic and private key) 
    2. Obtain or Transfer saMOR and saETH to this wallet for testing (suggest 1000 saMOR and at least 0.1 saETH) - Sepolia Arbitrum Chain 
2. Download latest release for your operating system: https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases
3. Extract the zip to a local folder (examples)
    * Windows: `(%USERPROFILE%)/Downloads/morpheus)` 
    * Linux & MacOS: `~/Downloads/morpheus`
    * On MacOS you may need to execute `xattr -c mor-launch proxy-router ui-desktop.app` in a command window to remove the quarantine flag on MacOS
4. Edit the .env file (this is a hidden file, please use your OS specific method to show hidden files) 
    * Change `ETH_NODE_ADDRESS=` (you can setup a free one in Alchemy or Infura) 
    * Change `WALLET_PRIVATE_KEY=` This will be the private key of the Wallet you setup previously

### Consumer (Local LLM, Proxy-Router & UI-Desktop): 
1. Assuming that you have already setup the prerequisites and downloaded the latest release, you can follow the steps below to get started
2. Launch the node - this should open a command window to see local LLM model server and proxy-router start and then should launch the user interface  
    * Windows: Double click the `mor-launch.exe` (You will need to tell Windows Defender this is ok to run) 
    * Linux & MacOS: Open a terminal and navigate to the folder and run `./mor-launch`
3. Startup User Interface: 
    1. Read & accept terms & Conditions 
    2. Set a strong password (this is for the UI-Desktop only)
    3. When prompted to create the wallet from a new Mnemonic, select **Recover your wallet** from `Saved Mnemonic` instead.  
        * This is important so that the private key for the proxy-router (in the .env file) and the wallet running in the UI are the same 
        * If you create a new wallet, the proxy-router will be listening to a different wallet address than the UI is managing and this will not work.
4. Local Test: Once the UI is up and running,
    1. You should see tokens for saETH and saMOR that you sent to this wallet earlier. If not, either the ETH_NODE_ADDRESS is incorrect or the wallet address is not the same as the one you sent the tokens to
    2. Click on the `Chat` icon on the left side of the screen
    3. Make sure the `Local Model` is selected
    4. Begin your conversation with the model by typing in the chat window and pressing `Enter`
        * You should see the model respond with the appropriate response to the prompt you entered, if not, then likely the configuration in the .env file is incorrect or the model is not running correctly
5. Remote Test: Once you've verified that your wallet can access the blockchain and you can see the local model working, you can switch to a remote model and test that as well
    1. In the `Chat` window, select `Change Model `
        1. Select a different model from remote providers
        2. DropDown and select the contract address of the model you want to use 
        3. Click Change 
        4. Click Open Session 
        5. MUST Enter at least **5** MOR to open session 
    3. You can now chat with the remote model and see the responses in the chat window 
6. Cleanup/Closeout 
    * Manually End all Remote Sessions: 
        * In the Chat Window, click on the Time icon to the right of the Model line - this will expand and show current sessions, click the "X" next to each one to make sure it's closed 
    * Closing the UI-Desktop window should leave the CMD window open
        * You’ll have to ctrl-c in the window to kill the local model and proxy-router
    * To FULLY delete and force a clean startup of the UI (including forcing new password and mnemonic recovery), delete the ui-desktop folder and start the UI again
        * Windows:  `%USERPROFILE%\AppData\Roaming\ui-desktop`
        * Linux: `~/.config/ui-desktop`
        * MacOS: `~/Library/Application Support/ui-desktop`

### Provider (Local LLM to offer, Proxy-Router running as background/service): 
This section is used for offering your hosted LLM model to the network for others to use.

**At this time, we are not onboarding any new providers, but you can follow the steps below to see how it would work and will be automated and simplified in the future.**
1. SETUP PROVIDER / MODEL / BID: 
    1. WEB3/Arbiscan/Metamask: Authorize Diamond Contract to spend on the Provider's behalf 
        1. https://sepolia.arbiscan.io/address/0xc1664f994fd3991f98ae944bc16b9aed673ef5fd#writeContract 
        2. Connect to Web3 (connect Provider wallet) 
        3. Click Approve 
        4. Spender Address = Diamond Contract 
        5. Authorized Amount = remember that this is in the form 1*10^18 so make sure there's enough MOR granted to cover the contract fees 
        6. The Diamond Contract is now authorized to spend MOR on provider's behalf 
    2. Create Provider in the contract:  
        1. Connect Wallet (approve via MM) 
        2. Select ProviderRegistry/providerRegister function 
            1. addr = Provider address 
            2. addStake = Amount of stake for provider to risk - Stake can be 0 now 
            3. Endpoint = Publicly accessible endpoint for provider (ip:port or fqdn:port no protocol) eg: `mycoolmornode.domain.com:3989`
    3. Create Model in the contract:
        1. Select ModelRegistry/modelRegister function
            1. modelId: random 32byte/hex that will uniquely identify model (uuid)
            2. ipfsCID: another random32byte/hex for future use (model library)
            3. Fee: fee for the model usage - 0 for now
            4. addStake: stake for model usage - 0 for now 
            5. Owner: Provider Wallet Address 
            6. name: Human Readable model like "Llama 2.0" or "Mistral 2.5" or "Collective Cognition 1.1" 
            7. tags: comma delimited tags for the model 
            8. Capture the `modelID` from the JSON response
    4. Offer Model Bid in the contract: 
        1. Select Marketplace/postModelBid function
            1. providerAddr: Provider Wallet Address
            2. modelID: Model ID Created in last step: 
            3. pricePerSecond: this is in 1*10^18 format so 100000000000 should make 5 minutes for the session around 37.5 saMOR) 
            4. Click WRITE and confirm via MM 
2. LAUNCH LOCAL MODEL:
    * On your server (that is accessible at the `provider endpoint` you provided in the contract), launch the local model server
    * You can use the provided `llama.cpp` and `tinyllama` model to test locally
    * If your local model is listening on a different port locally, you will need to modify the `OPENAI_BASE_URL` in the .env file to match the correct port
3. LAUNCH PROXY-ROUTER: 
    * On your server, launch the proxy-router with the modified .env file shown in the common pre-requisites section
        * Windows: Double click the `proxy-router.exe` (You will need to tell Windows Defender this is ok to run)  
        * Linux & MacOS: Open a terminal and navigate to the folder and run `./proxy-router`from the morpheus/proxy-router folder
    * This will start the proxy-router in the background and begin monitoring the blockchain for events
4. VERIFY PROVIDER SETUP 
    * On a separate machine and with a separate wallet, you can follow the consumer steps above to verify that your model is available and working correctly