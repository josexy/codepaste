import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Button, Col, Form, FormControl, InputGroup, OverlayTrigger, Row, Tooltip } from "react-bootstrap";
import { GPanel444 } from "./GPanel";
import {
    faIdCard,
    faLock,
    faUserAlt,
} from "@fortawesome/free-solid-svg-icons";
import React, { useState } from "react";
import request from "../api/request";
import { IResponseData, IUserInfo } from "../api/utils";
import { useNavigate } from "react-router-dom";
import GAlert from "./GAlert";

export default function GRegister() {

    const [username, setUsername] = useState('')
    const [nickname, setNickname] = useState('')
    const [password, setPassword] = useState('')
    const [validated, setValidated] = useState(false)

    const [pwType, setPwType] = useState('password')

    const [type, setType] = useState('danger')
    const [show, setShow] = useState(false)
    const [message, setMessage] = useState('')

    const [userInfo, setUserInfo] = useState<IUserInfo>()

    const navigate = useNavigate()

    const showAlert = (type: string, text: string) => {
        setShow(true)
        setType(type)
        setMessage(text)
        setTimeout(() => {
            setShow(false)
        }, 2000);
    }

    const createUser = () => {
        request.post<IResponseData<IUserInfo>>("/user/register", {
            data: {
                username: username,
                nickname: nickname,
                password: password
            }
        }).then(res => {
            switch (res.data.code) {
                case 1000:
                    if (res.data.data) {
                        setUserInfo(res.data.data)
                    }
                    showAlert('success', "注册成功，正在跳转到登录界面！")
                    setTimeout(() => {
                        navigate("/login")
                    }, 2000);
                    break
                case 3001:
                    showAlert('warning', "用户已存在！")
                    setUserInfo(undefined)
                    break
                case 3002:
                    showAlert('warning', "用户昵称已存在！")
                    setUserInfo(undefined)
                    break
                default:
                    showAlert('danger', "未知错误！")
                    setUserInfo(undefined)
                    break
            }
        }).catch(err => {
            console.log(err)
            showAlert('danger', err)
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
        createUser()
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
                                overlay={<Tooltip>昵称必须介于3到20个字符之间</Tooltip>}>
                                <InputGroup.Text>
                                    <FontAwesomeIcon icon={faIdCard} />
                                </InputGroup.Text>
                            </OverlayTrigger>
                            <FormControl type="text" minLength={3} maxLength={20} required value={nickname}
                                placeholder="请输入昵称" onChange={(e: any) => { setNickname(e.target.value) }} />
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
                        <Button disabled={!!userInfo} type="submit" variant="primary">
                            注册
                        </Button>
                    </Col>
                </Row>
            </Form>
            {<GAlert show={show} variant={type} text={message} />}
        </GPanel444 >
    )
}