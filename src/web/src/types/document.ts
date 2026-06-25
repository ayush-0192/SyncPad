export interface Document {
    id: string;
    title: string;
    updatedAt: string;
}

export interface FullDocument extends Document {
    content:string;
    version:number;
}