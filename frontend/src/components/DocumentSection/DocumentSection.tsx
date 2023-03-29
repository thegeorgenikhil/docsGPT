import { AxiosResponse } from "axios";
import { useEffect, useState } from "react";
import { api } from "../../api";
import { Document } from "../../types";
import { DocumentCard } from "../DocumentCard/DocumentCard";

type DocumentResponseType = {
  documents: Document[];
};

type DocumentSectionPropType = {
  selectedDocument: string;
  setSelectedDocument: (id: string) => void;
};

export const DocumentSection = ({
  selectedDocument,
  setSelectedDocument,
}: DocumentSectionPropType) => {
  const [documents, setDocuments] = useState<Document[]>([]);

  async function fetchDocument() {
    const res = await api.get("/documents");
    if (res.status === 200) {
      const data = res.data as DocumentResponseType;
      setDocuments(data.documents);
    }
  }

  useEffect(() => {
    fetchDocument();
  }, []);
  return (
    <>
      <section>
        {documents.length > 0 ? (
          <div className="flex flex-wrap">
            {documents.map((document) => (
              <DocumentCard
                selectedDocument={selectedDocument}
                setSelectedDocument={setSelectedDocument}
                document={document}
                key={document._id}
              />
            ))}
          </div>
        ) : (
          <div className="text-center p-16">
            <p className="text-xl text-gray-500">No Documents Uploaded</p>
          </div>
        )}
      </section>
    </>
  );
};
