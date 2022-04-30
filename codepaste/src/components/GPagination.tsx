import React, { useEffect, useState } from "react";
import { Pagination } from "react-bootstrap";
import { GPanel363 } from "./GPanel";

interface IProps {
    // 当前访问第几页
    page: number
    // 页内项数
    page_size: number
    // 所有项数
    total: number
    setPage: React.Dispatch<React.SetStateAction<number>>
}

export default function GPagination({ page, page_size, total, setPage }: IProps) {

    const [pageArray, setPageArray] = useState<number[]>()
    const [pageCount, setPageCount] = useState(1)

    useEffect(() => {
        const count_page = Math.ceil(total / page_size)
        setPageCount(count_page)

        // console.log('page: ', page);
        // console.log('total: ', total);
        // console.log('page count: ', count_page);

        let pageArr: number[] = []
        if (count_page >= 1) {
            // 页数少于5，则全部显示
            if (count_page <= 5) {
                let i = 1;
                while (i <= count_page) {
                    pageArr.push(i);
                    i++;
                }
            } else {
                // 多页，则省略一部分
                if (page <= 3) {
                    pageArr = [
                        1,  // 首页
                        2,
                        3,
                        4,
                        5,
                        -1, // 省略页
                        count_page // 尾页
                    ]
                } else if (count_page - page <= 4) {
                    pageArr = [
                        1, // 首页
                        -1, // 省略页
                        count_page - 4,
                        count_page - 3,
                        count_page - 2,
                        count_page - 1,
                        count_page // 尾页
                    ]
                } else {
                    pageArr = [
                        1, // 首页
                        -1, // 省略页
                        page - 1,
                        page,
                        page + 1,
                        -1, // 省略页
                        count_page // 尾页
                    ]
                }
            }
        }

        setPageArray(pageArr)
    }, [page, page_size, total])

    const paginationItemClick = (val: number) => {
        if (val < 1 || val > pageCount || val === page)
            return
        setPage(val)
    }

    return (
        <GPanel363>
            <div className="d-flex justify-content-center">
                <Pagination>
                    {/* 上一页 */}
                    <Pagination.Prev
                        disabled={page === 1}
                        onClick={() => { paginationItemClick(page - 1) }} />
                    {
                        pageArray && pageArray.map((val, index) => {
                            if (val === -1) {
                                return <Pagination.Ellipsis key={index} disabled />
                            }
                            return <Pagination.Item
                                onClick={() => { paginationItemClick(val) }}
                                key={index}
                                active={val === page}
                            >
                                {val}
                            </Pagination.Item>
                        })
                    }
                    {/* 下一页 */}
                    <Pagination.Next
                        disabled={page === pageCount}
                        onClick={() => { paginationItemClick(page + 1) }} />
                </Pagination>
            </div>
        </GPanel363>
    )
}