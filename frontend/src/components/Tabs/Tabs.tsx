import { TabEnum } from "../../types";

type TabsPropType = {
  currentTab: TabEnum;
  setCurrentTab: (tab: TabEnum) => void;
};
export const Tabs = ({ currentTab, setCurrentTab }: TabsPropType) => {
  const activeTabStyling = (tab: TabEnum, currentTab: TabEnum) => {
    if (tab === currentTab) {
      return "text-blue-600 border-b-2 border-blue-600 rounded-t-lg";
    }
    return "border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300";
  };
  return (
    <div className="text-md font-medium text-center text-gray-500 border-b border-gray-200">
      <ul className="flex flex-wrap -mb-px">
        <li className="mr-2">
          <button
            onClick={() => setCurrentTab(TabEnum.DOCUMENT)}
            className={`inline-block p-4 ${activeTabStyling(
              TabEnum.DOCUMENT,
              currentTab
            )}`}
            aria-current="page"
          >
            Documents
          </button>
        </li>
        <li className="mr-2">
          <button
            onClick={() => setCurrentTab(TabEnum.UPLOAD)}
            className={`inline-block p-4 ${activeTabStyling(
              TabEnum.UPLOAD,
              currentTab
            )}}`}
          >
            Upload
          </button>
        </li>
      </ul>
    </div>
  );
};
