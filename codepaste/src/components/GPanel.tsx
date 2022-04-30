import React from "react";
import { Col, Row } from "react-bootstrap";

interface IPanel {
    children?: React.ReactNode
}

function GPanel282({ children }: IPanel) {
    return (
        <Row>
            <Col md={2} lg={2}>
            </Col>
            <Col md={8} lg={8}>
                {children}
            </Col>
            <Col md={2} lg={8}>
            </Col>
        </Row>
    )
}

function GPanel101({ children }: IPanel) {
    return (
        <Row>
            <Col md={1} lg={1}>
            </Col>
            <Col md={10} lg={10}>
                {children}
            </Col>
            <Col md={1} lg={1}>
            </Col>
        </Row>
    )
}

function GPanel363({ children }: IPanel) {
    return (
        <Row>
            <Col md={3} lg={3}>
            </Col>
            <Col md={6} lg={6}>
                {children}
            </Col>
            <Col md={3} lg={3}>
            </Col>
        </Row>
    )
}

function GPanel444({ children }: IPanel) {
    return (
        <Row>
            <Col md={4} lg={4}>
            </Col>
            <Col md={4} lg={4}>
                {children}
            </Col>
            <Col md={4} lg={4}>
            </Col>
        </Row>
    )
}

export { GPanel101, GPanel282, GPanel363, GPanel444 }