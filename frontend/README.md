# Flow Multisign Demo - Frontend

This is a simple frontend that allows you to execute multisignature transactions on the Flow blockchain.
Forked from @jacob-tucker (https://github.com/jacob-tucker/multi-sign/tree/master/frontend)

## Prerequisites

Before running the frontend app, make sure you have the following installed:

- Node.js (at least version 14.17.0)

## Installation

1. Install frontend dependencies:

   ```bash
    npm install
   ```

2. Create a `.env.local` file in the root directory of the project and add the following environment variables:

   ```bash
    # FLOW
    REACT_ADMIN_ADDRESS=0x01cf0e2f2f715450 - Your flow account admin address
    REACT_ADMIN_KEY_INDEX=0 - Your flow account admin key index
   ```

## Usage

1. Start the frontend service using the following command:
   ```bash
    npm run start
   ```
