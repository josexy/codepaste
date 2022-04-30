import GAlert from "./GAlert";
import { GPanel282 } from "./GPanel";

export default function GPageNotFound() {
    return (
        <GPanel282>
            <GAlert show={true} variant={"danger"} text={"PAGE NOT FOUND :("} />
        </GPanel282>
    )
}