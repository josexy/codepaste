import React, { useState } from "react";
import { Button, Card, FormControl, InputGroup } from "react-bootstrap"
import GAlert from "./GAlert";
import { GPanel282 } from "./GPanel";

interface IProps {
    code: number,
    setPassword: React.Dispatch<React.SetStateAction<string>>
}

function CardPage({ code, setPassword }: IProps) {

    const [value, setValue] = useState('')

    const handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key === 'Enter') {
            click()
        }
    }
    const click = () => {
        setPassword(value)
    }

    return (
        <Card>
            <Card.Body>
                <GAlert variant="danger" show={true} text={code === 4001 ? "需要密码！" : "密码不正确！"} />
                <InputGroup className="mb-3">
                    <InputGroup.Text>请输入密码</InputGroup.Text>
                    <FormControl
                        type="password"
                        value={value}
                        onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setValue(e.target.value) }}
                        onKeyDown={handleKeyPress}
                    />
                </InputGroup>
                <Button variant="primary" onClick={click}>
                    确认
                </Button>
            </Card.Body>
        </Card>
    )
}

export default function GPasteAccess({ code, setPassword }: IProps) {
    let page = null;
    switch (code) {
        case 4000:
            page = <GAlert variant="danger" show={true} text={"Paste不存在！"} />
            break
        case 4001:
            page = <CardPage code={code} setPassword={setPassword} />
            break
        case 4002:
            page = <CardPage code={code} setPassword={setPassword} />
            break
        case 4004:
            page = <GAlert variant="danger" show={true} text={"此Paste已被保护，禁止访问！"} />
            break
        default:
            page = <GAlert variant="danger" show={true} text={"未知错误！"} />
            break
    }

    return (
        <GPanel282>
            {page}
        </GPanel282>
    )
}