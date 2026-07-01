import { Operation } from "@/types/document"

function findDiffPosition(oldText: string, newText: string): number {
    let i = 0
    while (i < newText.length && i < oldText.length && oldText[i] === newText[i]) {
        i++;
    }
    return i;
}

export default function generateOperation(oldText: string, newText: string, version: number): Operation | null {
    if (newText == oldText) {
        return null
    }

    const position = findDiffPosition(oldText,newText)

    if (newText.length > oldText.length) {
        const insertedText = newText.slice(position,position + (newText.length-oldText.length));
        return {
            type: "insert",
            position,
            text: insertedText,
            version,
        };
    }

    return {
        type: "delete",
        position,
        length: oldText.length-newText.length,
        version,
    };
}