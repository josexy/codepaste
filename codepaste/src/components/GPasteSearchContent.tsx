import { useEffect, useState } from "react";
import { Badge, Button, Card, Spinner } from "react-bootstrap";
import { useParams } from "react-router-dom";
import request from "../api/request";
import { copyToClipboard, IPasteInfo, IResponseData, supportLangMap } from "../api/utils";
import SyntaxHighlighter from 'react-syntax-highlighter';
import { atomOneLight } from 'react-syntax-highlighter/dist/esm/styles/hljs';
import GPasteAccess from "./GPasteAccess";
import LocalStorage from "../api/local";

export default function GPasteSearchContent() {
    const params = useParams()

    const [code, setCode] = useState(1000)
    const [paste, setPaste] = useState<IPasteInfo>()
    const [password, setPassword] = useState('')

    const getPaste = (key?: string, password?: string) => {

        let param = { password: password }
        let url = ''
        if (LocalStorage.getUsername() === '') {
            url = `/paste/${key}`
        } else {
            url = `/auth/pastes/${key}`
        }

        request.get<IResponseData<IPasteInfo>>(url, { params: param }).then(res => {
            setCode(res.data.code)
            switch (res.data.code) {
                case 1000:
                    if (res.data.data)
                        setPaste(res.data.data)
                    break
                case 4000: // paste不存在
                case 4001: // 需要密码
                case 4002: // 密码不正确
                case 4004: // 私有
                default:
                    setPaste(undefined)
                    break
            }
        }).catch(err => {
            console.log(err);
        })
    }

    useEffect(() => {
        getPaste(params.key, password.trim())
    }, [params.key, password])

    const copyPaste = () => {
        if (paste)
            copyToClipboard(paste.content)
    }

    return (
        <>
            {
                paste ?
                    <Card>
                        <Card.Header>
                            <div className="d-flex">
                                <div className="fw-bold">
                                    {paste.title ? paste.title : "[无标题]"}
                                    <Badge className="ms-2" bg="warning">
                                        {paste.lang ? paste.lang : "text"}
                                    </Badge>
                                </div>
                                <div className="ms-auto">
                                    <Button size="sm" variant="link" onClick={copyPaste}>
                                        复制
                                    </Button>
                                </div>
                            </div>
                        </Card.Header>
                        <Card.Body>
                            <SyntaxHighlighter
                                customStyle={{ backgroundColor: "white" }}
                                showLineNumbers={true}
                                startingLineNumber={1}
                                language={paste.lang ? supportLangMap.get(paste.lang)?.at(0) : "text"}
                                style={atomOneLight}
                                wrapLines={true}
                            >
                                {paste.content}
                            </SyntaxHighlighter>
                        </Card.Body>
                    </Card>
                    :
                    code === 1000 ?
                        <Spinner animation="border" variant="primary" />
                        : <GPasteAccess code={code} setPassword={setPassword} />
            }
        </>
    )
}