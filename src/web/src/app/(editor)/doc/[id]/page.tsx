"use client"
import {useParams} from "next/navigation"
import Editor from "@/component/Editor";
import { useEffect, useRef, useState } from "react";
import { FullDocument } from "@/types/document";

export default function DocumentPage() {
    const {id} = useParams()
    const[text, setText] = useState('')
    const[saving, setSaving] = useState(false)
    const lastSaved = useRef("")
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

    const saveDocumentContent = async () => {
        setSaving(true)
        try {
            await fetch(`http://localhost:8080/update-doc/${id}`, {
                method: "PATCH",
                headers: {
                    "content-type" : "application/json"
                },
                body: JSON.stringify({
                    "content":text,
                }),
            });

        }catch(error) {
            console.error(error)
        }finally{
            lastSaved.current = text
            setSaving(false)
        }

    };

    useEffect(() => {
        if(text == lastSaved.current) return;
        const timeout = setTimeout(() => {
            saveDocumentContent()
        },1000);
        return () => clearTimeout(timeout);
    },[text])

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