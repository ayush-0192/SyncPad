"use client"
import {useParams} from "next/navigation"
import Editor from "@/component/Editor";
import { useState } from "react";

export default function DocumentPage() {
    const {id} = useParams()
    const[text, setText] = useState('')
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