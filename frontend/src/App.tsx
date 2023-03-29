import { useState } from "react";
import {
  DocumentSection,
  Navbar,
  SearchSection,
  Tabs,
  UploadSection,
} from "./components";
import { TabEnum } from "./types";

function App() {
  const [currentTab, setCurrentTab] = useState<TabEnum>(TabEnum.DOCUMENT);
  const [selectedDocument, setSelectedDocument] = useState("");
  let content: React.ReactNode;
  switch (currentTab) {
    case TabEnum.DOCUMENT:
      content = (
        <>
          <DocumentSection
            selectedDocument={selectedDocument}
            setSelectedDocument={setSelectedDocument}
          />
          <SearchSection documentId={selectedDocument} />
        </>
      );
      break;
    case TabEnum.UPLOAD:
      content = (
        <>
          <UploadSection />
        </>
      );
      break;
    default:
      content = (
        <>
          <DocumentSection
            selectedDocument={selectedDocument}
            setSelectedDocument={setSelectedDocument}
          />
          <SearchSection documentId={selectedDocument} />
        </>
      );
  }
  return (
    <div className="px-6 md:px-10 py-12">
      <Navbar />
      <Tabs setCurrentTab={setCurrentTab} currentTab={currentTab} />
      {content}
    </div>
  );
}

export default App;
