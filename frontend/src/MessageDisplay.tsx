import { Message } from "./RoomMessages"

interface Props {
    message: Message
    username: string
}

export default function MessageDisplay({ message, username }:Props) {

    if (message.sender === username) {
        return (
            <div className="message" style={{display:'flex',alignItems:'center',marginRight:'30px',width:'100%',justifyContent:'flex-end'}}>
                <div className="hide">
                    {new Date(message.sendTime).toLocaleString()}
                </div>
                <div style={{borderRadius:'25px',margin:'10px',backgroundColor:'#e8d017',padding:'15px',maxWidth:'75%'}}>
                    {message.msg}
                </div>
            </div>
        )
    }

    return (
        <div style={{display:'flex', alignItems:'center', marginLeft:'30px'}}>
            <div>
                {message.sender}
            </div> 
            <div className="message" style={{borderRadius:'25px',margin:'10px',backgroundColor:'#8fd428',padding:'15px',maxWidth:'75%'}}>
                {message.msg}
            </div>
            <div className="hide">
                {new Date(message.sendTime).toLocaleString()}
            </div>
        </div>
    )
}