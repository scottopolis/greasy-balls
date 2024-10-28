export function getApiUrl() {
  switch (import.meta.env.mode) {
    case "development":
      return "http://localhost:8080";
    default:
      return "http://54.208.24.104";
  }
}
