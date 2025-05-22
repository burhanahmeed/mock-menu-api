# Waiter AI Agent n8n Workflow

This project contains an n8n workflow for a Waiter AI Agent that helps customers choose menu items based on a restaurant database.

## Features

- AI-powered menu recommendations
- Access to restaurant database
- Menu item customization options
- Time-based menu availability

## Tools Available

- Item Getter: Retrieves available menu items
- Customization Getter: Retrieves customization options for menu items

## Menu API Service

A simple Golang REST API for restaurant menu management.

## Description

This API provides endpoints to access restaurant menu data including categories, items, and customization options. The data is stored in a JSON file and served through a RESTful API built with Go and the Gorilla Mux router.

## Installation

1. Clone this repository
2. Install Go (1.19 or later)
3. Run `go mod tidy` to install dependencies
4. Run `go run main.go` to start the server

## API Endpoints

| Endpoint | Method | Parameters | Description |
|----------|--------|------------|-------------|
| **Menu Endpoints** |
| `/categories` | GET | None | List all menu categories |
| `/category` | GET | `id` | Get a specific category by ID |
| `/items` | GET | `category_id` | List all items in a specific category |
| `/all-items` | GET | None | List all menu items across all categories |
| `/item` | GET | `id` | Get a specific menu item by ID |
| `/customizations` | GET | `item_id` | List all customizations for a specific item |
| `/customization` | GET | `id` | Get a specific customization by ID |
| **Cart Endpoints** |
| `/cart` | GET | None | Get the current cart contents |
| `/cart/add` | POST | JSON body | Add an item to the cart |
| `/cart/remove` | POST | JSON body | Remove an item from the cart |
| `/cart/clear` | POST | None | Clear all items from the cart |

## Data Structure

The menu data is structured as follows:
- Categories (e.g., Appetizers, Main Courses)
  - Items (e.g., Calamari, Steak)
    - Customizations (e.g., Spice Level, Cooking Temperature)
      - Options (e.g., Mild, Medium, Well Done)

## Usage

1. Start the server: `go run main.go`
2. Access the API at `http://localhost:8080/`

## Setup Instructions

1. Install n8n
2. Import the workflow.json file
3. Configure the necessary credentials
4. Start the workflow
