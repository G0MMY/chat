import { Button, Modal, TextField } from "@mui/material";
import React, { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import RoomList from "./RoomList";
import RoomMessages from "./RoomMessages";

export interface Room {
    id: number
    name: string
}

interface Invitation {
    id: number,
    sender: string,
    receiver: string,
    roomId: number,
    roomName: string
}

export default function Rooms() {
    const [open, setOpen] = useState(false);
    const [token, setToken] = useState('');
    const navigate = useNavigate();
    const { state } = useLocation();
    const { username } = state === null? '':state ;
    const [error, setError] = useState(false);
    const [errorMessage, setErrorMessage] = useState('');
    const [roomName, setRoomName] = useState('');
    const [rooms, setRooms] = useState<Room[]>([]);
    const [selectedRoom, setSelectedRoom] = useState<Room>();
    const [invitations, setinvitations] = useState<Invitation[]>([]);

    const logoutClick = () => {
        sessionStorage.removeItem('Token');
        navigate('/');
    }

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
            setError(false);
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
            setError(false);
            const temp = rooms;
            temp.push(data)
            setRooms(temp);
            setOpen(false);
            setRoomName('');
        }).catch((err: Error) => {
            setError(true);
            setErrorMessage(err.message);
        });
    }

    useEffect(() => {
        const tempToken = sessionStorage.getItem('Token');
        if (tempToken !== null && state !== null) {
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
        <>
            <Modal open={open} onClose={()=>{setOpen(false)}}>
               <div className="centerDiv" style={{padding: '40px', backgroundColor: 'white', textAlign:'center'}}>
                    <h1 style={{marginBottom:'60px'}}>Create a new room</h1>
                    <div style={{display: 'flex', justifyContent: 'center', alignItems: 'center'}}>
                        <TextField label='Room Name' error={error} helperText={error? 'Invalid Name': 'Enter the room name'} onChange={(e)=> {
                            setRoomName(e.currentTarget.value);
                        }}/>
                        <Button style={{marginLeft: '10px',marginTop:'-25px'}} variant='contained' onClick={createRoom}>Create Room</Button>
                    </div>
               </div>
            </Modal>
            <div style={{display:"flex"}}>
                <RoomList rooms={rooms} setOpen={setOpen}/>
                <div style={{width:'100%'}}>
                    <div style={{width:'100%', height:'55px', borderBottom:'1px solid black',display:'flex', justifyContent:'flex-end'}}>
                        <h4 style={{marginRight:'40px', cursor:'pointer'}}>
                            Invitations {invitations.length === 0? '':invitations.length}
                        </h4>
                        <h4 style={{marginRight:'40px', cursor:'pointer'}} onClick={logoutClick}>
                            Logout
                        </h4>
                    </div>
                </div>
                {selectedRoom !== undefined? <RoomMessages room={selectedRoom}/>:<></>}
            </div>
        </>
    )
}