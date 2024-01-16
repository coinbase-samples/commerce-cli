# Coinbase Commerce CLI tool

The _Coinbase Commerce CLI Tool_ is a cli interface for interacting with the [_Commerce Go SDK_](https://github.com/coinbase-samples/commerce-sdk-go) sample library, an SDK built on top of the [Coinbase Commerce API](https://www.coinbase.com/commerce).

Easily create, retrieve, and review charges using command line arguments.

## Getting started

1. Set your API key as an environment variable named "COMMERCE_API_KEY" by running

```shell
export COMMERCE_API_KEY=123-YOUR-API-KEY
```

2. Build the CLI executable by running

```shell
go build -o commerce
```

3. Move executible in local bin

```shell
mv ./commerce usr/loca/bin
```

4. Create a charge by running

   ```shell
   commerce charges --setPrice 1.5
   ```

   This will create a $1.50 charge with a payment link (hosted_url)
   example output should be:

```json
{
  "data": {
    "brand_color": "#000000",
    "brand_logo_url": "https://res.cloudinary.com/commerce/image/upload/v1653516296/dlwoolpero6qgsffxmpz.jpg",
    "charge_kind": "WEB3",
    "code": "6EE5J2V8",
    "confirmed_at": "",
    "created_at": "2024-01-15T21:32:49Z",
    "expires_at": "2024-01-17T21:32:49Z",
    "hosted_url": "https://commerce.coinbase.com/pay/131bf7a0-69be-4b63-a8da-847ff831eb46",
    "id": "131bf7a0-69be-4b63-a8da-847ff831eb46",
    "organization_name": "Patrick's Fine Art",
    "pricing": {
      "local": {
        "amount": "1.5",
        "currency": "USD"
      },
      "settlement": {
        "amount": "1.5",
        "currency": "USDC"
      }
    },
    "pricing_type": "fixed_price",
    "redirects": {
      "cancel_url": "",
      "success_url": "",
      "will_redirect_after_success": false
    },
    "support_email": "patrick.hughes@coinbase.com",
    "timeline": [
      {
        "status": "NEW",
        "time": "2024-01-15T21:32:49Z"
      }
    ],
    "web3_data": {
      "transfer_intent": {
        "call_data": {
          "deadline": "",
          "fee_amount": "",
          "id": "",
          "operator": "",
          "prefix": "",
          "recipient": "",
          "recipient_amount": "",
          "recipient_currency": "",
          "refund_destination": "",
          "signature": ""
        },
        "metadata": {
          "chain_id": 0,
          "contract_address": "",
          "sender": ""
        }
      },
      "success_events": [],
      "failure_events": [],
      "contract_addresses": {
        "1": "0x7915f087685fffD71608E5d118f3B70c27D9eF4e",
        "137": "0x7f52269092F2a5EF06C36C91e46F9196deb90336",
        "8453": "0x9Bb4D44e6963260A1850926E8f6bEB8d5803836F"
      }
    }
  }
}
```

## Usage

The Coinbase Commerce CLI tool allows you to create and view charges with a valid API key as well as webhook events.

### Create a charge

```shell
commerce charges --setPrice [amount]
```

Replace `[amount]` with the desired charge amount (e.g., 1.5 for $1.50).

### Retrieve a charge

```shell
commerce charges --get [charge_id]
```

Replace [charge_id] with the specific ID of the charge you want to retrieve.

### Events

The Coinbase Commerce CLI tool also supports the retrieval of events.

To retrieve all events for you account run:

```shell
commerce events --all
```

### Retrieve a specific event

> [!TIP]
> Be sure to use the event id and **not** the charge id or charge code this will prevent any errors from occuring.

```shell
commerce events --get [event_id]
```

Replace [event_id] with the ID of the event you wish to retrieve.
