import React, { useEffect, useMemo, useState } from "react";
import { Room } from "./Rooms";
import { Autocomplete, List, ListItem, ListItemText, TextField } from "@mui/material";
import AddBoxIcon from '@mui/icons-material/AddBox';
import SearchBar from "material-ui-search-bar";

export default function RoomList(props: { rooms: Room[], setOpen: React.Dispatch<React.SetStateAction<boolean>>, setRooms: React.Dispatch<React.SetStateAction<Room[]>> }) {

    const [roomSearch, setRoomSearch] = useState('');
    const [showRooms, setShowRooms] = useState<Room[]>([]);

    const handleSelectClick = (room: Room, index: number) => {
        const newRooms = props.rooms.slice();

        if (props.rooms.length !== showRooms.length) {
            let counter = 0;
            let found = 0;
            for (let i=0;i<props.rooms.length;i++) {
                if (!props.rooms[i].name.includes(roomSearch)) {
                   counter++;
                } else {
                    if (found === index) {
                        index += counter
                        break;
                    }
                    found++
                }
            }
        }

        newRooms.splice(index, 1);
        newRooms.unshift(room);
        props.setRooms(newRooms);
    }

    const handleRoomSearch = () => {
        if (roomSearch === '') {
            setShowRooms(props.rooms);

            return;
        }

        const newRooms: Room[] = [];

        props.rooms.forEach((room) => {
            if (room.name.includes(roomSearch)) {
                newRooms.push(room);
            }
        })

        setShowRooms(newRooms);
    }

    const handleCancelSearch = () => {
        setShowRooms(props.rooms);
        setRoomSearch('');
    }

    useEffect(() => {
        if (roomSearch !== '') {
            handleRoomSearch();
        } else {
            setShowRooms(props.rooms);
        }
    }, [props.rooms])

    return (
        <div className='flexColumn' style={{borderRight:'1px solid black',height:'100vh',justifyContent:'flex-start'}}>
            <div style={{alignItems:'center',justifyContent:'space-evenly',display:'flex',width:'100%'}}>
                <h2>Rooms</h2>
                <div style={{cursor:'pointer'}} onClick={() => {props.setOpen(true)}}><AddBoxIcon/></div>
            </div>
            <div style={{marginLeft:'20px',marginRight:'20px'}}>
                <SearchBar value={roomSearch} onChange={(newValue) => setRoomSearch(newValue)} onRequestSearch={handleRoomSearch} onCancelSearch={handleCancelSearch}/>
            </div>
            <List sx={{width: '300px', display:'flex', flexDirection:'column', alignItems:'center', position: 'relative',overflow: 'auto', '& ul': { padding: 0 }}}>
                {showRooms !== null? showRooms.map((room, i) => (
                    <ListItem onClick={() => {handleSelectClick(room, i)}} className={i===0? "roomSelected":"roomList"} style={{textAlign:'center', borderRadius:'10px', cursor:'pointer', margin: '5px', width:'250px'}} key={room.id}>
                        <ListItemText primary={`${room.name}`} />
                    </ListItem>
                )): <></>}  
            </List>
        </div>
    )
}
