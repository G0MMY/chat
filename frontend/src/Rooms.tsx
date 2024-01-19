import { Button, Modal, TextField } from "@mui/material";
import React, { useEffect, useMemo, useRef, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import RoomList from "./RoomList";
import RoomMessages from "./RoomMessages";
import useWebSocket from "react-use-websocket";
import Invitations from "./Invitations";
import AddBoxIcon from '@mui/icons-material/AddBox';

export interface Room {
    id: number
    name: string
}

export interface Invitation {
    id: number,
    sender: string,
    receiver: string,
    roomId: number,
    roomName: string
}

export default function Rooms() {

    const [message, setMessage] = useState('');
    const [openCreateRoom, setOpenCreateRoom] = useState(false);
    const [openInvitations, setOpenInvitations] = useState(false);
    const [addInvitation, setAddInvitation] = useState(false);
    const [token, setToken] = useState('');
    const navigate = useNavigate();
    const { state } = useLocation();
    const { username } = state === null? '':state ;
    const [error, setError] = useState(false);
    const [errorMessage, setErrorMessage] = useState('');
    const [roomName, setRoomName] = useState('');
    const [usernameInvitation, setUsernameInvitation] = useState('');
    const [rooms, setRooms] = useState<Room[]>([]);
    const [invitations, setInvitations] = useState<Invitation[]>([]);
    const [socketUrl, setSocketUrl] = useState(`ws://127.0.0.1:8080/ws/rooms/${username}`);
    const roomsWebsocket = useWebSocket(socketUrl);
    const invitationsWebsocket = useWebSocket(`ws://127.0.0.1:8080/ws/invitations/${username}`);

    const selectedRoom: Room = useMemo<Room>(() => {
        if (rooms !== null) {
            return rooms[0];
        }

        return {id: 0, name: ''}
    }, [rooms]);

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
            getInvitations(currentToken);
        }).catch((err: Error) => {
            setError(true);
            setErrorMessage(err.message);
        });
    }

    const getInvitations = (currentToken?: string) => {
        const requestOptions = {
            method: "GET",
            headers: {
                "Content-type": "application/json",
                "Token": currentToken? currentToken : token
            }
        };
        fetch(`/invitations/${username}`, requestOptions).then((resp) => {
            if (resp.ok) {
                return resp.json();
            }
            return resp.text().then(text => { throw new Error(text) })
        }).then((data) => {
            setError(false);
            setInvitations(data !== null? data : []);
        }).catch((err: Error) => {
            setError(true);
            setErrorMessage(err.message);
        });
    }

    const createInvitation = () => {
        console.log(JSON.stringify({
            sender: username,
            receiver: usernameInvitation,
            roomName: roomName
        }))
        invitationsWebsocket.sendMessage(JSON.stringify({
            sender: username,
            receiver: usernameInvitation,
            roomName: roomName
        }));
        setOpenInvitations(false);
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
            joinRoom(data.id);
        }).catch((err: Error) => {
            setError(true);
            setErrorMessage(err.message);
        });
    }

    const deleteInvitation = (invitation: Invitation, changeSocket: boolean) => {
        const requestOptions = {
            method: "DELETE",
            headers: {
                "Content-type": "application/json",
                "Token": token
            }
        };
        fetch(`/invitations/${invitation.id}`, requestOptions).then((resp) => {
            if (resp.ok) {
                return resp.json();
            }
            return resp.text().then(text => { throw new Error(text) })
        }).then((data) => {
            setInvitations((prev) => (prev.splice(prev.indexOf(invitation), 1)));
            if (changeSocket) {
                setSocketUrl(`ws://127.0.0.1:8080/ws/rooms/${username}`);
            }
        })
    }

    const joinRoom = (id: number, invitation?: Invitation) => {
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
            setRooms((prev) => (prev.concat(data)));
            setOpenCreateRoom(false);
            setRoomName('');
            setSocketUrl('');
            setTimeout(()=>{
                if (invitation !== undefined) {
                    deleteInvitation(invitation, true);
                } else {
                    setSocketUrl(`ws://127.0.0.1:8080/ws/rooms/${username}`);
                }
            }, 10)
        }).catch((err: Error) => {
            setError(true);
            setErrorMessage(err.message);
        });
    }

    const handleSendClick = () => {
        if (message !== '') {
            roomsWebsocket.sendMessage(JSON.stringify({
                roomId: selectedRoom.id,
                sender: username,
                sendTime: Date.now(),
                msg: message
            }))

            setMessage('');
        }
    }

    const handleMessageChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        setMessage(e.target.value);
    }

    useEffect(() => {
        if (invitationsWebsocket.lastMessage !== null) {
            setInvitations((prev)=> prev === null? [JSON.parse(invitationsWebsocket.lastMessage!.data)] : prev.concat(JSON.parse(invitationsWebsocket.lastMessage!.data)));
        }
    }, [invitationsWebsocket.lastMessage])

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
            <Modal open={openCreateRoom} onClose={()=>{setOpenCreateRoom(false)}}>
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
            <Modal open={openInvitations} onClose={()=>{setOpenInvitations(false)}}>
               <div className="centerDiv" style={{padding: '40px', backgroundColor: 'white',alignItems:'flex-start'}}>
                    <div style={{display:'flex',alignItems:'center',justifyContent:'center'}}>
                        <h1 style={{marginRight:'60px'}}>Invitations</h1>
                        <div style={{cursor:'pointer'}} onClick={()=>{setAddInvitation(true)}} ><AddBoxIcon/></div>
                    </div>
                    {addInvitation? 
                        <div style={{display: 'flex', flexDirection:'column', justifyContent: 'center', alignItems: 'center',marginTop:'20px'}}>
                            <TextField label='Room Name' error={error} helperText={error? 'Invalid Name': 'Enter the room name'} onChange={(e)=> {
                                setRoomName(e.currentTarget.value);
                            }}/>
                            <TextField label='Username' error={error} helperText={error? 'Invalid Username': 'Enter the username'} onChange={(e)=> {
                                setUsernameInvitation(e.currentTarget.value);
                            }}/>
                            <Button style={{marginTop:'20px'}} variant='contained' onClick={createInvitation}>Create Invitation</Button>
                        </div> : 
                        <Invitations invitations={invitations} joinRoom={joinRoom} deleteInvitation={deleteInvitation}/>
                    }
               </div>
            </Modal>
            <div style={{display:"flex"}}>
                <RoomList rooms={rooms} setOpen={setOpenCreateRoom} setRooms={setRooms}/>
                <div className="flexColumn" style={{width:'100%',maxHeight:'100vh', justifyContent:'flex-start'}}>
                    <div style={{width:'100%', height:'55px', borderBottom:'1px solid black',display:'flex', justifyContent:'flex-end'}}>
                        <h4 style={{marginRight:'40px', cursor:'pointer', display:'flex'}} onClick={()=>{setOpenInvitations(true)}}>
                            Invitations &nbsp; <div style={{color:'red'}}>{invitations.length === 0? '':invitations.length}</div>
                        </h4>
                        <h4 style={{marginRight:'40px', cursor:'pointer'}} onClick={logoutClick}>
                            Logout
                        </h4>
                    </div>
                    {selectedRoom !== undefined? <RoomMessages lastMessage={roomsWebsocket.lastMessage} room={selectedRoom} token={token} username={username}/>:<></>}
                    <div style={{display:'flex',alignItems:'center', width:'100%',marginTop:'auto'}}>
                        <div style={{margin:'5px',marginLeft:'20px',width:'100%'}}>
                            <TextField value={message} onChange={(e) => handleMessageChange(e)} fullWidth multiline label="Message" />
                        </div>
                        <Button sx={{marginRight:'20px'}} variant="contained" onClick={handleSendClick}>Send</Button>
                    </div>
                </div>
            </div>
        </>
    )
}