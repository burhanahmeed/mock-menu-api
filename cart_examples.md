# Cart API Examples

## Add to Cart

### Request
```json
POST /cart/add
Content-Type: application/json

{
  "item_id": "app1",
  "quantity": 2,
  "customizations": [
    {
      "id": "spice_level",
      "name": "Spice Level",
      "selected_options": [
        {
          "id": "medium",
          "name": "Medium",
          "price": 0
        }
      ]
    },
    {
      "id": "dipping_sauce",
      "name": "Extra Dipping Sauce",
      "selected_options": [
        {
          "id": "aioli",
          "name": "Garlic Aioli",
          "price": 1.50
        }
      ]
    }
  ]
}
```

### Response
```json
{
  "items": [
    {
      "id": "app1_0",
      "item_id": "app1",
      "name": "Crispy Calamari",
      "price": 12.99,
      "quantity": 2,
      "customizations": [
        {
          "id": "spice_level",
          "name": "Spice Level",
          "selected_options": [
            {
              "id": "medium",
              "name": "Medium",
              "price": 0
            }
          ]
        },
        {
          "id": "dipping_sauce",
          "name": "Extra Dipping Sauce",
          "selected_options": [
            {
              "id": "aioli",
              "name": "Garlic Aioli",
              "price": 1.50
            }
          ]
        }
      ],
      "total_price": 28.98
    }
  ],
  "total_items": 2,
  "total_price": 28.98
}
```

## Get Cart

### Request
```
GET /cart
```

### Response
```json
{
  "items": [
    {
      "id": "app1_0",
      "item_id": "app1",
      "name": "Crispy Calamari",
      "price": 12.99,
      "quantity": 2,
      "customizations": [
        {
          "id": "spice_level",
          "name": "Spice Level",
          "selected_options": [
            {
              "id": "medium",
              "name": "Medium",
              "price": 0
            }
          ]
        },
        {
          "id": "dipping_sauce",
          "name": "Extra Dipping Sauce",
          "selected_options": [
            {
              "id": "aioli",
              "name": "Garlic Aioli",
              "price": 1.50
            }
          ]
        }
      ],
      "total_price": 28.98
    }
  ],
  "total_items": 2,
  "total_price": 28.98
}
```

## Remove from Cart

### Request
```json
POST /cart/remove
Content-Type: application/json

{
  "cart_item_id": "app1_0"
}
```

### Response
```json
{
  "items": [],
  "total_items": 0,
  "total_price": 0
}
```

## Clear Cart

### Request
```
POST /cart/clear
```

### Response
```json
{
  "items": [],
  "total_items": 0,
  "total_price": 0
}
```
