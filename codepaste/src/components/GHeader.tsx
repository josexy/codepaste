import React, { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { Button, Container, FormControl, InputGroup, Nav, Navbar, NavDropdown } from "react-bootstrap";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSearch, faSignInAlt, faSignOutAlt, faUserAlt, faUserEdit } from "@fortawesome/free-solid-svg-icons";
import LocalStorage from "../api/local";

export default function GHeader() {
    const [key, setKey] = useState('')
    const navigate = useNavigate()

    const [title, setTitle] = useState('用户')

    const handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key === 'Enter') {
            search()
        }
    }

    const search = () => {
        if (key.trim().length === 0)
            return
        navigate(`/s/${key.trim()}`)
    }

    useEffect(() => {
        const username = LocalStorage.getUsername()
        if (username !== '') {
            setTitle(`${username} | 用户`)
        } else {
            setTitle('用户')
        }
    }, [LocalStorage.getUsername()])

    return (
        <Navbar collapseOnSelect bg="dark" variant="dark" fixed="top" expand="lg">
            <Container fluid>
                <Navbar.Brand className="btn ms-3 fw-bold" onClick={() => { window.location.replace('/') }}>
                    CodePaste
                </Navbar.Brand>
                <Navbar.Toggle />
                <Navbar.Collapse>
                    <div className="me-auto">
                        <InputGroup >
                            <InputGroup.Text>
                                索引
                            </InputGroup.Text>
                            <FormControl
                                type="search"
                                value={key}
                                onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setKey(e.target.value) }}
                                onKeyDown={handleKeyPress}
                            />

                            <Button variant="primary" onClick={search} >
                                <FontAwesomeIcon icon={faSearch} />
                            </Button>
                        </InputGroup>
                    </div>
                    <Nav className="me-4">
                        <NavDropdown title={title}>
                            <NavDropdown.Item href="/register">
                                <FontAwesomeIcon icon={faUserEdit} />
                                注册
                            </NavDropdown.Item>
                            <NavDropdown.Item href="/login">
                                <FontAwesomeIcon icon={faSignInAlt} />
                                登录
                            </NavDropdown.Item>
                            <NavDropdown.Item href="/logout">
                                <FontAwesomeIcon icon={faSignOutAlt} />
                                注销
                            </NavDropdown.Item>
                            <NavDropdown.Divider />
                            <NavDropdown.Item href="/user">
                                <FontAwesomeIcon icon={faUserAlt} />
                                用户管理
                            </NavDropdown.Item>
                        </NavDropdown>
                        <Nav.Link as={Link} to={"/about"}>关于</Nav.Link>
                        <Nav.Link target={"_blank"} href="https://github.com/josexy/codepaste">Github</Nav.Link>
                    </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    )
}