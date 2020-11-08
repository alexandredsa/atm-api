# ATM-API

## About

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