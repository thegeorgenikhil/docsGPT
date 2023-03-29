import React from "react";
import { api } from "../../api";
import { Loader } from "../Loader/Loader";

export const UploadSection = () => {
  const [file, setFile] = React.useState<File | null>(null);
  const [isLoading, setIsLoading] = React.useState(false);

  function handleFileChange(event: React.ChangeEvent<HTMLInputElement>) {
    if (event.target.files) {
      setFile(event.target.files[0]);
    }
  }

  async function handleUpload() {
    const formData = new FormData();
    if (file) {
      setIsLoading(true);
      formData.append("file", file);
      try {
        const res = await api.post("/documents/upload", formData);
        if (res.status === 200) {
          alert("File uploaded successfully");
        }
      } catch (err) {
        console.log(err);
        alert("Something went wrong");
      } finally {
        setIsLoading(false);
      }
    }
  }

  return (
    <div>
      <div className="pt-8">
        <label
          className="block mb-2 text-sm font-medium text-gray-900"
          htmlFor="file_input"
        >
          Upload file
        </label>
        <input
          className="block w-full text-sm text-gray-900 border border-gray-300 rounded-sm cursor-pointer bg-gray-50"
          aria-describedby="file_input_help"
          id="file_input"
          accept="application/pdf"
          type="file"
          onChange={handleFileChange}
        />
        <p className="mt-1 text-sm text-gray-500" id="file_input_help">
          *.pdf files only. Max file size 10MB.
        </p>
      </div>
      <button
        onClick={handleUpload}
        className="bg-blue-500 w-40 text-center text-white px-4 py-2 text-sm rounded-sm mt-2 hover:opacity-95"
      >
        {isLoading ? <Loader /> : "Upload"}
      </button>
    </div>
  );
};
