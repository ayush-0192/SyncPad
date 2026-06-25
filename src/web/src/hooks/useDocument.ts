import { useRouter } from "next/navigation"
import { useState } from "react"
import { Document } from "@/types/document"
export function useDocument() {
    const router = useRouter()

    const [documents,setDocument] = useState<Document[]>([])

    const createDocument = () => {
        const newDocument = {
            id: crypto.randomUUID(),
            title: "untitled",
            updatedAt: new Date().toISOString()
        };
        setDocument((prev) => [
            newDocument,
            ...prev,
        ]);

        router.push(`/doc/${newDocument.id}`);
    };

    return {
        documents,
        createDocument
    }
}