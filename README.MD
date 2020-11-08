# ATM-API

## About

## Available Commands

The project has a Makefile with the following commands.
- run: Run API at port 8080
- test
- coverage: Coverage output at `./cover.html`


## Endpoints

### POST - /withdrawals

#### Request

```
{
  "value": 170
}
```

#### Response 

```
{
  "banknotes": [
    {
      "value": 50,
      "quantity": 3
    },
    {
      "value": 20,
      "quantity": 1
    }
  ]
}
```