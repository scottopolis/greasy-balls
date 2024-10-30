export function getApiUrl() {
  switch (import.meta.env.mode) {
    case "development":
      return "http://localhost:8080";
    default:
      return "http://3.84.18.189";
  }
}
