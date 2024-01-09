import { Button, TextField } from "@mui/material";
import React, { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import RoomList from "./RoomList";
import RoomMessages from "./RoomMessages";

export interface Room {
    id: number
    name: string
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
    const [selectedRoom, setSelectedRoom] = useState<Room>();

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
            setError(false);
            setRooms(data.items !== null? data.items : []);
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
            setError(false);
            const temp = rooms;
            temp.push(data)
            setRooms(temp);
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
        <div style={{display:"flex"}}>
            <RoomList rooms={rooms}/>
            <RoomMessages room={selectedRoom}/>
        </div>
        // <div className="flexColumn">
        //     {/* <div style={{display: 'flex', justifyContent: 'center'}}>
        //         <TextField label='Room Name' onChange={(e)=> {
        //             setRoomName(e.currentTarget.value);
        //         }}/>
        //         <Button variant='contained' onClick={createRoom}>Create Room</Button>
        //     </div> */}
        //     {error? errorMessage: ''}
        //     <div>
        //         <RoomList rooms={rooms}/>
        //         <RoomMessages room={selectedRoom}/>
        //     </div>
        // </div>
    )
}