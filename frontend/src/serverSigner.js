import * as fcl from '@onflow/fcl';
import dotenv from 'dotenv';

dotenv.config();

const API = 'http://localhost:8080';

const addr = process.env.ADMIN_ADDRESS;
const keyId = process.env.ADMIN_KEY_ID;

const signingFunction = async (signable) => {
  const response = await fetch(`${API}/auth`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ method: 'purchase_tokens', payload: signable }),
  });
  console.log(await response.json());
  const signature = Buffer.from(await response.json(), 'base64').toString('hex');
  console.log({
    addr: fcl.withPrefix(addr),
    keyId,
    signature,
  });
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
