# DataBet SPA Example

This repository provides boilerplate code as an introduction to help you build your own betting application using the 
[DataBet SPA](https://docs.data.bet/betting-spa/).

## Documentation

- [Betting Integration](https://docs.data.bet/betting-integration/)
- [Authorization](https://docs.data.bet/betting-integration-auth/)

## Getting Started

### Prerequisites

To run this project, you need to have 
[Node.js 18+ and NPM](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) installed on your system.

### Setting up the Project

#### 1. Clone the Repository and Install Dependencies

```bash
git clone https://github.com/databet-cloud/spa-example
cd spa-example
npm install
```

#### 2. Retrieve Authorization Certificates from DataBet and Generate an Authorization Token for Your Project

```bash
curl --location https://betting.int.databet.cloud/token/create \
--cert    tls.crt \
--key     tls.key \
--data '{ "locale": "en", "currency": "EUR"}'
```

#### 3. Modify the `index.html` File

Open `index.html` in your editor and replace the placeholders:
1. Replace `<PUT_YOUR_GRAPHQL_ADDRESS_HERE>` with the GQL address provided by DataBet.
2. Replace `<PUT_TOKEN_HERE>` with the token retrieved from step #2.

```html
<!-- ... -->
<script>
    bettingLoader.load(
        {
            token: 'YOUR_TOKEN_HERE',
            url: {
                gqlEndpoint: 'YOUR_GQL_ADDRESS_HERE',
                bettingPath: '/',
                pagePath: '',
                staticEndpoint: '//demo-static.data.bet',
            },
<!-- ... -->
```

#### 4. Build and Start the Server

```bash
npm run dev
```

Your browser will open automatically, and you'll see a configured SPA with sports events configured for your account.

![SPA example](img/page-example.png)

Congratulations! You've successfully set up the SPA example!

## What's Next?

Once set up, you should integrate the [Callbacks API](https://docs.data.bet/betting-bet/) within your server. Explore 
any additional APIs you need to realize the **Betting Application of Your Dreams**!
