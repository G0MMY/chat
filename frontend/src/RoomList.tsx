import React, { useEffect, useMemo } from "react";
import { Room } from "./Rooms";
import { Autocomplete, List, ListItem, ListItemText, TextField } from "@mui/material";
import AddBoxIcon from '@mui/icons-material/AddBox';

export default function RoomList(props: { rooms: Room[] }) {

    const createRoomClick = () => {

    }

    // const createRoom = () => {
    //     const requestOptions = {
    //         method: "POST",
    //         headers: {
    //             "Content-type": "application/json",
    //             "Token": token
    //         },
    //         body: JSON.stringify({
    //             name: roomName
    //         })
    //     };
    //     fetch('/rooms', requestOptions).then((resp) => {
    //         if (resp.ok) {
    //             return resp.json();
    //         }
    //         return resp.text().then(text => { throw new Error(text) })
    //     }).then((data) => {
    //         setError(false);
    //         joinRoom(data.id)
    //     }).catch((err: Error) => {
    //         setError(true);
    //         setErrorMessage(err.message);
    //     });
    // }

    return (
        <div className='flexColumn'>
            <div style={{alignItems:'center',justifyContent:'space-evenly',display:'flex',width:'100%'}}>
                <h2>Rooms</h2>
                <div style={{cursor:'pointer'}} onClick={createRoomClick}><AddBoxIcon/></div>
            </div>
            <Autocomplete disablePortal options={props.rooms !== null? props.rooms.map((room) => ({label:room.name})): []} sx={{ width: 300 }} renderInput={(params) => <TextField {...params} label="Search" />}/>
            <List sx={{width: '300px',position: 'relative',overflow: 'auto', '& ul': { padding: 0 }}}>
                {props.rooms !== null? props.rooms.map((room) => (
                    <ListItem key={room.id}>
                        <ListItemText primary={`${room.name}`} />
                    </ListItem>
                )): <></>}  
            </List>
        </div>
    )
}