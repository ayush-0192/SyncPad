"use client"
import {useParams} from "next/navigation"
import Editor from "@/component/Editor";
import { useState } from "react";

export default function DocumentPage() {
    const {id} = useParams()
    const[text, setText] = useState('')
    useEffect(() => {
        const fetchDocumentContent = async () => {
            try {
                    const response = await fetch(`http://localhost:8080/get-document/${id}`, {
                        method: "GET",
                    })
                    if(!response.ok) {
                    throw new Error("Failed to fetch document")
                }
                const data:FullDocument = await response.json();
                setText(data.content)
                lastSaved.current = data.content
            }catch(error) {
             console.log(error)
            }
        }
        fetchDocumentContent()
    },[id]);
    return (
        <div className="p-20">
            <Editor 
                documentId={id as string}
                text = {text}
                onChange= {setText}
            />
        </div>
        
    )
}