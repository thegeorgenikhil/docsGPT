# DocsGPT

DocsGPT is a powerful web app that allows you to embed your documents then query them using natural language.

## How It Works

Using OpenAI's Embedding API, it embeds your documents into set of vectors and then stores the content in a NoSQL and the vectors in a database(Pinecone). When you query the app, it uses the same API to embed your query and then uses cosine similarity to find the most similar documents.

From the similar documents, it selects the top 3 results and creates a request to OpenAI's Completions API with the top result along with your query and ask it to return a meaningful response by taking the top results as reference.

## Setup

For this to work, you need to have Docker installed on your machine.

It uses Pinecone to store all the vectors.So you need to have a Pinecone account and a Pinecone API key. You can get one [here](https://www.pinecone.io/).

For the OpenAI API, you need to have an API key. You can get one [here](https://platform.openai.com/account/api-keys/).

1. Clone the repo

```bash
git clone https://github.com/thegeorgenikhil/docsGPT.git
```

2. Rename the `.env.example` file to `.env` and fill in the values.

```bash
PORT=9000
OPENAI_API_KEY= # This is the API key you get from OpenAI
PINECONE_API_KEY= # This is the API key you get from Pinecone
PINECONE_INDEX_URL= # This is the URL of the index you want to use.
MONGO_URI=mongodb://mongodb:27017
```

3. Run the app

```bash
docker-compose up
```

Your app should be running on `localhost:80` with backend on `localhost:9000`.

## Screenshots

![image](https://user-images.githubusercontent.com/56214901/228553999-4bb07e12-bf72-48c9-9a5f-1e01b742d7d9.png)

## TODO

- [ ] Replace all the alerts in the UI with toasts
- [ ] Handle the error when the prompt exceeds the max token limit
