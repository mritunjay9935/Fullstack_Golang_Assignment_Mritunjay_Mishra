import React, { useState, useEffect } from 'react';
import { w3cwebsocket as W3CWebSocket } from 'websocket';
import axios from 'axios';

const client = new W3CWebSocket('ws://127.0.0.1:8000/');

function App() {
    const [cacheItems, setCacheItems] = useState({});
    const [key, setKey] = useState('');
    const [value, setValue] = useState('');
    const [expiration, setExpiration] = useState('');
    const [fetchedValue, setFetchedValue] = useState('');

    const handleFetchItem = async () => {
        try {
            const response = await axios.get(`http://localhost:8000/item?key=${key}`);
            setFetchedValue(response.data.value);
        } catch (error) {
            if (error.response) {
                // The request was made and the server responded with a status code
                // that falls out of the range of 2xx
                console.log(error.response.data);
                console.log(error.response.status);
                console.log(error.response.headers);
            } else if (error.request) {
                // The request was made but no response was received
                console.log(error.request);
            } else {
                // Something happened in setting up the request that triggered an Error
                console.log('Error', error.message);
            }
            console.log(error.config);
        }
    };

    useEffect(() => {
        client.onopen = () => console.log('WebSocket Client Connected');
        client.onmessage = (message) => setCacheItems(JSON.parse(message.data));
        return () => client.close();
    }, []);

    useEffect(() => {
        const interval = setInterval(() => {
            const now = Date.now();
            setCacheItems((prevItems) => {
                return Object.fromEntries(
                    Object.entries(prevItems).filter(([_, item]) => item.expiration >= now)
                );
            });
        }, 1000);
        return () => clearInterval(interval);
    }, []);

    const handleSetItem = async (e) => {
        e.preventDefault();
        try {
            await axios.post('http://localhost:8000/item', {
                key,
                value,
                expiration: parseInt(expiration, 10)
            });
            // rest of your code
        } catch (error) {
            if (error.response) {
                // The request was made and the server responded with a status code
                // that falls out of the range of 2xx
                console.log(error.response.data);
                console.log(error.response.status);
                console.log(error.response.headers);
            } else if (error.request) {
                // The request was made but no response was received
                console.log(error.request);
            } else {
                // Something happened in setting up the request that triggered an Error
                console.log('Error', error.message);
            }
            console.log(error.config);
        }
    };

    const handleDeleteItem = async (keyToDelete) => {
        await axios.delete(`http://localhost:8000/item?key=${keyToDelete}`);
        setCacheItems((prevItems) => {
            const updatedItems = { ...prevItems };
            delete updatedItems[keyToDelete];
            return updatedItems;
        });
    };

    return (
        <div>
            <h1>LRU Cache Items</h1>
            <ul>
                {Object.entries(cacheItems).map(([key, item]) => (
                    <li key={key} className="cache-item">
                        {key}: {item.value}
                        <button onClick={() => handleDeleteItem(key)}>Delete</button>
                    </li>
                ))}
            </ul>
            <form onSubmit={handleSetItem} className="form-add-item">
                <input type="text" value={key} onChange={(e) => setKey(e.target.value)} placeholder="Key" required />
                <input type="text" value={value} onChange={(e) => setValue(e.target.value)} placeholder="Value" required />
                <input type="number" value={expiration} onChange={(e) => setExpiration(e.target.value)} placeholder="Expiration (seconds)" required />
                <button type="submit">Set Item</button>
            </form>
            <h1>Cache Item Fetcher</h1>
            <input
                type="text"
                value={key}
                onChange={(e) => setKey(e.target.value)}
                placeholder="Enter key to fetch"
            />
            <button onClick={handleFetchItem}>Fetch Item</button>
            {fetchedValue && (
                <p>
                    Value: <strong>{fetchedValue}</strong>
                </p>
            )}
        </div>
    );
}

export default App;