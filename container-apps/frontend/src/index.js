import React from 'react';
import ReactDOM from 'react-dom';
import {
  ApolloClient,
  InMemoryCache,
  ApolloProvider,
} from "@apollo/client";
import { 
  Route,
  Routes,
  BrowserRouter as Router 
} from "react-router-dom";

import { Home } from "./containers/Home";
import { Users } from "./containers/Users";
import App from './App';
import 'bootstrap/dist/css/bootstrap.min.css';
import './index.css';

const client = new ApolloClient({
  uri: 'https://api-gateway.calmdune-59e37c6b.westeurope.azurecontainerapps.io:3005/query',
  cache: new InMemoryCache()
});

ReactDOM.render(
  <React.StrictMode> 
    <ApolloProvider client={client}>
      <Router>
        <Routes>
          <Route path="/" element={<App />}>
            <Route index element={<Home />} />
            <Route path="users" element={<Users />}>
            </Route>
          </Route>
        </Routes>
      </Router>
    </ApolloProvider>
  </React.StrictMode>,
  document.getElementById('root')
);
