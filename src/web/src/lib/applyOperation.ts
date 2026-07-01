import { Operation } from "@/types/document";

export default function applyOperation(text: string, operation: Operation) : string {
    if(operation.type === "insert") {
        return (
            text.slice(0,operation.position) + operation.text + text.slice(operation.position)
        );
    }

    return (
        text.slice(0,operation.position) + text.slice(operation.position + operation.length)
    )
}