# Example models config file.  Local model configurations are stored in this file
* `root_key` (required) is the model id
* `modelName` (required) is the name of the model
* `apiType` (required) is the type of the model api.  Currently supported values are "prodia" and "openai"
* `apiUrl` (required) is the url of the LLM server or model API
* `apiKey` (optional) is the api key for the model
* `cononcurrentSlots` (optional) are number of available distinct chats on the llm server and used for capacity policy
* `capacityPolicy` (optional) can be one of the following: "idle_timeout", "simple" 

## Examples of models-config.json entries 
* The first key (`0x000...0000`) is the model id of the local default model (llama2)
* The middle two keys are examples of externally hosted and owned models where the morpheus-proxy-router enables proxying requests to the external model API
* The last key is an example of a model hosted on a server owned by the morpheus-proxy-router operator

```bash
{
    "0x0000000000000000000000000000000000000000000000000000000000000000": {
        "modelName": "llama2",
        "apiType": "openai",
        "apiUrl": "http://localhost:8080/v1"
    }, 
    "0x60d5900b10534de1a668fd990bd790fa3fe04df8177e740f249c750023a680fb": {
        "modelName": "v1-5-pruned-emaonly.safetensors [d7049739]",
        "apiType": "prodia-sd",
        "apiUrl": "https://api.prodia.com/v1",
        "apiKey": "replace-with-your-api-key"
    },
    "0x0d90cf8ca0a811a5cd2347148171e9a2401e9fbc32f683b648c5c1381df91ff7": {
        "modelName": "animagineXLV3_v30.safetensors [75f2f05b]",
        "apiType": "prodia-sdxl",
        "apiUrl": "https://api.prodia.com/v1",
        "apiKey": "replace-with-your-api-key"
    },
    "0x06c7b502d0b7f14a96bb4fda5f2ba941068601d5f6b90804c9330e96b093b0ce": {
        "modelName": "LMR-Collective Cognition Mistral 7B",
        "apiType": "openai",
        "apiUrl": "http://llmserver.domain.io:8080/v1",
        "concurrentSlots": 8,
        "capacityPolicy": "simple"
    }
}
```