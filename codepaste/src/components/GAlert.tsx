import { Alert } from "react-bootstrap";

interface IProps {
    show: boolean,
    variant: string
    text: string,
}

export default function GAlert(props: IProps) {
    return (
        <div>
            <Alert
                show={props.show}
                variant={props.variant}
            >
                <div>
                    {props.text}
                </div>
            </Alert>
        </div>
    )
}