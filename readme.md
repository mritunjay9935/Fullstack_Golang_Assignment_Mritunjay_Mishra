# Apica Assignment

This project is a simple implementation of a Least Recently Used (LRU) cache system. It consists of a backend server written in Go and a frontend application built with React.

## Backend

The backend server is responsible for managing the cache items. It provides a WebSocket connection for real-time updates and a REST API for CRUD operations on cache items.

### Endpoints

- `POST /item`: Add a new item to the cache.
- `GET /item`: Fetch an item from the cache.
- `DELETE /item`: Remove an item from the cache.

### WebSocket

The server provides a WebSocket connection at `ws://localhost:8000/`. Clients can connect to this WebSocket to receive real-time updates whenever the cache changes.

## Frontend

The frontend application is a simple React app that provides a user interface for interacting with the cache. It connects to the backend server via the WebSocket connection to receive real-time updates and uses Axios to make HTTP requests to the REST API.

### Features

- Add a new item to the cache.
- Fetch an item from the cache.
- Delete an item from the cache.
- Display a list of all items in the cache.

## Setup

### Backend

Navigate to the backend directory and run the following command to start the server:

```bash
go run .
```

### Frontend

Navigate to the frontend directory and install the dependencies:

```bash
npm install
```

Then, start the React app:

```bash
npm start
```

The app will be available at `http://localhost:3000`.

## Technologies Used

- Go
- Gorilla WebSocket
- React
- Axios
- WebSocket

## Note

Ensure that both the backend server and the frontend application are running simultaneously for the system to function correctly.