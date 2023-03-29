import { useState } from "react";
import { api } from "../../api";
import { Loader } from "../Loader/Loader";

type QueryResponseType = {
  content: string;
};

type SearchSectionPropType = {
  documentId: string;
};

export const SearchSection = ({ documentId }: SearchSectionPropType) => {
  const [userQuery, setUserQuery] = useState("");
  const [output, setOutput] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  async function handlerQuery() {
    setOutput(null);
    if (documentId === "") {
      alert("Please select a document to query on");
      return;
    }
    if (userQuery === "") {
      alert("Please enter a query");
      return;
    }
    setIsLoading(true);
    try {
      const res = await api.post("/query", {
        query: userQuery,
        document_id: documentId,
      });
      if (res.status === 200) {
        const data = res.data as QueryResponseType;
        setIsLoading(false);
        setOutput(data.content);
      }
    } catch (err) {
      setIsLoading(false);
      alert("Something went wrong");
    }
  }
  return (
    <div className="px-2 py-8">
      <label
        htmlFor="message"
        className="block mb-2 text-sm font-medium text-gray-900"
      >
        Ask your queries here
      </label>
      <textarea
        id="message"
        rows={4}
        onChange={(e) => setUserQuery(e.target.value)}
        className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500"
        placeholder="Write your thoughts here..."
      ></textarea>
      <button
        onClick={handlerQuery}
        className="bg-blue-500 ml-[0.5px] text-white px-4 py-2 text-sm rounded-sm mt-2 hover:opacity-95"
      >
        {isLoading ? <Loader /> : "Ask Away"}
      </button>
      {output && (
        <div className="bg-gray-100 border-[1px] mt-5 p-4 rounded-sm">
          <p className="font-medium">GPT Response:</p>
          <p className="text-gray-600">{output}</p>
        </div>
      )}
    </div>
  );
};
