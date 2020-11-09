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

#### Responses 

- `201` - Created
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

- `422` - Unprocessable Entity

In case of available notes that don't attend the value

Scenario: 
- Available Notes: `[50, 10]`
- Response

```
{
  "Reason": "Withdrawal not supported for value, left over: 5"
}
``` 