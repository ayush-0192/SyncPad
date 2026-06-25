"use client"
import { useState } from "react";

interface EditorProps {
    documentId: string;
    text: string
    onChange: (value: string) => void
}

export default function Editor(
    { documentId,text,onChange }: EditorProps
) {

    return (
        <div>
            <textarea
            value = {text}
            onChange = {(e) => 
                onChange(e.target.value)
            }
            placeholder="Start Collabrating"
            className="
                w-[1200px]
                h-[800px]
                p-6
                bg-white
                border-4
                border-purple-500
                rounded-xl
                shadow-lg
                resize-none
                outline-none
                text-3xl
                text-black
                "
            />
        </div>
    )
}