import React, { useState } from "react";
import { TextField, Button } from "@mui/material";
import { useNavigate } from "react-router-dom";

function isEmail(email: string) {
    var re = /\S+@\S+\.\S+/;
    return re.test(email);
}

export default function Auth() {
    const [login, setLogin] = useState(true);
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [pwd, setPwd] = useState('');
    const [error, setError] = useState(false);
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    const changeLogin = () => {
        setLogin(!login);
    }

    const loginApp = () => {
        const requestOptions = {
            method: "POST",
            headers: {
                "Content-type": "application/json",
            },
            body: JSON.stringify({
                username: isEmail(username)? '':username,
                pwd: pwd,
                email: isEmail(username)? username:'',
            })
        };
        
        fetch('/login', requestOptions).then((resp) => {
            if (resp.ok) {
                return resp.json();
            }
            return resp.text().then(text => { throw new Error(text) })
        }).then((data) => {
            //store jwt token
            navigate('/rooms')
        }).catch((err: Error) => {
            setError(true);
            setErrorMessage(err.message);
        });
    }

    const signUpApp = () => {
        const requestOptions = {
            method: "POST",
            headers: {
                "Content-type": "application/json",
            },
            body: JSON.stringify({
                username: username,
                pwd: pwd,
                email: email,
            })
        };
        
        fetch('/user', requestOptions).then((resp) => {
            if (resp.ok) {
                return resp.json();
            }
            return resp.text().then(text => { throw new Error(text) })
        }).then((data) => {
            loginApp();
        }).catch((err: Error) => {
            setError(true);
            setErrorMessage(err.message);
        });
    }

    if (login) {
        return (
            <div className="centerDiv flexColumn">
                <TextField label='Username or Email' onChange={(e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
                    setUsername(e.currentTarget.value)}
                }/>
                <TextField label='Password' style={{marginTop: '10px'}} type='password' onChange={(e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
                    setPwd(e.currentTarget.value)}
                }/>
                <Button variant='contained' style={{marginTop: '20px'}} onClick={loginApp}>Login</Button>
                {error? <div className='errorMessage'>{errorMessage.substring(1, errorMessage.length-2)}</div>:<></>}
                <div style={{marginTop: '20px'}}>You don't have an account? <a className="clickLink" onClick={changeLogin}>Sign up here</a></div>
            </div>
        )
    }
    return (
        <div className="centerDiv flexColumn">
            <TextField label='Email'onChange={(e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
                    setEmail(e.currentTarget.value)}
                }/>
            <TextField label='Username' style={{marginTop: '10px'}}onChange={(e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
                    setUsername(e.currentTarget.value)}
                }/>
            <TextField label='Password' style={{marginTop: '10px'}} type='password' onChange={(e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
                    setPwd(e.currentTarget.value)}
                }/>
            <Button variant='contained' style={{marginTop: '20px'}} onClick={signUpApp}>Sign up</Button>
            {error? <div className='errorMessage'>{errorMessage.substring(1, errorMessage.length-2)}</div>:<></>}
            <div style={{marginTop: '20px'}}>You already have an account? <a className="clickLink" onClick={changeLogin}>Login here</a></div>
        </div>
    )
}

