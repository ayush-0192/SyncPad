export interface Document {
    id: string;
    title: string;
    updatedAt: string;
}

export interface FullDocument extends Document {
    content:string;
    version:number;
}

export interface DocumentSummary {
    id: string
    title: string
}


export type Operation =
  | {
      type: "insert";
      position: number;
      text: string;
      version: number;
    }
  | {
      type: "delete";
      position: number;
      length: number;
      version: number;
    };