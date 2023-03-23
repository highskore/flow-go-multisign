import * as fcl from '@onflow/fcl';
import dotenv from 'dotenv';

dotenv.config();

const API = 'http://localhost:8080';

const addr = process.env.REACT_APP_ADMIN_ADDRESS;
const keyId = process.env.REACT_APP_ADMIN_KEY_INDEX;

const signingFunction = async (signable) => {
  const response = await fetch(`${API}/cosign`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ method: 'purchase_tokens', payload: signable }),
  });
  const signature = Buffer.from(await response.json(), 'base64').toString('hex');
  return {
    addr: fcl.withPrefix(addr),
    keyId,
    signature,
  };
};

export const serverAuthorization = async (account) => {
  return {
    ...account,
    tempId: `${addr}-${keyId}`,
    addr: fcl.sansPrefix(addr),
    keyId: Number(keyId),
    signingFunction,
  };
};
