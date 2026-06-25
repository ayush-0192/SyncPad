"use client"

import Link from "next/link"
import { useDocument } from "@/hooks/useDocument"

export default function Sidebar() {
    const {documents, createDocument} = useDocument()

    return (
        <aside className="w-72 border-r border-black flex flex-col">
            <button
               onClick={createDocument}
                className="m-4 p-3 bg-black rounded-lg text-white"
            >
                + New Document
            </button>

            <div className="flex-1 overflow-auto">
                {documents.map((doc) => (
                    <Link
                    key={doc.id}
                    href={`/doc/${doc.id}`}
                    >
                        <div className="p-3 hover:bg-gray-100">
                            doc.title
                        </div>
                    </Link>
                ))}
            </div>
        </aside>

    )
}