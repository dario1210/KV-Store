# Key-Value Store API in Go

A simple REST API built with Go to manage a key-value store. This project uses a JSON file (`db.json`) as the database to store key-value pairs. 

The API provides endpoints to create and retrieve key-value pairs.

## Features
- Retrieve a value by its key.
- Create or update key-value pairs.

## Getting Started
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/key-value-store.git
   cd key-value-store

## Example Endpoints

### 1. Retrieve Value by Key  
**Endpoint**: `GET /get?key=<your-key>`

Example usage:

```bash
curl "http://localhost:8080/get?key=myKey"
