import React, { useEffect, useState } from "react"
import { Badge, Button, Card, Col, Container, Form, FormControl, InputGroup, Modal, OverlayTrigger, Row, Tooltip } from "react-bootstrap"
import request from "../api/request"
import { copyToClipboard, IPasteInfo, IResponseData, supportLangList } from "../api/utils"

interface IPageProps {
    paste: IPasteInfo,
}

interface IProps {
    show: boolean,
    setShowModal: React.Dispatch<React.SetStateAction<boolean>>
    onOk: () => void
}

function CustomModal({ show, setShowModal, onOk }: IProps) {
    return (
        <Modal
            show={show}
            size="sm"
            centered
            onHide={() => { setShowModal(false) }}
        >
            <Modal.Header closeButton>
                提示
            </Modal.Header>
            <Modal.Body>
                是否删除当前便利贴?
            </Modal.Body>
            <Modal.Footer>
                <Button size="sm" variant="danger" onClick={() => { setShowModal(false) }}>关闭</Button>
                <Button size="sm" variant="primary" onClick={() => { setShowModal(false); onOk(); }}>是的</Button>
            </Modal.Footer>
        </Modal>
    )
}

export default function GPaste({ paste }: IPageProps) {

    const [show, setShow] = useState(true)
    const [title, setTitle] = useState(paste.title ? paste.title : '')
    const [content, setContent] = useState(paste.content)
    const [lang, setLang] = useState(paste.lang)
    const [password, setPassword] = useState(paste.password ? paste.password : '')
    const [isPrivate, setIsPrivate] = useState(paste.private ? paste.private : false)

    const [text, setText] = useState('应用')
    const [showModal, setShowModal] = useState(false)

    useEffect(() => {
        setShow(true)
        setShowModal(false)
        setTitle(paste.title ? paste.title : '')
        setContent(paste.content)
        setLang(paste.lang)
        setPassword(paste.password ? paste.password : '')
        setIsPrivate(paste.private ? paste.private : false)
    }, [paste])

    const updatePaste = () => {
        if (content.length === 0) {
            setText('内容不能为空！')
            setTimeout(() => {
                setText('应用')
            }, 2000);
            return
        }
        request.put<IResponseData<any>>(`/auth/pastes/${paste.id}`, {
            data: {
                title: title.trim(),
                content: content,
                lang: lang,
                password: password.trim(),
                'private': isPrivate
            }
        }).then(res => {
            if (res.data.code === 1000) {
                setText('成功！')
                setTimeout(() => {
                    setText('应用')
                }, 2000);
            } else if (res.data.code === 1004) {
                setText('请检查信息格式是否正确！')
                setTimeout(() => {
                    setText('应用')
                }, 2000);
            }
        }).catch(err => {
            console.log(err)
            if (err.response && err.response.status === 401) {
                setText('登录回话已过期，需要重新登录！')
            }
        })
    }

    const deletePaste = () => {
        setShowModal(true)
    }

    const onOk = () => {
        setShow(false)
        request.delete<IResponseData<any>>(`/auth/pastes/${paste.id}`).catch(err => {
            console.log(err)
        })
    }

    const copyLink = () => {
        copyToClipboard(`${window.location.origin}/s/${paste.key}`)
    }

    return (
        show ?
            <div className="mb-3">
                <Card>
                    <Card.Header>
                        <div className="d-flex">
                            <div>
                                <span className="fw-bold">{title ? title : "[无标题]"}</span>
                                <Badge className="ms-1" bg={"success"}>索引: {paste.key}</Badge>
                                <Badge className="ms-1" bg={"danger"}>访问: {isPrivate ? "私有" : "公开"}</Badge>
                                <Badge className="ms-1" bg={"warning"}>密码: {password ? password : "无"}</Badge>
                                <Badge className="ms-1" bg={"primary"}>格式: {lang ? lang : "text"}</Badge>
                            </div>
                            <div className="ms-auto">
                                <Button size="sm" variant="link" onClick={copyLink}>复制链接</Button>
                                <Button size="sm" variant="link" >
                                    <a href={`${window.location.origin}/s/${paste.key}`} target={"_blank"} rel={"noreferrer"} >访问</a>
                                </Button>
                                <Button size="sm" variant="link" onClick={deletePaste}>删除</Button>
                            </div>
                        </div>
                    </Card.Header>
                    <Card.Body>
                        <Container>
                            <Row>
                                <Col>
                                    <InputGroup className="mb-2" size="sm">
                                        <OverlayTrigger placement="left"
                                            overlay={<Tooltip>标题长度小于25</Tooltip>}>
                                            <InputGroup.Text>标题</InputGroup.Text>
                                        </OverlayTrigger>
                                        <FormControl
                                            type="text"
                                            placeholder="无标题"
                                            value={title}
                                            onChange={(e: any) => { setTitle(e.target.value) }}
                                            maxLength={25}
                                        />
                                    </InputGroup>
                                    <InputGroup className="mb-2" size="sm">
                                        <InputGroup.Text>格式</InputGroup.Text>
                                        <Form.Select
                                            defaultValue={lang}
                                            onChange={(e: React.ChangeEvent<HTMLSelectElement>) => { setLang(e.target.value) }}
                                        >
                                            {
                                                supportLangList.map((val, index) =>
                                                    <option
                                                        key={index}
                                                        value={val[0]}
                                                    >
                                                        {val[1][1]}
                                                    </option>
                                                )
                                            }
                                        </Form.Select>
                                    </InputGroup>
                                    <InputGroup className="mb-2" size="sm">
                                        <OverlayTrigger placement="left"
                                            overlay={<Tooltip>密码长度小于20</Tooltip>}>
                                            <InputGroup.Text>密码</InputGroup.Text>
                                        </OverlayTrigger>
                                        <FormControl
                                            type="text"
                                            placeholder="无密码"
                                            value={password}
                                            onChange={(e: any) => { setPassword(e.target.value) }}
                                            maxLength={20}
                                        />
                                    </InputGroup>
                                    <InputGroup className="mb-2" size="sm">
                                        <Form.Check
                                            inline
                                            type="switch"
                                            label="私有"
                                            checked={isPrivate}
                                            onChange={(e: any) => { setIsPrivate(e.target.checked) }}
                                        />
                                    </InputGroup>
                                    <Form.Control
                                        className="mb-2"
                                        as="textarea"
                                        rows={10}
                                        value={content}
                                        onChange={(e: any) => { setContent(e.target.value) }}
                                    />
                                    <Button
                                        variant="outline-success"
                                        size="sm"
                                        onClick={updatePaste}
                                    >{text}
                                    </Button>
                                </Col>
                            </Row>
                        </Container>
                    </Card.Body>
                </Card>
                <CustomModal key="" show={showModal} setShowModal={setShowModal} onOk={onOk} />
            </div>
            : <></>
    )
}