import React, { useEffect, useMemo } from "react";
import { Room } from "./Rooms";
import { Autocomplete, List, ListItem, ListItemText, TextField } from "@mui/material";
import AddBoxIcon from '@mui/icons-material/AddBox';

export default function RoomList(props: { rooms: Room[], setOpen: React.Dispatch<React.SetStateAction<boolean>> }) {

    return (
        <div className='flexColumn' style={{borderRight:'1px solid black',height:'100vh',justifyContent:'flex-start'}}>
            <div style={{alignItems:'center',justifyContent:'space-evenly',display:'flex',width:'100%'}}>
                <h2>Rooms</h2>
                <div style={{cursor:'pointer'}} onClick={() => {props.setOpen(true)}}><AddBoxIcon/></div>
            </div>
            <Autocomplete disablePortal options={props.rooms !== null? props.rooms.map((room) => ({label:room.name})): []} sx={{ width: 250 }} renderInput={(params) => <TextField {...params} label="Search" helperText="Search for a room" />}/>
            <List sx={{width: '300px', display:'flex', flexDirection:'column', alignItems:'center', position: 'relative',overflow: 'auto', '& ul': { padding: 0 }}}>
                {props.rooms !== null? props.rooms.map((room) => (
                    <ListItem style={{textAlign:'center', border:'1px solid black', borderRadius:'10px', cursor:'pointer', margin: '5px', width:'250px'}} key={room.id}>
                        <ListItemText primary={`${room.name}`} />
                    </ListItem>
                )): <></>}  
            </List>
        </div>
    )
}
