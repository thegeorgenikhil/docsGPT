import { Document } from "../../types";

type DocumentCardProps = {
  document: Document;
  selectedDocument: string;
  setSelectedDocument: (id: string) => void;
};

export const DocumentCard = ({
  document,
  selectedDocument,
  setSelectedDocument,
}: DocumentCardProps) => {
  const activeCardStyling = (id: string) => {
    if (id === selectedDocument) {
      return "border-indigo-700 text-blue-700 font-medium shadow-sm shadow-blue-400/5";
    } else {
      return "border-gray-300 text-gray-700";
    }
  };

  const selectedTagVisibility = (id: string) => {
    if (id === selectedDocument) {
      return "visible";
    } else {
      return "invisible";
    }
  };

  return (
    <div
      onClick={() => {
        console.log("clicked");
        setSelectedDocument(document._id);
      }}
      className={`w-[300px] flex flex-col justify-between cursor-pointer text-sm rounded-md border-2 p-4 m-4 ${activeCardStyling(
        document._id
      )}`}
    >
      <h3>{document.name}</h3>
      <div className="flex justify-between items-center">
        <p
          className={`text-xs bg-blue-700 text-white w-fit p-1 rounded-full px-2 font-semibold  ${selectedTagVisibility(
            document._id
          )}`}
        >
          SELECTED
        </p>
        <div>
          <p className="text-right font-medium text-gray-500 text-xs">
            {(document.size / (1024 * 1024)).toFixed(2)} MB
          </p>
          <p className="text-right text-xs italic text-gray-500">
            {`Uploaded on ${new Date(document.createdAt).toLocaleDateString()}`}
          </p>
        </div>
      </div>
    </div>
  );
};
