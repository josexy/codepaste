import { useEffect, useState } from "react";
import { Button, Col, Form, FormControl, InputGroup, OverlayTrigger, Row, Tooltip } from "react-bootstrap";
import { IPasteInfo, IResponseData, supportExpireTime, supportLangList } from "../api/utils";
import { GPanel282 } from "./GPanel";
import LocalStorage from "../api/local";
import request from "../api/request";
import GAlert from "./GAlert";
import GResult from "./GResult";
import AceEditor from "react-ace";
import "ace-builds/src-noconflict/mode-text";
import "ace-builds/src-noconflict/mode-golang";
import "ace-builds/src-noconflict/mode-c_cpp";
import "ace-builds/src-noconflict/mode-java";
import "ace-builds/src-noconflict/mode-javascript";
import "ace-builds/src-noconflict/mode-python";
import "ace-builds/src-noconflict/mode-rust";
import "ace-builds/src-noconflict/mode-ruby";
import "ace-builds/src-noconflict/mode-html";
import "ace-builds/src-noconflict/mode-markdown";
import "ace-builds/src-noconflict/mode-json";
import "ace-builds/src-noconflict/mode-sql";
import "ace-builds/src-noconflict/mode-php";
import "ace-builds/src-noconflict/theme-tomorrow";
import "ace-builds/src-noconflict/ext-language_tools";

export default function GHome() {

    const [title, setTitle] = useState('')
    const [lang, setLang] = useState('text')
    const [expireTime, setExpireTime] = useState(0)
    const [password, setPassword] = useState('')
    const [content, setContent] = useState('')
    const [isPrivate, setIsPrivate] = useState(false)

    const [pwType, setPwType] = useState('password')
    const [paste, setPaste] = useState<IPasteInfo>()
    const [message, setMessage] = useState('')
    const [show, setShow] = useState(false)
    const [type, setType] = useState('danger')

    const [login, setLogin] = useState(false)

    useEffect(() => {
        if (LocalStorage.getUsername() !== '') {
            setLogin(true)
        } else {
            setLogin(false)
        }
    }, [])


    const showAlert = (type: string, message: string) => {
        setType(type)
        setShow(true)
        setMessage(message)
        setTimeout(() => {
            setShow(false);
        }, 2000);
    }

    // ????????????????????????????????????
    const savePaste = () => {
        if (content.length === 0) {
            showAlert('danger', '?????????????????????')
            return
        }

        let url = ''
        if (login) {
            url = '/auth/pastes'
        } else {
            url = '/paste'
        }
        request.post<IResponseData<IPasteInfo>>(url, {
            data: {
                title: title.trim(),
                content: content,
                lang: lang,
                password: password.trim(),
                'private': isPrivate,
                expire_second: expireTime
            }
        }).then(res => {
            // ??????
            if (res.data.code === 1000) {
                showAlert('success', '???????????????')
                if (res.data.data)
                    setPaste(res.data.data)
            } else if (res.data.code === 4003) {
                showAlert('danger', '?????????????????????????????????')
            } else if (res.data.code === 1004) {
                showAlert('warning', '????????????????????????????????????')
            }
        }).catch(err => {
            console.log(err)
            if (err.response && err.response.status === 401) {
                showAlert('danger', '?????????????????????????????????????????????')
            }
        })
    }

    return (
        <GPanel282>
            {
                paste ?
                    <GResult mkey={paste.key} title={paste.title}
                        lang={paste.lang} password={paste.password} private={paste.private} /> :
                    <div>
                        <Row>
                            <Col md={7} lg={5}>
                                <InputGroup className="mb-3">
                                    <OverlayTrigger placement="left"
                                        overlay={<Tooltip>??????????????????25</Tooltip>}>
                                        <InputGroup.Text>??????</InputGroup.Text>
                                    </OverlayTrigger>
                                    <FormControl
                                        type="text"
                                        placeholder="?????????"
                                        maxLength={25}
                                        value={title} onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setTitle(e.target.value) }}
                                    />
                                </InputGroup>
                            </Col>
                        </Row>

                        <Row>
                            <Col md={7} lg={5}>
                                <InputGroup className="mb-3">
                                    <InputGroup.Text>??????</InputGroup.Text>
                                    <Form.Select value={lang} onChange={(e: React.ChangeEvent<HTMLSelectElement>) => { setLang(e.target.value) }}>
                                        {
                                            supportLangList.map((val, index) =>
                                                <option key={index}
                                                    defaultChecked={index === 0}
                                                    value={val[0]}
                                                >
                                                    {val[1][1]}
                                                </option>
                                            )
                                        }
                                    </Form.Select>
                                </InputGroup>
                            </Col>
                        </Row>
                        <Row>
                            <Col md={7} lg={5}>
                                <InputGroup className="mb-3">
                                    <OverlayTrigger placement="left"
                                        overlay={<Tooltip>??????????????????20</Tooltip>}>
                                        <InputGroup.Text>??????</InputGroup.Text>
                                    </OverlayTrigger>
                                    <FormControl
                                        placeholder="?????????"
                                        maxLength={20}
                                        type={pwType}
                                        value={password}
                                        onChange={(e: any) => { setPassword(e.target.value) }}
                                    />
                                    <InputGroup.Text>
                                        <Form.Check
                                            type="switch"
                                            defaultChecked={true}
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
                            <Col md={7} lg={5}>
                                <InputGroup className="mb-3">
                                    <InputGroup.Text>????????????</InputGroup.Text>
                                    <Form.Select value={expireTime}
                                        onChange={(e: React.ChangeEvent<HTMLSelectElement>) => { setExpireTime(parseInt(e.target.value, 10)) }}>
                                        {
                                            supportExpireTime.map((val, index) =>
                                                <option key={index}
                                                    defaultChecked={index === 0}
                                                    value={val[0]}
                                                >
                                                    {val[1]}
                                                </option>
                                            )
                                        }
                                    </Form.Select>
                                </InputGroup>
                            </Col>
                        </Row>
                        <Row>
                            <Col>
                                <AceEditor
                                    wrapEnabled
                                    className="mb-3 rounded-3"
                                    style={{ "border": "1px solid lightgray" }}
                                    width="100%"
                                    placeholder="???????????????"
                                    mode={lang}
                                    theme="tomorrow"
                                    name="code_content"
                                    fontSize={18}
                                    showPrintMargin={false}
                                    showGutter={true}
                                    highlightActiveLine={true}
                                    value={content}
                                    onChange={(value: string) => { setContent(value) }}
                                    onLoad={editorInstance => {
                                        editorInstance.container.style.resize = "vertical";
                                        document.addEventListener("mouseup", e => (
                                            editorInstance.resize()
                                        ));
                                    }}
                                    setOptions={{
                                        enableBasicAutocompletion: true,
                                        enableLiveAutocompletion: true,
                                        enableSnippets: false,
                                        showLineNumbers: true,
                                    }} />
                            </Col>
                        </Row>
                        <Row>
                            <Col>
                                <div className="mb-3">
                                    <Button
                                        className="me-2"
                                        variant="primary"
                                        onClick={savePaste}>
                                        ??????
                                    </Button>
                                    <Form.Check
                                        inline
                                        disabled={!login}
                                        type="switch"
                                        label="??????(?????????)"
                                        onClick={(e: any) => { setIsPrivate(e.target.checked) }}
                                    />
                                </div>
                            </Col>
                        </Row>
                        <Row>
                            <Col>
                                <GAlert show={show} variant={type} text={message} />
                            </Col>
                        </Row>
                    </div>
            }
        </GPanel282>
    )
}