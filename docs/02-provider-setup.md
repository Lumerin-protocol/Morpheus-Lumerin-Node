
# Provider Hosting (Local LLM to offer, Proxy-Router running as background/service): 

## Assumptions: 
* Your AI model has been configured, started and made available to the proxy-router server via a private endpoint (IP:PORT or DNS:PORT) eg: `http://mycoolaimodel.domain.com:8080`
    * Optional
        * You can use the provided `llama.cpp` and `tinyllama` model to test locally
        * If your local model is listening on a different port locally, you will need to modify the `models-config.json` file to match the correct port
* You have an existing funded wallet with MOR and ETH and also have the `private key` for the wallet (this will be needed for the .env file configuration)
* You have created an Alchemy or Infura free account and have a private API key for the Arbitrum Sepolia testnet (wss://arb-sepolia.g.alchemy.com/v2/<your_private_alchemy_api_key>)
* Your proxy-router must have a **publicly accessible endpoint** for the provider (ip:port or fqdn:port no protocol) eg: `mycoolmornode.domain.com:3333` - this will be used when creating the provider on the blockchain

## Installation & Configuration Steps:
1. Obtain the software: 
    1. Package: Download latest release for your operating system: https://github.com/Lumerin-protocol/Morpheus-Lumerin-Node/releases
        * Mainnet releases will be prefixed with `main-*`
        * Testnet releases will be prefixed with `test-*` 
    2. Source Code: 
        * Clone the repository: `git clone -b branch https://github.com/Lumerin-protocol/Morpheus-Lumerin-Node.git` 
        * Mainnet Branch = `main` 
        * Testnet Branch = `test`
        * Development Branch = `dev`(not recommended unless directed by the development team)

1. Extract the zip to a local folder (examples)
    * Windows: `(%USERPROFILE%)/Downloads/morpheus)` 
    * Linux & MacOS: `~/Downloads/morpheus`
    * On MacOS you may need to execute `xattr -c proxy-router` in a command window to remove the quarantine flag on MacOS

1. Environment configuration 
    * In most cases, the default .env file will work for the proxy-router...In some cases you will want to modify the .env file with advanced capability (log entries, private keys, private endpoints, etc)
    * Please see [proxy-router.all.env](proxy-router.all.env) for more information on the key values available in the .env file
    1. Choose OS environment you are working with: 
        * Linux/Mac: `env.example`  or `env.example.win` for Windows
        * Change the name of the desired file to `.env` 
        * Edit values within the file (Wallet_Private_Key, for example) as desired
    2. Choose the **blockchain** you'd like to work on...**Arbitrum MAINNET is the default** 
        * To operate on the Sepolia Arbitrum TESTNET,  
        * Edit the .env file and 
        * Uncomment the `TESTNET VALUES` and comment the `MAINNET VALUES` lines & save the file

1. **(OPTIONAL)** - External Provider or Pass through 
    * In some cases you will want to leverage external or existing AI Providers in the network via their own, private API
    * Dependencies: 
        * `model-config.json` file in the proxy-router directory
        * proxy-router .env file for proxy-router must also be edited to adjust `MODELS_CONFIG_PATH=<path_to_proxy-router>/models-config.json`
    * Once your provider is up and running, deploy a new model and model bid via the API interface (you will need the `model_ID` for the configuration)
    * Edit the model-config.json to the following json format
        * The JSON ID will be the ModelID that you created above, modelName, apiTYpe, apiURL and apiKey are from the external provider and specific to their offered models 
        * Full explanation of models-config.json can be found here [proxy-router models-config.json](proxy-router.models-config.json.md)
    * Once the model-config.json file is updated, the morpheus node will need to be restarted to pick up the new configuration (not all models (eg: image generation can be utilized via the MorpheusUI, but API integration is possible)

## Start the Proxy Router 
1. On your server, launch the proxy-router with the modified .env file shown above
    * Windows: Double click the `proxy-router.exe` (You will need to tell Windows Defender this is ok to run)  
    * Linux & MacOS: Open a terminal and navigate to the folder and run `./proxy-router`from the morpheus/proxy-router folder
1.  This will start the proxy-router and begin monitoring the blockchain for events
    
## Validating Steps:
1. Once the proxy-router is running, you can navigate to the Swagger API Interface (http://localhost:8082/swagger/index.html as example) to validate that the proxy-router is running and listening for blockchain events
1. You can also check the logs in the `./data` directory for any errors or issues that may have occurred during startup
1. Once validated, you can move on and create your provider, model and bid on the blockchain [03-provider-offer.md](03-provider-offer.md)
