import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import LocalStorage from "../api/local";
import request from "../api/request";
import { IResponseData } from "../api/utils";

export default function GLogout() {

    const navigate = useNavigate()

    useEffect(() => {
        if (LocalStorage.getUsername() === '') {
            navigate("/")
            return
        }
        request.delete<IResponseData<any>>("/auth/user").then(res => {
            LocalStorage.removeToken()
            LocalStorage.removeUsername()
            navigate("/")
        }).catch(err => {
            console.log(err);
        })
    }, [navigate])

    return (
        <></>
    )
}