import './App.css';
import { useState, useEffect } from 'react';

import { serverAuthorization } from './serverSigner';

const fcl = require('@onflow/fcl');
const t = require('@onflow/types');
fcl
  .config()
  .put('accessNode.api', 'https://testnet.onflow.org')
  .put('discovery.wallet', 'https://flow-wallet-testnet.blocto.app/authn');

const App = () => {
  const [user, setUser] = useState();

  const [transaction, setTransaction] = useState();

  useEffect(() => {
    fcl.currentUser().subscribe(setUser);
  }, []);
  const multiSign = async () => {
    setTransaction('');
    const transactionId = await fcl
      .send([
        fcl.transaction`
      transaction() {
        prepare(frontendUser: AuthAccount, backendAdmin: AuthAccount) {

        }
      }
      `,
        fcl.payer(serverAuthorization),
        fcl.proposer(fcl.authz),
        fcl.authorizations([fcl.authz, serverAuthorization]),
        fcl.limit(9999),
      ])
      .then(fcl.decode);

    setTransaction(transactionId);
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>{user && user.addr ? user.addr : 'Not logged in.'}</h1>
        <div>
          <button onClick={() => fcl.authenticate()}>Log In</button>
          <button onClick={() => fcl.unauthenticate()}>Log Out</button>
        </div>
        <button onClick={() => multiSign()}>Run Multi-Sign Tx</button>
        {transaction && (
          <p>
            <a href={`https://flow-view-source.com/testnet/tx/${transaction}`} target="_blank" rel="noreferrer">
              {transaction}
            </a>
          </p>
        )}
      </header>
    </div>
  );
};

export default App;
