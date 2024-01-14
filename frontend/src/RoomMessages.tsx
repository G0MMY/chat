import React, { useEffect, useState } from "react";
import { Room } from "./Rooms";
import Message from "./MessageDisplay";
import MessageDisplay from "./MessageDisplay";
import { Button, List, ListItem, TextField } from "@mui/material";

interface Props {
    room: Room
    token: string
    username: string
    lastMessage: MessageEvent<any> | null
}

export interface Message {
    id: number
    roomId: number
    sender: string
    sendTime: number
    msg: string
}

export default function RoomMessages({ room, token, username, lastMessage }:Props) {

    const [messages, setMessages] = useState<Message[]>([]);

    const getMessages = () => {
        const requestOptions = {
            method: "GET",
            headers: {
                "Content-type": "application/json",
                "Token": token
            }
        };
        fetch(`/messages/${room.id}`, requestOptions).then((resp) => {
            if (resp.ok) {
                return resp.json();
            }
            return resp.text().then(text => { throw new Error(text) })
        }).then((data) => {
            setMessages(data.items);
        }).catch((err: Error) => {
            console.log(err);
        });
    }

    useEffect(() => {
        getMessages();
    }, [])

    useEffect(() => {
        const messageList = document.getElementById("messageList")!;
        messageList.scrollTop = messageList.scrollHeight;
    }, [messages])

    useEffect(() => {
        if (lastMessage !== null) {
            setMessages((prev)=> prev.concat(JSON.parse(lastMessage.data)));
        }
    }, [lastMessage])

    return (
        <List id="messageList" sx={{width: '100%', position: 'relative',overflow: 'auto', '& ul': { padding: 0 }}}>
            {messages !== null? messages.map((message) => (
                <ListItem key={message.id}>
                    <MessageDisplay message={message} username={username}/>
                </ListItem>
            )): <></>}
        </List>
    )
}