import React from 'react';

import { ApolloProvider } from '@apollo/react-hooks';

import ApolloClient from 'apollo-boost';
import { gql } from 'apollo-boost';
import fetch from 'node-fetch';

const client = new ApolloClient({
    uri: 'https://48p1r2roz4.sse.codesandbox.io',
    fetch: fetch
});

client
  .query({
    query: gql`
      {
        rates(currency: "USD") {
          currency
        }
      }
    `
  })
  .then(result => console.log(result));

const App = () => (
  <ApolloProvider client={client}>
    <div>
      <h2>My first Apollo app 🚀</h2>
    </div>
  </ApolloProvider>
);

export default App;
