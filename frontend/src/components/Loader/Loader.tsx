import React from "react";
import { StageSpinner } from "react-spinners-kit";
export const Loader = () => {
  return (
    <div className="flex justify-center">
      <StageSpinner
        size={1.2}
        sizeUnit={"rem"}
        color="#FFFFFF"
        className="margin-auto"
      />
    </div>
  );
};
