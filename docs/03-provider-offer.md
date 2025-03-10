
# Creating Provider, Model and Bid on the Blockchain:
**Contract Minimums** As of 11/22/2024: 
* "providerMinStake": `200000000000000000`, (0.2 MOR)
* "modelMinStake": `100000000000000000`, (0.1 MOR)
* "marketplaceBidFee": `300000000000000000`, (0.3 MOR)
* "bidPricePerSeconMin `10000000000`, (0.00000001 MOR)

**Needed information (samples):**
* Provider/Owner:   `0x9E26Fea97F7d644BAf62d0e20e4d4b8F836C166c` # Your ERC-20 Wallet with MOR & ETH
* Endpoint:         `server.domain.com:3333` # Internet **publicly accessible** server/node access point 
* Model ID:         `0xe1e6e3e77148d140065ef2cd4fba7f4ae59c90e1639184b6df5c84` # Random 32byte/hex that you generate 
* ipfcCID:          `0xc2d3a5e4f9b7c1a2c8f0b1d5e6c78492fa7bcd34e9a3b9c9e18f25d3be47a1f6` # 32byte/hex from CID of the model file in the ipfs network
* Model Name:       `CapybaraHermes-v2.5-Mistral-7b` # Human Readable name for the model
* Bid Cost:         `10000000000` (0.00000001 MOR) # What will the model cost per second to use

## Steps
1. To complete these steps, you will need to be running the proxy-router and also have access to the API Port (default=8082)for the Swagger API Interface
    1. http://localhost:8082/swagger/index.html

1. Authorize Diamond Contract to spend on the Provider's behalf 
    1. http://localhost:8082/swagger/index.html#/transactions/post_blockchain_approve 
    1. Spender Address = Diamond Contract 
    1. Authorized Amount = remember that this is in the form `1*10^18` so make sure there's enough MOR granted to cover the contract fees 
        1. To become a provider, offer a model and offer a bid based on the current minimums, you will need to authorize the contract to spend at least `600000000000000000` (0.6 MOR) on your behalf
    1. The Diamond Contract is now authorized to spend MOR on provider's behalf 

1. Create Provider in the Diamond contract via swagger api:
    1. http://localhost:8082/swagger/index.html#/providers/post_blockchain_providers  
        1. addStake = Amount of stake for provider to risk 
            - Minimum Provider stake is `200000000000000000`, (0.2 MOR) 
            - Provider stake to become a subnet is `10000000000000000000000`, (10,000 MOR)
        1. Endpoint = Your **publicly accessible endpoint** for the proxy-router provider (ip:port or fqdn:port no protocol) eg: `mycoolmornode.domain.com:3333`
    
1. Add Model to IPFS:
    1. Go to http://localhost:8082/swagger/index.html#/ipfs/post_ipfs_add
        1. Obtain `Hash` from the JSON response. It's the `ipfsCID` for the model registry
        1. Pin the model to the IPFS network
            - Go to http://localhost:8082/swagger/index.html#/ipfs/post_ipfs_pin
            - Use the `Hash` from the previous step

1. Create Model in the contract:
    1. Go to http://localhost:8082/swagger/index.html#/models/post_blockchain_models and enter
        1. modelId: random 32byte/hex that will be used in conjunction with providerId to uniquely identify model (uuid)
        1. ipfsCID: 32byte/hex from the previous step
        1. Fee: fee for the model usage
        1. addStake: "modelMinStake": `100000000000000000`, (0.1 MOR) 
        1. Owner: Provider Wallet Address 
        1. name: Human Readable model like "Llama 2.0" or "Mistral 2.5" or "Collective Cognition 1.1" 
        1. tags: array of tag strings for the model 
        1. Capture the `modelID` from the JSON response 
            **NOTE** The returned `modelID` is a combination of your requested modelID and your providerID and will be required to update your models-config.json file AND when offering bids

1. Update the models-config.json file with the new modelID and restart the proxy-router
    1. Navigate to the proxy-router directory and open the `models-config.json` file
    1. Add the new modelID to the JSON array of models
    1. Save the file and restart the proxy-router

1. Offer Model Bid in the contract: 
    1. Navigate to http://localhost:8082/swagger/index.html#/bids/post_blockchain_bids and enter
        1. modelID: Model ID Created above
        1. pricePerSecond: "bidPricePerSeconMin `10000000000`, (0.00000001 MOR)
        1. Click Execute and capture the `bidID` from the JSON response

----------------