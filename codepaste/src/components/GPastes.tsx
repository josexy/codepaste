import React, { useEffect, useState } from "react";
import { Row, Col } from "react-bootstrap";
import request from "../api/request";
import { IPasteInfo, IResponseDataPagination } from "../api/utils";
import GPagination from "./GPagination";
import GPaste from "./GPaste";

interface IProps {
    setTotalPastes: React.Dispatch<React.SetStateAction<number>>
}

export default function GPastes({ setTotalPastes }: IProps) {
    const [pastes, setPastes] = useState<IPasteInfo[]>()

    // 每页最多有多少项
    const page_size = 5
    const [page, setPage] = useState(1)
    const [total, setTotal] = useState(0)

    // 获取分页数据
    const getPasteList = (page: number) => {
        request.get<IResponseDataPagination<IPasteInfo[]>>("/auth/pastes", {
            params: {
                page: page,
                page_size: page_size
            }
        }).then(res => {
            if (res.data.code === 1000) {
                if (res.data.data) {
                    setPastes(res.data.data)
                    setTotal(res.data.total)
                    setTotalPastes(res.data.total)
                }
            }
        }).catch(err => {
            console.log(err)
        })
    }

    useEffect(() => {
        getPasteList(page)
    }, [page])

    return (
        <>
            <Row>
                <Col>
                    {pastes && pastes.map((paste, index) =>
                        <GPaste
                            key={index}
                            paste={paste}
                        />
                    )}
                </Col>
            </Row>
            {
                total > 0 &&
                <GPagination page={page} page_size={page_size} total={total} setPage={setPage} />
            }
        </>
    )
}