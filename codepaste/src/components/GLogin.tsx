import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Button, Col, Form, FormControl, InputGroup, OverlayTrigger, Row, Tooltip } from "react-bootstrap";
import { GPanel444 } from "./GPanel";
import {
    faLock,
    faSignInAlt,
    faUserAlt,
} from "@fortawesome/free-solid-svg-icons";
import { useState } from "react";
import request from "../api/request";
import { IResponseData, IUserTokenInfo } from "../api/utils";
import { useNavigate } from "react-router-dom";
import GAlert from "./GAlert";
import LocalStorage from "../api/local";

export default function GLogin() {

    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [validated, setValidated] = useState(false)

    const [pwType, setPwType] = useState('password')

    const [type, setType] = useState('danger')
    const [show, setShow] = useState(false)
    const [message, setMessage] = useState('')

    const [tokenInfo, setTokenInfo] = useState<IUserTokenInfo>()

    const navigate = useNavigate()

    const showAlert = (type: string, text: string) => {
        setShow(true)
        setType(type)
        setMessage(text)
        setTimeout(() => {
            setShow(false)
        }, 2000);
    }

    const login = () => {
        const u = username.trim()
        const p = password.trim()

        request.post<IResponseData<IUserTokenInfo>>("/user/login", {
            data: {
                username: u,
                password: p
            }
        }).then(res => {
            switch (res.data.code) {
                case 1000:
                    if (res.data.data) {
                        setTokenInfo(res.data.data)
                        // 记录token
                        LocalStorage.setUsername(username)
                        LocalStorage.setToken(res.data.data.token)
                    }
                    showAlert('success', "登录成功，正在跳转到主界面！")
                    setTimeout(() => {
                        navigate("/user")
                    }, 2000)
                    break
                default:
                    showAlert('danger', "未知错误！")
                    setTokenInfo(undefined)
                    break
            }
        }).catch(err => {
            if (err.response && err.response.status === 401) {
                showAlert('danger', '登录失败！')
            }
        })
    }

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()
        event.stopPropagation()

        const form = event.currentTarget
        setValidated(true)

        if (!form.checkValidity()) {
            return
        }
        login()
    }

    return (
        <GPanel444>
            <Form noValidate validated={validated} onSubmit={handleSubmit}>
                <Row className="mt-5">
                    <Col>
                        <InputGroup className="mb-2">
                            <OverlayTrigger placement="left"
                                overlay={<Tooltip>用户名必须介于5到20个字符之间</Tooltip>}>
                                <InputGroup.Text>
                                    <FontAwesomeIcon icon={faUserAlt} />
                                </InputGroup.Text>
                            </OverlayTrigger>
                            <FormControl type="text" minLength={5} maxLength={20} required value={username}
                                placeholder="请输入用户名" onChange={(e: any) => { setUsername(e.target.value) }} />
                        </InputGroup>
                    </Col>
                </Row>
                <Row>
                    <Col>
                        <InputGroup className="mb-2">
                            <OverlayTrigger placement="left"
                                overlay={<Tooltip>密码必须介于8到20个字符之间</Tooltip>}>
                                <InputGroup.Text onClick={(e: any) => {
                                    if (pwType === 'password') {
                                        setPwType('text')
                                    } else {
                                        setPwType('password')
                                    }
                                }}>
                                    <FontAwesomeIcon icon={faLock} />
                                </InputGroup.Text>
                            </OverlayTrigger>

                            <FormControl type={pwType} minLength={8} maxLength={20} required value={password}
                                placeholder="请输入密码" onChange={(e: any) => { setPassword(e.target.value) }} />
                        </InputGroup>
                    </Col>
                </Row>
                <Row className="mb-3">
                    <Col className="d-grid gap-2">
                        <Button disabled={!!tokenInfo} type="submit" variant="success">
                            登录
                            <FontAwesomeIcon className="ms-1" icon={faSignInAlt} />
                        </Button>
                    </Col>
                </Row>
            </Form>
            {<GAlert show={show} variant={type} text={message} />}
        </GPanel444 >
    )
}