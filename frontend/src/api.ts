// axios

import axios from "axios";

export const api = axios.create({
  baseURL: "http://localhost:9000/api",
});

