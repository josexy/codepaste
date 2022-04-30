import { useEffect, useState } from "react"
import { Button, Col, Form, FormControl, InputGroup, OverlayTrigger, Row, Tooltip } from "react-bootstrap"
import request from "../api/request"
import { IResponseData, IUserInfo } from "../api/utils"

interface IProps {
    showAlert: (type: string, text: string) => void
}

export default function GUser({ showAlert }: IProps) {

    const [userInfo, setUserInfo] = useState<IUserInfo>()

    const [nickname, setNickname] = useState('')
    const [password, setPassword] = useState('')

    const [pwType, setPwType] = useState('password')

    const getUserInfo = () => {
        request.get<IResponseData<IUserInfo>>("/auth/user").then(res => {
            if (res.data.code === 1000) {
                if (res.data.data) {
                    setUserInfo(res.data.data)
                    setNickname(res.data.data.nickname)
                }
            }
        }).catch(err => {
            console.log(err)
        })
    }

    useEffect(() => {
        getUserInfo()
    }, [])

    const update = () => {
        let n = nickname.trim()
        let p = password.trim()
        if (n.length === 0 || p.length === 0) {
            showAlert("warning", "昵称和密码不能为空！")
            return
        }

        request.put<IResponseData<any>>("/auth/user", {
            data: {
                nickname: n,
                password: p,
            }
        }).then(res => {
            if (res.data.code === 3002) {
                showAlert("warning", "用户昵称已经存在！")
            } else if (res.data.code === 1000) {
                if (userInfo)
                    userInfo.nickname = nickname
                showAlert("success", "更新成功！")
            } else if (res.data.code === 1004) {
                showAlert("warning", "请检查信息格式是否正确！")
            }
        }).catch(err => {
            console.log(err)
            if (err.response && err.response.status === 401) {
                showAlert('danger', '登录回话已过期，需要重新登录！')
            }
        })
    }

    return (
        <>
            {
                userInfo &&
                <>
                    <Row>
                        <Col>
                            <InputGroup className="mb-2">
                                <InputGroup.Text>
                                    用户名
                                </InputGroup.Text>
                                <FormControl type="text"
                                    readOnly
                                    minLength={5}
                                    maxLength={20} required
                                    value={userInfo.username} disabled />
                            </InputGroup>
                        </Col>
                    </Row>
                    <Row>
                        <Col>
                            <InputGroup className="mb-2">
                                <OverlayTrigger placement="left"
                                    overlay={<Tooltip>昵称必须介于3到20个字符之间</Tooltip>}>
                                    <InputGroup.Text>昵称</InputGroup.Text>
                                </OverlayTrigger>
                                <FormControl type="text"
                                    minLength={3}
                                    maxLength={20} required
                                    value={nickname} onChange={(e: any) => setNickname(e.target.value)} />
                            </InputGroup>
                        </Col>
                    </Row>
                    <Row>
                        <Col>
                            <InputGroup className="mb-2">
                                <OverlayTrigger placement="left"
                                    overlay={<Tooltip>密码必须介于8到20个字符之间</Tooltip>}>
                                    <InputGroup.Text>密码</InputGroup.Text>
                                </OverlayTrigger>
                                <FormControl
                                    type={pwType}
                                    minLength={8}
                                    maxLength={20} required
                                    value={password}
                                    onChange={(e: any) => setPassword(e.target.value)} />
                                <InputGroup.Text>
                                    <Form.Check
                                        type="switch"
                                        defaultChecked
                                        onClick={(e: any) => {
                                            if (e.target.checked) {
                                                setPwType("password")
                                            } else {
                                                setPwType("text")
                                            }
                                        }}
                                    />
                                </InputGroup.Text>
                            </InputGroup>
                        </Col>
                    </Row>
                    <Row>
                        <Col>
                            <Button className="mb-2" variant="success" onClick={update}>
                                更新
                            </Button>
                        </Col>
                    </Row>
                </>
            }
        </>
    )
}