import { List, ListItem, ListItemText } from "@mui/material"
import { Invitation } from "./Rooms"
import DoneIcon from '@mui/icons-material/Done';
import CloseIcon from '@mui/icons-material/Close';

interface Props {
    invitations: Invitation[]
    joinRoom: (id: number, invitation?: Invitation) => void
    deleteInvitation: (invitation: Invitation) => void
}

export default function Invitations({ invitations, joinRoom, deleteInvitation }: Props) {

    return (
        <div>
            <List sx={{width: '300px', display:'flex', flexDirection:'column', alignItems:'center', position: 'relative',overflow: 'auto', '& ul': { padding: 0 }}}>
                {invitations !== null? invitations.map((invitation) => (
                    <ListItem style={{textAlign:'center', borderRadius:'10px', margin: '5px', width:'250px'}} key={invitation.id}>
                        <ListItemText primary={`${invitation.sender} invited you to join the room ${invitation.roomName}`} />
                        <DoneIcon style={{cursor:'pointer',marginLeft:'10px',marginRight:'10px'}} onClick={() => {
                            joinRoom(invitation.roomId, invitation);
                        }} />
                        <CloseIcon style={{cursor:'pointer'}} onClick={()=>{deleteInvitation(invitation)}}/>
                    </ListItem>
                )): <></>}  
            </List>
        </div>
    )
}