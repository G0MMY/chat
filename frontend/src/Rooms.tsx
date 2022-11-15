import { Button, TextField } from "@mui/material";
import React, { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";

interface Room {
    id: number
    name: number
}

export default function Rooms() {
    const [token, setToken] = useState('');
    const navigate = useNavigate();
    const { state } = useLocation();
    const { username } = state;
    const [error, setError] = useState(false);
    const [errorMessage, setErrorMessage] = useState('');
    const [roomName, setRoomName] = useState('');
    const [rooms, setRooms] = useState<Room[]>([]);

    const getRooms = (currentToken?: string) => {
        const requestOptions = {
            method: "GET",
            headers: {
                "Content-type": "application/json",
                "Token": currentToken? currentToken : token
            }
        };
        fetch("/rooms?username=" + username, requestOptions).then((resp) => {
            if (resp.ok) {
                return resp.json();
            }
            return resp.text().then(text => { throw new Error(text) })
        }).then((data) => {
            console.log("t")
           console.log(data)
        }).catch((err: Error) => {
            setError(true);
            setErrorMessage(err.message);
        });
    }

    const createRoom = () => {
        const requestOptions = {
            method: "POST",
            headers: {
                "Content-type": "application/json",
                "Token": token
            },
            body: JSON.stringify({
                name: roomName
            })
        };
        fetch('/rooms', requestOptions).then((resp) => {
            if (resp.ok) {
                return resp.json();
            }
            return resp.text().then(text => { throw new Error(text) })
        }).then((data) => {
            joinRoom(data.id)
        }).catch((err: Error) => {
            setError(true);
            setErrorMessage(err.message);
        });
    }

    const joinRoom = (id: number) => {
        const requestOptions = {
            method: "POST",
            headers: {
                "Content-type": "application/json",
                "Token": token
            },
            body: JSON.stringify({
                username: username,
                roomId: id
            })
        };
        fetch('/rooms/join', requestOptions).then((resp) => {
            if (resp.ok) {
                return resp.json();
            }
            return resp.text().then(text => { throw new Error(text) })
        }).then((data) => {
            // add to rooms
        }).catch((err: Error) => {
            setError(true);
            setErrorMessage(err.message);
        });
    }

    useEffect(() => {
        const tempToken = sessionStorage.getItem('Token');
        if (tempToken !== null) {
            setToken(tempToken);
            fetch("/validate?Token=" + tempToken).then((resp) => {
                if (resp.status !== 200) {
                    navigate("/");
                } else {
                    getRooms(tempToken);
                }
            })
        } else {
            navigate("/");
        }
    }, [])

    return (
        <div>
            <h1>Welcome {username}</h1>
            {error? errorMessage: ''}
            <div>
                <TextField label='Room Name' onChange={(e)=> {
                    setRoomName(e.currentTarget.value);
                }}/>
                <Button variant='contained' onClick={createRoom}>Create Room</Button>
            </div>
        </div>
    )
}