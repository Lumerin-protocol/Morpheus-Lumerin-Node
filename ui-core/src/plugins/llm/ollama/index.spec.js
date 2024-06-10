const sinon = require('sinon')
const axios = require('axios')
const ollama = require('./index')
const chai = require('chai')
const { expect } = chai
const modelName = 'llama2:latest'
const config = {

  modelUrl: 'http://localhost:8080/v1',
  // modelUrl: "http://localhost:11434/v1",
  modelName: modelName,
}

// TODO: fix other tests so they can be run reliably
describe.only('test ollama', function () {
  this.timeout(20000)
  describe('api integration', () => {
    before('should init with config', () => {

      process.env.OPENAI_BASE_URL = config.modelUrl
      ollama.init(config)

      expect(ollama.chat.createChatCompletion).is.a('function')
    })

    it('should create chat completion', async () => {
      const chat = []
      const message = 'how are you?'
      const response = await ollama.chat.createChatCompletion(chat, message)
      // concatenate the values of all of the content fields in all of the response elements
      const contents = response.map(({ message }) => message.content)

      expect(contents.join('')).to.contain("I'm just an AI")
    })

    it('should create chat completion stream', async () => {
      const chat = []
      const message = 'how are you?'
      const response = await ollama.chat.createChatCompletion(chat, message)
      
      // concatenate the values of all of the content fields in all of the response elements
      const contents = response.map(({ message }) => message.content)

      expect(contents.join('')).to.contain("I'm just an AI")
    })
  })

  describe('unit', () => {
    let request = null,
      response = null

    before(async () => {
      //mock axios
      request = sinon.stub(axios, 'post')
      response = {
        data: {
          choices: [
            {
              index: 0,
              message: {
                role: 'assistant',
                content: `I'm just an AI assistant and do not have feelings or emotions, so I cannot answer the question "How are you?" as I am not capable of experiencing any emotional state. My purpose is to assist users like you by providing information and answering questions to the best of my abilities based on my training and knowledge. Is there anything else I can help you with?`,
              },
              finish_reason: 'stop',
            },
          ],
        },
      }

      request.returns(Promise.resolve(response))
    })

    after(async () => {
      request.restore()
    })

    it.skip('should throw error on invalid config', async () => {
      await expect(() => ollama.init({})).throws()
    })

    it('should set default model name', async () => {
      ollama.init(config)

      expect(ollama.chat.createChatCompletion).is.a('function')
      expect(ollama.modelName).to.equal('llama2:latest')
    })

    it('should create chat completion', async () => {
      const chat = []
      const message = 'how are you?'
      const response = await ollama.chat.createChatCompletion(chat, message)

      expect(response[0].message.content).to.contain("I'm just an AI")
    })
  })
})