import { useState } from "react";
import { Badge, Button, Card } from "react-bootstrap";
import { copyToClipboard, IResultPaste } from "../api/utils";
import QRCode from 'qrcode.react'

export default function GResult(props: IResultPaste) {
    const copyLink = () => {
        copyToClipboard(`${window.location.origin}/s/${props.mkey}`)
        setText('复制成功')
        setTimeout(() => {
            setText('复制链接')
        }, 2000);
    }

    const [text, setText] = useState('复制链接')

    const backToHome = () => {
        window.location.reload()
    }

    return (
        <div className="mt-5">
            <Card>
                <Card.Header as={"h5"}>Code Paste</Card.Header>
                <Card.Body>
                    <Card.Title><h4>保存成功</h4></Card.Title>
                    <Card.Body>
                        <ul>
                            <li>
                                索引：<Badge bg="primary">{props.mkey}</Badge>
                            </li>
                            <li>
                                标题：<Badge bg="secondary">{props.title ? props.title : "无"}</Badge>
                            </li>
                            <li>
                                格式： <Badge bg="info">{props.lang ? props.lang : "text"}</Badge>
                            </li>
                            <li>
                                密码：<Badge bg="warning">{props.password ? props.password : "无"}</Badge>
                            </li>
                            <li>
                                访问：<Badge bg="danger">{props.private ? "私有" : "公开"}</Badge>
                            </li>
                            <li>
                                <span>在浏览器中访问：</span>
                                <span className="m-2">
                                    <a
                                        target={"_blank"}
                                        rel="noopener noreferrer"
                                        href={`${window.location.origin}/s/${props.mkey}`} >
                                        {`${window.location.origin}/s/${props.mkey}`}
                                    </a>
                                </span>
                                <span>
                                    <Button
                                        variant="success"
                                        className="badge"
                                        onClick={copyLink}
                                    >
                                        {text}
                                    </Button>
                                </span>

                            </li>
                            <li>
                                扫描下方二维码
                                <div>
                                    <QRCode
                                        value={`${window.location.origin}/s/${props.mkey}`}
                                        size={200}
                                        fgColor="#000000"
                                    />
                                </div>
                            </li>
                        </ul>
                    </Card.Body>
                    <Button variant="primary" onClick={backToHome}>
                        返回主页
                    </Button>
                </Card.Body>
            </Card>
        </div>
    )
}