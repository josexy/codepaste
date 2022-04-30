import { Col, Container, Row } from "react-bootstrap";
import { Outlet } from "react-router-dom";
import GFooter from './GFooter';
import GHeader from './GHeader';

function GContainer() {
    return (
        <Container fluid style={{ "height": "100%" }}>
            <Row>
                <Col md={12}>
                    <GHeader />
                </Col>
            </Row>
            <Row style={{ "marginTop": "75px" }}>
                <Col>
                    <Outlet />
                </Col>
            </Row>
            <Row>
                <Col md={12}>
                    <GFooter />
                </Col>
            </Row>
        </Container>
    );
}

export default GContainer;
