import { useEffect, useState } from "react";
import { Tab, Tabs } from "react-bootstrap";
import { useNavigate } from "react-router-dom";
import LocalStorage from "../api/local";
import GAlert from "./GAlert";
import { GPanel363 } from "./GPanel";
import GPastes from "./GPastes";
import GUser from "./GUser";

export default function GUserManager() {

    const [type, setType] = useState('danger')
    const [show, setShow] = useState(false)
    const [message, setMessage] = useState('')
    const [login, setLogin] = useState(false)

    const [totalPastes, setTotalPastes] = useState(0)

    const navigate = useNavigate()

    const showAlert = (type: string, text: string) => {
        setShow(true)
        setType(type)
        setMessage(text)
        setTimeout(() => {
            setShow(false)
        }, 2000);
    }

    useEffect(() => {
        if (LocalStorage.getUsername() !== '') {
            setLogin(true)
        } else {
            setLogin(false)
            navigate("/login")
        }
    }, [navigate])

    return (
        <>
            {login &&
                <GPanel363>
                    <GAlert show={show} variant={type} text={message} />
                    <Tabs defaultActiveKey="user" className="mt-3 mb-3">
                        <Tab eventKey="user" title={"用户管理"}>
                            <GUser showAlert={showAlert} />
                        </Tab>
                        <Tab eventKey="pastes" title={`便利贴管理(${totalPastes ? totalPastes : 0})`}>
                            <GPastes setTotalPastes={setTotalPastes} />
                        </Tab>
                    </Tabs>
                </GPanel363>
            }
        </>
    )
}