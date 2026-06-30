import { useRouter } from "next/navigation"
import { useEffect, useState } from "react"
import { Document, DocumentSummary } from "@/types/document"


export function useDocument() {
    const router = useRouter()

    const [documents,setDocument] = useState<Document[]>([])
    const [documentSummary,setDocumentSummary] = useState<DocumentSummary[]>([])

    useEffect(() => {
        fetchDocumentsTitleList()
    }, [])

    // fetch the list of title of the document to show in the sidebar
    const fetchDocumentsTitleList = async() => {
        try {
            const response = await fetch(
                "http://localhost:8080/get-document-title-list"
            );

            if(!response.ok) {
                throw new Error("Failed to fetch the response")
            }

            const data: DocumentSummary[] = await response.json();
            setDocumentSummary(data)
        }catch(error) {
            console.log(error)
        };
        
    };

    // create the document with untitled and document with empty content and shown in ui
    const createDocument = async() => {
        try {
            const response = await fetch(
                "http://localhost:8080/create-document",
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify({
                        title: "Untitled",
                        content: ""
                    }),
                }
                

            );

            const newDocument : Document = await response.json();
            setDocumentSummary((prev) => [
                {
                    id: newDocument.id,
                    title: newDocument.title,
                },
                ...prev
            ]);
            router.push(`/doc/${newDocument.id}`)
        }catch(error) {
            console.log(error)
        }
        
    }

    return {
        documents,
        createDocument,
        documentSummary
    }
}