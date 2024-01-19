import { List, ListItem, ListItemText } from "@mui/material"
import { Invitation } from "./Rooms"
import DoneIcon from '@mui/icons-material/Done';
import CloseIcon from '@mui/icons-material/Close';

interface Props {
    invitations: Invitation[]
    joinRoom: (id: number) => void
    setInvitations: React.Dispatch<React.SetStateAction<Invitation[]>>
    token: string
}

export default function Invitations({ invitations, joinRoom, setInvitations, token }: Props) {

    const deleteInvitation = (invitation: Invitation) => {
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
            setInvitations((prev) => (prev.splice(prev.indexOf(invitation), 1)))
        })
    }

    return (
        <div>
            <List sx={{width: '300px', display:'flex', flexDirection:'column', alignItems:'center', position: 'relative',overflow: 'auto', '& ul': { padding: 0 }}}>
                {invitations !== null? invitations.map((invitation, i) => (
                    <ListItem style={{textAlign:'center', borderRadius:'10px', margin: '5px', width:'250px'}} key={invitation.id}>
                        <ListItemText primary={`${invitation.sender} invited you to join the room ${invitation.roomName}`} />
                        <DoneIcon style={{cursor:'pointer',marginLeft:'10px',marginRight:'10px'}} onClick={() => {
                            joinRoom(invitation.roomId);
                            setTimeout(() => {
                                deleteInvitation(invitation);
                            }, 10)
                        }} />
                        <CloseIcon style={{cursor:'pointer'}} onClick={()=>{deleteInvitation(invitation)}}/>
                    </ListItem>
                )): <></>}  
            </List>
        </div>
    )
}